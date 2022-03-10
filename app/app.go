package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
