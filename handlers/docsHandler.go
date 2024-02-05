package handlers

import (
	"boularbahsmail/gowebserver/shared"
	"fmt"
	"net/http"
)

// docsHandler function to handle /docs route
func DocsHandler(writer http.ResponseWriter, req *http.Request) {
	// Parsing the form
	shared.FormParser(writer, req)

	fmt.Fprintf(writer, "POST Request Successed!!\n\n")
	email := req.FormValue("email")
	password := req.FormValue("password")
	fmt.Fprintf(writer, "Email : %s\nPassword : %s\n", email, password)
}
