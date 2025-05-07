package dbops3

import (
	"database/sql"
	"fmt"

	//"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDetails struct {
	Id    int16
	Name  string `gorm:"not null"`
	Phone sql.NullInt16
	Email sql.NullString
}

var db *gorm.DB

func ConnectDb() {
	dsn := "root:bullet@/golangdb"

	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//
	db.AutoMigrate(&UserDetails{})
}

func CloseDb() {
	sqlDb, _ := db.DB()
	sqlDb.Close()
}

func InsertNewUser() {
	user1 := UserDetails{
		Name: "John Doe",
		Phone: sql.NullInt16{
			Int16: 2255,
			Valid: true,
		},
	}
	result := db.Create(user1)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func RetrieveUser() {
	var usr1 UserDetails
	//db.First(usr1) //gets the first record
	db.First(&usr1, 11) //gets user with id 11

}

func RetrieveUserList() {
	var users []UserDetails

	db.Find(&users) //get all users

	for i, val := range users {
		fmt.Printf("")
	}
}

func RetrieveUserBasedOnCOnditions() {
	var users []UserDetails

	db.Where("Name=?", "mmm").Find(&users) //get all users

	for i, val := range users {
		fmt.Printf("")
	}
}

//ALSO CHECK
//find and update a user
//method 1: find and then update then save
//method 2: directly udpate using where clause without retrieving the user using db.Model()

//delete user
