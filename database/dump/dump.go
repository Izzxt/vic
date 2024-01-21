package main

import (
	"os"
	"path/filepath"

	"github.com/jarvanstack/mysqldump"
)

func main() {
	dsn := "root:root@tcp(localhost:49152)/vic?parseTime=true"
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filepath.Join(wd, "database", "schema.sql"))
	if err != nil {
		panic(err)
	}

	if err := mysqldump.Dump(
		dsn, mysqldump.WithWriter(f),
	); err != nil {
		panic(err)
	}
}
