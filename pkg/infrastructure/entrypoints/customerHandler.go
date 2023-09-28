package entrypoints

import (
	"Customers/internal/domain/usecase"
	"Customers/pkg/infrastructure/drivenadapters"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	customerUseCase usecase.CustomerUseCase
}

func NewCustomerHandler() CustomerHandler {
	customerUseCase := usecase.NewCustomerUseCase(drivenadapters.NewCustomerRepositoryDb())
	return CustomerHandler{customerUseCase: customerUseCase}
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

func (c *CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)
	if err != nil {
		fmt.Println("Error id conversion process")
		return
	}

	w.Header().Add("Content-Type", "application/json")

	customer, err := c.customerUseCase.GetCustomerById(id)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Customer not found"))
	}
	json.NewEncoder(w).Encode(customer)
}
