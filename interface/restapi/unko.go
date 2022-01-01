package restapi

import (
	"fmt"
	"net/http"
)

func Unko(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintln(w, "うんこ OK!!!")
		if err != nil {
			fmt.Println(err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintln(w, "Unko Request")
		if err != nil {
			fmt.Println(err)
		}
	}
}
