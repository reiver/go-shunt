package shuntdbsqldriver


import (
	"database/sql/driver"
	"io"
	"time"
)


type internalRows struct {
	iterator Iterator
	cachedColumns []string
}

func newRows(stmt *internalStmt, args []driver.Value) (driver.Rows, error) {

	if nil == stmt {
		return nil, errNilStmt
	}

	conn := stmt.conn
	if nil == conn {
		return nil, errNilConn
	}

	iterator := conn.iterator
	if nil == iterator {
		return nil, errNilIterator
	}

	cachedColumns, err := iterator.Columns()
	if nil != err {
		return nil, err
	}

	rows := internalRows{
		iterator:iterator,
		cachedColumns:cachedColumns,
	}

	return &rows, nil
}



func (rows *internalRows) Columns() []string {
	if nil == rows {
		panic(errNilReceiver)
	}

	return rows.cachedColumns
}


func (rows *internalRows) Close() error {
	if nil == rows {
		return errNilReceiver
	}

	return rows.iterator.Close()
}


func (rows *internalRows) Next(dest []driver.Value) error {
	if nil == rows {
		return errNilReceiver
	}

	iterator := rows.iterator
	if nil == iterator {
		return errNilIterator
	}


	if ! iterator.Next() {
		return io.EOF
	}


	lenColumns := len(rows.cachedColumns)

	scanDest := make([]interface{}, lenColumns)

	if err := iterator.Scan(scanDest...); nil != err {
		return err
	}


	for i, datum := range scanDest {
		if lenColumns <= i {
			break
		}

		switch x := datum.(type) {
		case int64:
			dest[i] = x
		case float64:
			dest[i] = x
		case bool:
			dest[i] = x
		case []byte:
			dest[i] = x
		case string:
			dest[i] = []byte(x)
		case time.Time:
			dest[i] = x
		default:
//@TODO: Should this produce an error instead?
			dest[i] = datum
		}
	}


	if err := iterator.Err(); nil != err {
		return err
	}


	return nil
}
