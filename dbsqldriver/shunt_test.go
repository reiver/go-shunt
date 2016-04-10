package shuntdbsqldriver


import (
	"testing"
)


func TestShunt(t *testing.T) {

	const numTests = 500
	for testNumber := 0; testNumber < numTests; testNumber++ {

		emptyIterator := internalEmptyIterator{
			Cols: []string{"apple", "banana", "cherry"},
		}


		shunted, err := Shunt(emptyIterator)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v", testNumber, err)
			continue
		}

		if err := shunted.Close(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v", testNumber, err)
			continue
		}
	}
}
