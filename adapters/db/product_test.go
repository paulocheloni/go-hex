package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/paulocheloni/gohex/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func init() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db)
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
	stnt.Exec()

}

func insertProduct(db *sql.DB) {
	insert := `INSERT INTO products ("abc", "name", 0, "Disabled") VALUES (?, ?, ?, ?)`
	stnt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stnt.Exec()

}

func TestProductDb_Get(t *testing.T) {
	testing.Init()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "abc", product.GetName)
	require.Equal(t, 0.0, product.GetPrice)
	require.Equal(t, "Disabled", product.GetStatus)
}
