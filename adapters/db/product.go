package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paulocheloni/gohex/application"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stnt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stnt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}
