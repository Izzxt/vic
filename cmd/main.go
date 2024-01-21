package main

import (
	"github.com/Izzxt/vic"
)

func main() {

	dsn := "root:root@tcp(localhost:49152)/vic?parseTime=true"

	vic := vic.Vic{
		Dsn: dsn,
	}
	vic.Init()

}
