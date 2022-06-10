package models

import "gorm.io/gorm"

type DatabaseManager struct{
	db *gorm.DB
}

var manager DatabaseManager
func SetManager(dbReference *gorm.DB){
	manager = DatabaseManager{
		db: dbReference,
	}
}