# go-shunt

Package shunt enables you to create "middleware" for the the "database/sql" package.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-shunt

[![GoDoc](https://godoc.org/github.com/reiver/go-shunt?status.svg)](https://godoc.org/github.com/reiver/go-shunt)


## Example Usage
```
import (
	"github.com/reiver/go-shunt"

	"database/sql"
)

// ...

rows, err := db.Query(query)
if nil != err {
	return err
}
defer rows.Close()

iterator := head(rows, 5) // <---- iterator has the same methods as *sql.Rows (but isn't *sql.Rows).
rows, err = shunt.Rows(iterator) // <---- Here we are using shunt.Rows func.

for rows.Next() {
	if err := rows.Scan(&name); nil != err {
		return err
	}
}
if err := rows.Err(); nil != err {
	return err
}

```
