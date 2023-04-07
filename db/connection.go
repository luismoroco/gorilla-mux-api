package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var setupDBConfig = os.Getenv("POSTGRESS_CONNECTION_STR")
var GatewayForConnection *gorm.DB

func SetupConnectionWithPostgreSQL() {
	var err error
	GatewayForConnection, err = gorm.Open(postgres.Open(setupDBConfig), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("Database setup SUCESSFULLY")
}
