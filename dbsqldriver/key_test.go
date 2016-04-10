package shuntdbsqldriver


import (
	"testing"
)


func TestGenerateKey(t *testing.T) {

	const numTests = 100
	for testNumber := 0; testNumber < numTests; testNumber++ {

		key := generateKey()


		if expected, actual := keyLen, len(key); expected != actual {
			t.Errorf("For random test #%d, expected %d, but actually got %d; key = %q.", testNumber, expected, actual, key)
			continue
		}
	}
}
