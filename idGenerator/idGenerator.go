package idgenerator

import "math/rand"

const KEY_LENGTH = 12

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateUniqueString() string {
	// TODO: make it a separate service - save to a db, check for uniqueness etc.
	newKey := randStringBytes(KEY_LENGTH)
	return newKey
}

// TODO: when ID is given send a call to main DB service to make the loc not releasable (gRPC)

// Job1 - Check if there is 100 IDs available (messages: giveMe100Ids ->, receive100Ids <-)
//    1. no - generate 10 x 100 IDs -> 2
//    2. lock in main DB 100 IDs
//    3. return to a caller that saves them in cache for 1 hour

// Job2 - Release locked IDs in main db (by timestamp + 1 hour + 30 min buffer time)
