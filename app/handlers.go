package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/veezyjay/banka/service"
)

// type Customer struct {
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zip_code"`
// }

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

type CustomerHandlers struct {
	service service.CustomerService
}
