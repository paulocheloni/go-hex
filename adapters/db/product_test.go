package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/paulocheloni/gohex/adapters/db"
	"github.com/paulocheloni/gohex/application"
	"github.com/stretchr/testify/require"
)

func openDb() *sql.DB {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	// One connection keeps the in-memory database stable for the whole test.
	conn.SetMaxOpenConns(1)
	createTable(conn)
	insertProduct(conn)
	return conn
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		price float,
		status string
	)`
	stnt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stnt.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func insertProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)`
	stnt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stnt.Exec("abc", "name", 0, "Disabled")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestEnsureProductsTable(t *testing.T) {
	conn, err := sql.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	conn.SetMaxOpenConns(1)
	defer conn.Close()

	require.Nil(t, db.EnsureProductsTable(conn))
	require.Nil(t, db.EnsureProductsTable(conn))

	productDb := db.NewProductDb(conn)
	product := application.NewProduct()
	product.Name = "Notebook"
	product.Price = 19.9

	saved, err := productDb.Save(product)
	require.Nil(t, err)

	got, err := productDb.Get(saved.GetId())
	require.Nil(t, err)
	require.Equal(t, "Notebook", got.GetName())
	require.Equal(t, 19.9, got.GetPrice())
	require.Equal(t, product.GetStatus(), got.GetStatus())
}

func TestProductDb_Get(t *testing.T) {
	testing.Init()
	conn := openDb()
	defer conn.Close()
	productDb := db.NewProductDb(conn)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "abc", product.GetId())
	require.Equal(t, "name", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "Disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	testing.Init()
	conn := openDb()
	defer conn.Close()
	productDb := db.NewProductDb(conn)

	product := application.NewProduct()
	product.Name = "name"
	product.Price = 2

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "Enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}
