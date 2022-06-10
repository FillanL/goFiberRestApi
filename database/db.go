package database

import (
	"goFiberRestApi/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabase(){
	var db *gorm.DB
	var err error

	connectionString := config.Env.DBbUrl

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil{
		log.Fatalln("could not connect to postgres")
	}
	log.Println("connected to postgres")
	err = db.AutoMigrate()
}