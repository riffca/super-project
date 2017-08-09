package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Model struct {
}

var connection *gorm.DB

func InitDB(host, user, dbname, password, port string) {
	var err error

	conOptions := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", host, user, dbname, password, port)
	connection, err = gorm.Open("postgres", conOptions)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	connection.LogMode(true)
}

func New() *gorm.DB {
	return connection.New()
}
