package main

import (
	"boularbahsmail/gowebserver/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // ./static/index.html

	// Routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handlers.LoginHandler)
	http.HandleFunc("/docs", handlers.DocsHandler)

	fmt.Println("Server running on PORT 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // Error running the server...
	}
}
