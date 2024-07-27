package idgenerator

import (
	"errors"
	"log"
	idgeneratorclient "poll-lite/id-generator-client"
)

const KEY_LENGTH = 12
const KEYS_TO_GENERATE_COUNT = uint32(100)

var cache = []string{}

func TakeNextUniqueKey() (string, error) {
	if !generatedKeysAvailable() {
		generatedKeys, err := idgeneratorclient.GenerateIds(KEYS_TO_GENERATE_COUNT) //generateKeys(KEYS_TO_GENERATE_COUNT)
		if err != nil {
			log.Fatalf(err.Error())
			return "", err
		}

		saveKeys(&generatedKeys)
	}

	nextKey, err := popNextFreeKey()

	if err != nil {
		log.Fatalf(err.Error())
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
		return "", errors.New("there are no keys in the cache")
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

// TODO: when ID is given send a call to main DB service to make the loc not releasable (gRPC)

// Job1 - Check if there is 100 IDs available (messages: giveMe100Ids ->, receive100Ids <-)
//    1. no - generate 10 x 100 IDs -> 2
//    2. lock in main DB 100 IDs
//    3. return to a caller that saves them in cache for 1 hour

// Job2 - Release locked IDs in main db (by timestamp + 1 hour + 30 min buffer time)
