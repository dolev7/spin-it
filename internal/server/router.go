package server

import (
	"github.com/dolev7/spin-it/internal/users"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// SetupRouter initializes all routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", users.SignupHandler).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/swagger.json") // ✅ Correct relative path
	}).Methods("GET")

	return router
}
