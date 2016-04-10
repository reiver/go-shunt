package shuntdbsqldriver


func Shunt(x Iterator) (Shunted, error) {

	var key string

	const limit = 7
	for numTries := 0 ; numTries < limit; numTries++ {

		key = generateKey()

		if err := insertIterator(key, x); nil == err {
			continue
		}

		break
	}

	if "" == key {
		return Shunted{""}, errTooManyTries
	}


	shunted := Shunted{
		key:key,
	}

	return shunted, nil
}


type Shunted struct {
	key string
}


func (s Shunted) Close() error {

	deleteIterator(s.key)

	return nil
}


func (s Shunted) DataSourceName() string {
	return s.key
}
