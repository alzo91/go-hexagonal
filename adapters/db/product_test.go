package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/alzo91/go-hexagonal/adapters/db"
	"github.com/alzo91/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp(){
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	insertProduct(Db)
}

func createTable(Db *sql.DB){
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
		);`

	stmt, err := Db.Prepare(table)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()						
}

func insertProduct(Db *sql.DB){
	stmt, err := Db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec("1", "Product 1", 10, "enabled")
}

func TestProductDb_Get(t *testing.T){
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "1", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
}

func TestProductDb_Save(t *testing.T){
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	

	product := application.NewProduct()
	product.ID = "2"
	product.Name = "Product 2"
	product.Price = 20
	product.Status = "disabled"

	productSaved, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetID(), productSaved.GetID())
	require.Equal(t, product.GetName(), productSaved.GetName())
	require.Equal(t, product.GetPrice(), productSaved.GetPrice())
	require.Equal(t,"disabled", productSaved.GetStatus())
	
	product.Name = "Product 2 updated"
	product.Status = "enabled"
	productUpdated, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), productUpdated.GetName())
	require.Equal(t,"enabled", productUpdated.GetStatus())
}