package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
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

	//Clean Up
	cleanUp(db)
}
func createRecords(db *gorm.DB) {
	log.Println("Creating records...")
	defer log.Println("Done!")
	sampleBooks := [5]Book{}
	sampleBooks[0] = Book{
		Title:  "How to Code Well",
		Author: "John Cena",
		Year:   2004,
	}
	sampleBooks[1] = Book{
		Title:  "Mind, Body, and Soul",
		Author: "Yikes Smith",
		Year:   2007,
	}
	sampleBooks[2] = Book{
		Title:  "Yoga Basics",
		Author: "Guru69",
		Year:   2001,
	}
	sampleBooks[3] = Book{
		Title:  "Mentoring 101",
		Author: "Mr. Dr. Strange",
		Year:   2020,
	}
	sampleBooks[4] = Book{
		Title:  "Up to No Good",
		Author: "Who Though",
		Year:   2055,
	}
	log.Print("Records: ")
	db.Create(&sampleBooks)
	spew.Dump(sampleBooks)
}
func cleanUp(db *gorm.DB) {
	db.Exec("DELETE FROM books")
}
