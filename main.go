package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"main.go/database"
	"database/sql"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "26032004"
	dbname = "product"
)


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
	// Update the connection string to match your database name
	database.Connect("postgres://postgres:26032004@localhost/product?sslmode=disable")

	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("client/static")))) //import static/styles css file
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images")))) //import images

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	products, err := database.GetProducts()
	if err != nil {
			http.Error(w, "Unable to load products", http.StatusInternalServerError)
			return
	}

	tmpl, err := template.ParseFiles("client/templates/amazon.html")
	if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
	}

	tmpl.Execute(w, products)
}
