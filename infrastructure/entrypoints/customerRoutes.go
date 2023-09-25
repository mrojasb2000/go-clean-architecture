package entrypoints

import (
	"Customers/domain/usecase"
	"Customers/infrastructure/drivenadapters"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	customerHandler := CustomerHandler{customerUseCase: usecase.NewCustomerUseCase(drivenadapters.NewCustomerRepositoryStub())}
	router := mux.NewRouter()
	router.HandleFunc("/", customerHandler.getAllCustomers).Methods(http.MethodGet)
	return router
}
