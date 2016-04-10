package shuntdbsqldriver


type internalEmptyIterator struct{
	Cols []string
}


func (internalEmptyIterator) Close() error {
	return nil
}

func (iterator internalEmptyIterator) Columns() ([]string, error) {
	return iterator.Cols, nil
}

func (internalEmptyIterator) Err() error {
	return nil
}

func (internalEmptyIterator) Next() bool {
	return false
}

func (internalEmptyIterator) Scan(...interface{}) error {
	return nil
}
