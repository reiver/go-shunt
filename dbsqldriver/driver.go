package shuntdbsqldriver


import (
	"database/sql/driver"
)


// internalDriver fits the `Driver` interface in the package "database/sql/driver".
//
// In the "register.go", in this packagee, it calls the sql.Register func and registers
// this driver with the name found in the `Name` variable in this package; which is: "shunt".
//
// The effect of this is that this becomes a database driver. Meaning that if someone
// runs code like the following them they will "get" this driver:
//
//	import (
//		_ "github.com/reiver/go-shunt/dbsqldriver"
//	
//		"database/sql"
//	)
//	
//	// ...
//	
//	db, err := sql.Open("shunt", dataSourceName)
//
// Of course, most peope are probably NOT going to make a call like this themselves.
//
// Instead they will probably do something like:
//
//	import (
//		"github.com/reiver/go-shunt"
//	)
//	
//	//...
//	
//	rows, err := shunt.Rows(iterator) // This conceptually "converts" an interface of type shunt.Iterator to *sql.Rows.
//
// The whole point of this driver is to work-around that "database/sql" does not return
// interfaces, but instead returns (point to) structs.
type internalDriver struct {}


func newDriver() driver.Driver {

	driver := internalDriver{}

	return &driver
}


func (driver *internalDriver) Open(name string) (driver.Conn, error) {
	if nil == driver {
		return nil, errNilReceiver
	}

	return newConn(name)
}
