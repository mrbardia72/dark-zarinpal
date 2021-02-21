package helpers

import (
	"fmt"
	"net/http"
)

func LogWriteHeader(w http.ResponseWriter, msg string, statuscode int) {
	fmt.Fprintln(w, msg)
	w.WriteHeader(statuscode)
}
