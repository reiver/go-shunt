package shunt


import (
	"io"
)


type Float64Iterator struct {

	Slice []float64

	i int

	Label string
}


func NewFloat64Iterator(float64s ...float64) *Float64Iterator {
	slice := make([]float64, len(float64s))
	for i, datum := range float64s {
		slice[i] = datum
	}

	iterator := Float64Iterator{
		Slice:slice,
		Label:"value",
	}

	return &iterator
}


func (iterator Float64Iterator) Close() error {
	return nil
}

func (iterator Float64Iterator) Columns() ([]string, error) {
	return []string{iterator.Label}, nil
}

func (iterator Float64Iterator) Err() error {
//@TODO: Should this ever return io.EOF?
	return nil
}

func (iterator Float64Iterator) Next() bool {
	return len(iterator.Slice) > iterator.i
}

func (iterator *Float64Iterator) Scan(dest ...interface{}) error {

//@TODO: Should we return an error if `dest` is the wrong length?

	if ! iterator.Next() {
//@TODO: Is this the correct error?
		return io.EOF
	}

	value := iterator.Slice[iterator.i]

	iterator.i++

	if 1 <= len(dest) {
		dest[0] = value
	}

	return nil
}
