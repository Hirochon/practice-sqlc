package main

import (
	"fmt"
	"github.com/Hirochon/practice-sqlc/interface/restapi"
	"net/http"
)

func main() {
	router := restapi.NewServer()
	http.Handle("/", router)
	fmt.Println(http.ListenAndServe(":8080", router))
}
