package configs

import (
	"assignment2/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "root"
	dbname   string = "assignment2"
)

func NewPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Order{}) {
		db.AutoMigrate(&models.Order{})
	}

	if !db.HasTable(&models.Item{}) {
		db.AutoMigrate(&models.Item{})
	}

	return db
}
