package config

import (
	"log"
	"post-api/model"
	"post-api/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = util.GetConfig("DB_PASSWORD")
	DB_NAME     string = util.GetConfig("DB_NAME")
	DB_HOST     string = util.GetConfig("DB_HOST")
	DB_PORT     string = util.GetConfig("DB_PORT")
)

func InitDB() {
	dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to database")
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.Post{}, &model.Tag{})
}
