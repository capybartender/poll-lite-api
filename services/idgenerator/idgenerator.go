package idgenerator

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const KEY_LENGTH = 12
const KEYS_TO_GENERATE_COUNT = 100

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var cache = []string{}

func randStringBytes(keyLength int) string {
	b := make([]byte, keyLength)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
}

// TODO: make it a separate service - save to a db, check for uniqueness etc.

func TakeNextUniqueKey() (string, error) {
	if !generatedKeysAvailable() {
		generatedKeys := generateKeys(KEYS_TO_GENERATE_COUNT)
		saveKeys(generatedKeys)
	}

	nextKey, err := popNextFreeKey()

	if err != nil {
		// log error
		fmt.Println(err.Error())
		return "", err
	}

	return nextKey, nil
}

func generatedKeysAvailable() bool {
	return len(cache) > 0
}

func popNextFreeKey() (string, error) {
	lastIndex := len(cache) - 1

	if lastIndex < 0 {
		return "", errNoKeysInCacheError
	}

	element := cache[lastIndex]
	cache = cache[:lastIndex]

	return element, nil
}

func saveKeys(nextKeys *[]string) error {
	cache = append(cache, *nextKeys...)
	//todo: use real DB and real cache
	return nil
}

func generateSingleKey(result *string, keyLength int, wg *sync.WaitGroup) {
	defer wg.Done()
	*result = randStringBytes(keyLength)
}

func generateKeys(keysCount int) *[]string {
	keys := make([]string, keysCount)

	var wg sync.WaitGroup
	wg.Add(keysCount)

	for i := range keysCount {
		go generateSingleKey(&keys[i], KEY_LENGTH, &wg)
	}

	wg.Wait()

	keys = *removeDuplicateStr(&keys)

	return &keys
}

// TODO: when ID is given send a call to main DB service to make the loc not releasable (gRPC)

// Job1 - Check if there is 100 IDs available (messages: giveMe100Ids ->, receive100Ids <-)
//    1. no - generate 10 x 100 IDs -> 2
//    2. lock in main DB 100 IDs
//    3. return to a caller that saves them in cache for 1 hour

// Job2 - Release locked IDs in main db (by timestamp + 1 hour + 30 min buffer time)
