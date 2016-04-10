package shunt


import (
	// This strickly isn't needed (since it is imported elsewhere), but adding it here for clarifiy.
	// This cause the "shunt" database driver to be registered with the "database/sql" package.
	_ "github.com/reiver/go-shunt/dbsqldriver"
)
