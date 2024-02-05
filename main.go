package main

import (
	"fmt"
	"log"
	"net/http"
)

// Shared functions
func __ParseForm(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() Error : %v", err)
		return
	}
}

// formHandler function to handle /form route => (* => POINTER : It's pointing to the http request)
func loginHandler(writer http.ResponseWriter, req *http.Request) {

	// Checking the request url path validity
	if req.URL.Path != "/form" {
		http.Error(writer, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// Checking the request method (It should be GET method)
	if req.Method != "GET" {
		http.Error(writer, "Method Is Not Supported (GET Method Requests Only)", http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "Hello From Form Route!!")
}

// docsHandler function to handle /docs route
func docsHandler(writer http.ResponseWriter, req *http.Request) {
	// Parsing the form
	__ParseForm(writer, req)

	fmt.Fprintf(writer, "POST Request Successed!!\n\n")
	email := req.FormValue("email")
	password := req.FormValue("password")
	fmt.Fprintf(writer, "Email : %s\nPassword : %s\n", email, password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // ./static/index.html

	// Routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", loginHandler)
	http.HandleFunc("/docs", docsHandler)

	fmt.Println("Server running on PORT 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // Error running the server...
	}
}
