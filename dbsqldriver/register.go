package shuntdbsqldriver


import (
	"database/sql"
)


const (
	Name = "shunt"
)


func init() {
	register()
}


func register() {
	driver := newDriver()

	sql.Register(Name, driver)
}
