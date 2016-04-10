package shuntdbsqldriver


import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)


const (
	keyPrefix = "shunt:"

	keyLen = 126
)


var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	runes = []rune{
		'0','1','2','3','4','5','6','7','8','9',
		'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
		'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z',
	}
)


func generateKey() string {

	var buffer bytes.Buffer

	buffer.WriteString(keyPrefix)

	now := time.Now().Unix()
	s := strconv.FormatInt(now, 36)
	buffer.WriteString(s)
	buffer.WriteRune('_')

	for buffer.Len() < keyLen {
		index := randomness.Intn(len(runes))

		r := runes[index]

		buffer.WriteRune(r)
	}

	return buffer.String()
}
