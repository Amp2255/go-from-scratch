package dbops2

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Demotexts struct {
	some_txt string
}

var Db *sql.DB

func CheckDbConnection() {
	Db, err := sql.Open("mysql", "testuser:bullet@/golangdb")
	// handle error, if any.
	if err != nil {
		panic(err)
	}
	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func CloseDbConnection() {
	Db.Close()
}
