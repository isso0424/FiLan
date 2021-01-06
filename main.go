package main

import (
	"FiLan/server"
)

func main() {
	server.Serve("db.sqlite3", "~/.filan/storage")
}
