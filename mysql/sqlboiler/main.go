package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryutah/go-sample-codes/mysql/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/foo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	boil.SetDB(db)
	foos, err := models.Foos().All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	for _, foo := range foos {
		log.Printf("%#v", foo)
	}
}
