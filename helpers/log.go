package helpers

import (
	"fmt"
	"net/http"
	"time"
)

func LogWriteHeader(w http.ResponseWriter, msg string, statuscode int) {
	date_now := time.Now().Format("02-01-2006")
	time_now := time.Now().Format("15:04:05")
	fmt.Fprintln(w, "msg",msg," Status-code-Request: ",statuscode,"data-now",date_now,"time-now",time_now)
	w.WriteHeader(statuscode)
	fmt.Println(statuscode)
}
