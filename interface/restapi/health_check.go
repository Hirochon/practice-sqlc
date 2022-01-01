package restapi

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(w, "Health Check OK!!!")
		if err != nil {
			fmt.Println(err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "Bad Request")
		if err != nil {
			fmt.Println(err)
		}
	}
}
