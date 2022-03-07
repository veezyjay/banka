package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/veezyjay/banka/domain"
	"github.com/veezyjay/banka/service"
)

func Start() {
	
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}