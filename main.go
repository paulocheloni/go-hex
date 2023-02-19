package main

import (
	"database/sql"

	db2 "github.com/paulocheloni/gohex/adapters/db"
	"github.com/paulocheloni/gohex/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productAdapater := db2.NewProductDb(db)
	productService := application.NewProductService(productAdapater)
	product, _ := productService.Save("abcd", 3.0)

	productService.Enable(product)

}
