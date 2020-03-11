package routes

import (
	"LunarAssignment/LunarServer/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RestServer struct {
	service *service.BridgeService
}

func StartRestServer(service *service.BridgeService) {
	rs := &RestServer{service: service}

	router := mux.NewRouter()

	router.HandleFunc("/update", rs.UpdateAccount).Methods("POST")
	router.HandleFunc("/account/{id}", rs.GetAccount).Methods("GET")
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	fmt.Println("Server: Listning on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
