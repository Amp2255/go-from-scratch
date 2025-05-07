package dbops

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// create a struct with fields same as the coloumns in db table
type Product struct {
	Pid   int
	Name  string
	Price int
}

// a reference varibale of sql
var db *sql.DB

func CheckDbConnection() {
	//Open(db,DSN)
	//DSN complete form ::> username:password@protocol(address)/dbname?param=value
	//db, err := sql.Open("mysql", "testuser:bullet@/tcp(localhost:3306)/golangdb")
	//db, err := sql.Open("mysql", "testuser:bullet@/golangdb") //if localmachine omit the address part
	var err error
	db, err = sql.Open("mysql", "testuser:bullet@tcp(0.0.0.0:3306)/golangdb")
	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// Now its  time to connect with oru database,
	// database object has a method Ping.
	// Ping returns error, if unable connect to database.
	err = db.Ping()
	// handle error
	if err != nil {
		panic(err)
	}
	log.Println("Connected to MySQL database")

}

func CloseDbConnection() {
	closeErr := db.Close()
	if closeErr != nil {
		log.Fatal(closeErr)
	}
	fmt.Println("Db connection closed")

}

func InsertNewProduct() {

	var product Product
	product.Name = "P_Three"
	product.Price = 6234
	product.Pid = 17
	stmt, _ := db.Prepare("insert into product values(?,?,?)")
	res, inserterr := stmt.Exec(product.Pid, product.Name, product.Price)
	if inserterr != nil {
		fmt.Println("Database insert failed!")
	} else {
		num, _ := res.RowsAffected()
		fmt.Printf("%d Rows inserted\n", num)
	}
}

func ReadAllProducts() []Product {
	var prodlist []Product
	res, readErr := db.Query("select pid,name, price from product")

	if readErr != nil {
		fmt.Println("Error reading data from table")
	} else {
		for res.Next() {
			var p Product
			res.Scan(&p.Pid, &p.Name, &p.Price)
			prodlist = append(prodlist, p)
		}
	}
	return prodlist
}
