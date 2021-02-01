//go:generate go run http.go
package main

import (
	"fmt"
	"log"
	"net/http"

	dbConf "easyfood/config/db"
	httpConf "easyfood/config/http"
	"easyfood/pkg/graphql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", dbConf.DatabaseConn)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:    httpConf.GraphqlAddr,
		Handler: graphql.NewHandler(db),
	}

	fmt.Printf("Listening on %s", httpConf.GraphqlAddr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
