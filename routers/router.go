package routers

import (
	"../config"
	"../service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RouterMap() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/Bank{price}").HandlerFunc(service.Bank)
	r.Methods("GET").Path("/CallBack{price}").HandlerFunc(service.CallBack)
	log.Fatal(http.ListenAndServe(config.SERVER_PORT, r))
}