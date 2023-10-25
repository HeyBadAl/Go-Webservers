package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Use the database connection to perform database operations.
	// For example, you can query data from the database here.
	// db.Query() or db.Exec() can be used to execute SQL queries.

	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	fmt.Printf("Albums found: %v\n", albums)

	// Set up the route to the helloWorldHandler
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
