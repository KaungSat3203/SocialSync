package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from SocialSync Backend!"))
	}).Methods("GET")


	AuthRoutes(r)
	RegisterUserRoutes(r)
	// Add more like RegisterPostRoutes(r), etc.

	return r
}
