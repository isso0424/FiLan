package main

import (
	"FiLan/server"
)

func main() {
	err := server.Serve("db.sqlite3", "~/.filan/storage")
	if err != nil {
		panic(err)
	}
}
