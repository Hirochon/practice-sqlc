package restapi

import (
	"database/sql"
	"fmt"
	"net/http"

	sqlc "github.com/Hirochon/practice-sqlc/infra/gen/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

func Unko(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctx := r.Context()

		db, err := sql.Open("mysql", "noda:nodanoda@(mysql:3306)/nodasql?parseTime=true&autocommit=0")
		if err != nil {
			fmt.Println(err)
		}

		queries := sqlc.New(db)

		allUnko, err := queries.GetAllUnko(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v¥n", allUnko)

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprintln(w, "うんこ OK!!!")
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
