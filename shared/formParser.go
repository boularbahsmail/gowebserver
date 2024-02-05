package shared

import (
	"fmt"
	"net/http"
)

func FormParser(writer http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() Error : %v", err)
		return
	}
}
