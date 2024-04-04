package database

import (
	"fmt"
	"strconv"

	"github.com/emmaperez2197/go-fiber/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnnectDB() {

	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("Faild to parse database port")
	}

	dns := fmt.Sprintf("host=%s port=%d user=%s  password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Fail connection to database")
	}

	fmt.Println("Connection open to database")

}
