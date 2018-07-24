package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&Product{Code: "L1234", Price: 3000})

	// Read
	var product Product
	db.First(&product, 1) // Find product with id 1
	spew.Dump(product)

	product = Product{}
	db.First(&product, "code = ?", "L1234") // find product with code L1212
	spew.Dump(product)
}
