package routes

import (
	"swTest/controllers"

	"github.com/gorilla/mux"
)

var AirCompanyRoutes = func(router *mux.Router) {
	/*router.HandleFunc("/aircompanyscontroller/airline", controllers.CreateAirCompany).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/airline/{code}", controllers.DeleteAirCompany).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/provider", controllers.CreateProvider).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/provider/{id}", controllers.DeleteProvider).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/airline/{code}/providers/{id:.*}", controllers.RedactProvidersList).Methods("PUT")
	router.HandleFunc("/aircompanyscontroller/schema", controllers.CreateSchema).Methods("POST")*/
	router.HandleFunc("/aircompanyscontroller/schema/{name}", controllers.GetSchema).Methods("GET")
	/*router.HandleFunc("/aircompanyscontroller/schema/{id}", controllers.RedactSchema).Methods("PUT")
	router.HandleFunc("/aircompanyscontroller/schema/{id}", controllers.DeleteSchema).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/account", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/aircompanyscontroller/account/{accId}/schema/{scId}", controllers.RedactAccountSchema).Methods("PUT")
	router.HandleFunc("/aircompanyscontroller/account", controllers.DeteleAccount).Methods("DELETE")
	router.HandleFunc("/aircompanyscontroller/account/{id}/airlines", controllers.GetAccountAirlines).Methods("GET")
	router.HandleFunc("/aircompanyscontroller/provider/{id}/airlines", controllers.GetProviderAirlines).Methods("GET")*/
}
