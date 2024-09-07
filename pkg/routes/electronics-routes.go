package routes

import (
	"github.com/alexandra1044/electronics-store-manager/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterElectronicsStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/electronics/", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/electronics/", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/electronics/{productID}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/electronics/{productID}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/electronics/{productID}", controllers.DeleteProduct).Methods("DELETE")
}
