package generator

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const KEY_LENGTH = 12

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var cache = []string{}

func randStringBytes(keyLength int) string {
	b := make([]byte, keyLength)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
}

func removeDuplicateStr(strSlice *[]string) *[]string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range *strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return &list
}

func generateSingleKey(result *string, keyLength int, wg *sync.WaitGroup) {
	defer wg.Done()
	*result = randStringBytes(keyLength)
}

// TODO: make length an argument
func GenerateKeys(keysCount uint32) ([]string, error) {
	keys := make([]string, keysCount)

	var wg sync.WaitGroup
	wg.Add(int(keysCount))

	for i := range keysCount {
		go generateSingleKey(&keys[i], KEY_LENGTH, &wg)
	}

	wg.Wait()

	uniqueKeys := removeDuplicateStr(&keys)

	err := saveGeneratedKeys(uniqueKeys)

	if err != nil {
		//TODO: either log error, or return it EVERYWHERE
		log.Fatalf("failed to save unique keys to the DB: %s", err.Error())
		return nil, err
	}

	return *uniqueKeys, nil
}

func saveGeneratedKeys(keys *[]string) error {
	// TODO: save to the DB instead of cache
	cache = append(cache, *keys...)
	return nil
}
