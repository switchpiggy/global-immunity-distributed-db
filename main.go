package main

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func main() {
	// Different API versions may expose different runtime behaviors.
	fdb.MustAPIVersion(630)

	// Open the default database from the system cluster
	db := fdb.MustOpenDefault()

}
