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

func (p *ProductDb) save(product application.ProductInterface) (application.ProductInterface, error) {
	stnt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stnt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}
	err = stnt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := p.save(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil

}
