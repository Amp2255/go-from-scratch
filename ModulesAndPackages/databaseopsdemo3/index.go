package main

import (
	"databasedemo2/dbops3"
)

func main() {
	dbops3.ConnectDb()
	dbops3.InsertNewUser()
	defer dbops3.CloseDb()

}
