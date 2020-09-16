package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Book ...
type Book struct {
	gorm.Model
	Title  string
	Author string
	Year   int
}

func main() {
	//Set DSN and connect
	dsn := "zane:52455245@tcp(127.0.0.1:3306)/dummy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//Migrations
	db.AutoMigrate(&Book{})

	//Create Sample records
	createRecords(db)
}
func createRecords(db *gorm.DB) {
	log.Println("Creating records...")
	defer log.Println("Done!")

}
