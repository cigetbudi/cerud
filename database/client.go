package database

import (
	"cerud/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Debe *gorm.DB
var err error

func Connect(connectionString string) {
	Debe, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Ggal terhubung dengna database")
	}
	log.Println("Terhubung dengan database")
}

func Migrate() {
	Debe.AutoMigrate(&entities.Product{})
	log.Println("Migrasi database berhasil")
}
