package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"main.go/database"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "26032004"
	dbname   = "product"
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

	//routes:
	http.HandleFunc("/", homeHandler)             //amazon.html route
	http.HandleFunc("/checkout", checkoutHandler) //checkout.html route
	http.HandleFunc("/orders", ordersHandler)     //orders.html route
	http.HandleFunc("/tracking", trackingHandler) //tracking.html route
	http.HandleFunc("/search", searchHandler)     //searching product via keywords

	//import css files. Js currently error
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("client/static")))) //import static/styles css file
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))        //import images

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homepage/amazon.html route
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

// Checkout page handler (checkout.html router)
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("client/templates/checkout.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Checkout page handler (checkout.html router)
func ordersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("client/templates/orders.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Checkout page handler (checkout.html router)
func trackingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("client/templates/tracking.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// this is the route that will allow user to "search" for a product via keywords
// inside the search bar.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")

	//ideally, this code redirect to an empty page with an error message for searching nothing
	//but, we won't want it here. Just return the main page if search nothing.

	// if searchQuery == "" {
	// 	http.Error(w, "Missing search query", http.StatusBadRequest)
	// 	return
	// }

	// Fetch products based on search query
	products, err := database.SearchProducts(searchQuery)
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
