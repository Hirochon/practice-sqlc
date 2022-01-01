package restapi

import "github.com/gorilla/mux"

func NewServer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheck)
	router.HandleFunc("/unko", Unko)
	return router
}
