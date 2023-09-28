package entrypoints

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routes() *mux.Router {
	customerHandler := NewCustomerHandler()
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}", customerHandler.getCustomerById).Methods(http.MethodGet)

	return router
}
