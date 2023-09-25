package entrypoints

import (
	"Customers/domain/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomerHandler struct {
	customerUseCase usecase.CustomerUseCase
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request customer")
}

func (c *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	customers, err := c.customerUseCase.GetAllCustomer()
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Customer not found"))
	}
	json.NewEncoder(w).Encode(customers)
}
