package config

import (
	"database/sql"

	db2 "github.com/paulocheloni/gohex/adapters/db"
	"github.com/paulocheloni/gohex/application"
)

func InjectService() application.ProductServiceInterface {
	var db, _ = sql.Open("sqlite3", "db.sqlite")
	var productAdapater = db2.NewProductDb(db)
	var productService = application.ProductService{Persistence: productAdapater}
	return &productService
}
