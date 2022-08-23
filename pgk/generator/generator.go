package generator

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetTime(oldTime time.Time, appendData time.Duration) time.Time {
	if oldTime.Unix() >= time.Now().Unix() {
		return oldTime.Add(appendData)
	} else {
		return time.Now().Add(appendData)
	}
}
