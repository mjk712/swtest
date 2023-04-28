package routes

import (
	"swTest/handlers"

	"github.com/gorilla/mux"
)

var AirCompanyRoutes = func(router *mux.Router) {
	router.HandleFunc("/aircompanyscontroller/airline", handlers.CreateAirCompany).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/airline/{code}", handlers.DeleteAirCompany).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/provider", handlers.CreateProvider).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/provider/{id}", handlers.DeleteProvider).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/airline/{code}/providers/{id:.*}", handlers.RedactProvidersList).Methods("GET")
	router.HandleFunc("/aircompanyscontroller/schema", handlers.CreateSchema).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/schema/{name}", handlers.GetSchema).Methods("GET")
	router.HandleFunc("/aircompanyscontroller/schema/{id}", handlers.RedactSchema).Methods("PUT")
	router.HandleFunc("/aircompanyscontroller/schema/{id}", handlers.DeleteSchema).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/account", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/account/{accId}/schema/{scId}", handlers.RedactAccountSchema).Methods("PUT")
	router.HandleFunc("/aircompanyscontroller/account/{id}", handlers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/account/{id}/airlines", handlers.GetAccountAirlines).Methods("GET")
	router.HandleFunc("/aircompanyscontroller/provider/{id}/airlines", handlers.GetProviderAirlines).Methods("GET")
}
