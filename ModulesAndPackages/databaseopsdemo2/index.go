package main

import "databasedemo2/dbops2"

func main() {
	dbops2.CheckDbConnection()
	defer dbops2.CloseDbConnection()
}
