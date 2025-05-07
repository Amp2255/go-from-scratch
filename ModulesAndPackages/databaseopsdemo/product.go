package main

import (
	"databaseopsdemo/dbops"
	"fmt"
)

func main() {

	dbops.CheckDbConnection()
	defer dbops.CloseDbConnection()

	dbops.InsertNewProduct()
	prods := dbops.ReadAllProducts()

	for p, prods := range prods {
		fmt.Printf("Row %d ProductId:%d Name:%s Price:%d \n", p, prods.Pid, prods.Name, prods.Price)
	}

}
