package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
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

var form = `
	<form action="/" method="get">
	<input name="q"><button type="submit">Search</button>
	</form>
`

func main() {
	//Set DSN and connect
	dsn := "zane:52455245@tcp(127.0.0.1:3306)/dummy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//Migrations
	db.AutoMigrate(&Book{})
	//Clean Up
	cleanUp(db)

	//Create Sample records
	createRecords(db)

	//Set Up Webserver
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Serve Page
		serveSearchPage(w, r)
		log.Println("Page Served")
		//Do Search
		if search, ok := r.URL.Query()["q"]; ok {
			search := search[0]
			log.Println("Search:", search)
			results := performSearch(search, db)
			//Return results
			for _, result := range results {
				fmt.Fprintf(w, "<ul>")
				fmt.Fprintf(w, "<b>"+result.Title+"</b>")
				fmt.Fprintf(w, "<li>by "+result.Author+"</li>")
				fmt.Fprintf(w, "<li>"+strconv.Itoa(result.Year)+"</li>")
				fmt.Fprintf(w, "</ul>")
				fmt.Fprintf(w, "<br>")
			}
		}
	})
	panic(http.ListenAndServe("localhost:8888", router))
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

func serveSearchPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!doctype HTML>"+form)

}
func performSearch(search string, db *gorm.DB) []Book {
	searchObj := []Book{}
	db.Where("title LIKE ?", "%"+search+"%").Find(&searchObj)
	return searchObj
}
