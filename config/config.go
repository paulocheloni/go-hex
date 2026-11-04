package config

import (
	"database/sql"
	"log"

	db2 "github.com/paulocheloni/gohex/adapters/db"
	"github.com/paulocheloni/gohex/application"
)

func InjectService() application.ProductServiceInterface {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	// db.sqlite is created in the working directory on first use. The products
	// table is created with it, so a fresh checkout can run without a seed file.
	if err = db2.EnsureProductsTable(db); err != nil {
		log.Fatal(err)
	}
	var productAdapater = db2.NewProductDb(db)
	var productService = application.ProductService{Persistence: productAdapater}
	return &productService
}
