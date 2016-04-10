package shunt


import (
	"testing"
)


func TestRows(t *testing.T) {

	tests := []struct{
		Iterator *Float64Iterator
	}{
		{
			Iterator: &Float64Iterator{
				Label:"value",
				Slice: []float64{},
			},
		},
		{
			Iterator: &Float64Iterator{
				Label:"value",
				Slice: []float64{-5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}


	TestLoop: for testNumber, test := range tests {

		rows, err := Rows(test.Iterator)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one: %v", testNumber, err)
			continue
		}
		if nil == rows {
			t.Errorf("For test #%d, did a value of nil for rows but actually got that: %v", testNumber, rows)
			continue
		}

		length := 0
		for rows.Next() {
			var datum interface{}
			if err := rows.Scan(&datum); nil != err {
				t.Errorf("For test #%d and row #%d, did not expect an error from Scan but actually got one %v", testNumber, length, err)
				continue TestLoop
			}

			length++
		}

		if expected, actual := len(test.Iterator.Slice), length; expected != actual {
			t.Errorf("For test #%d, expected %d but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}
