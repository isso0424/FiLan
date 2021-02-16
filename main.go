package main

import (
	"FiLan/server"
	"FiLan/setup"
)

const (
	storageDir = "filan/storage"
	dbFile     = "db.sqlite3"
)

func main() {
	obj := setup.Setup("", storageDir, dbFile)
	if obj.Err != nil {
		panic(obj.Err)
	}

	err := server.Serve(&obj.FsRepository, &obj.DbRepository)
	if err != nil {
		panic(err)
	}
}
