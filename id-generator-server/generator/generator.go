package generator

import (
	"math/rand"
	"sync"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

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

func GenerateKeys(keysCount int, keyLength int) ([]string, error) {
	keys := make([]string, keysCount)

	var wg sync.WaitGroup
	wg.Add(keysCount)

	for i := range keysCount {
		go generateSingleKey(&keys[i], keyLength, &wg)
	}

	wg.Wait()

	uniqueKeys := removeDuplicateStr(&keys)

	return *uniqueKeys, nil
}
