package routers

import (
	"github.com/gorilla/mux"
	"github.com/mrbardia72/dark-zarinpal/config"
	"github.com/mrbardia72/dark-zarinpal/service"
	"log"
	"net/http"
)

func RouterMap() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/Bank{price}").HandlerFunc(service.Bank)
	r.Methods("GET").Path("/CallBack{price}").HandlerFunc(service.CallBack)
	log.Fatal(http.ListenAndServe(config.SERVER_PORT, r))
}