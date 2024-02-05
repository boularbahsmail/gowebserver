package handlers

import (
	"fmt"
	"net/http"
)

// formHandler function to handle /form route => (* => POINTER : It's pointing to the http request)
func LoginHandler(writer http.ResponseWriter, req *http.Request) {

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
