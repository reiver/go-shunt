package shunt


// Iterator is an interface that provides the exact same methods as *sql.Rows
// in the "database/sql" package.
//
// This exists because the "database/sql" package does not provide an interface
// for this itself.
type Iterator interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(...interface{}) error
}
