package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	dbConn *gorm.DB
)

func CreateConnection() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	User := os.Getenv("SQL_USER")
	Pass := os.Getenv("SQL_PASSWORD")
	Host := os.Getenv("SQL_HOST")
	Port := os.Getenv("SQL_PORT")
	DB := os.Getenv("SQL_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host,
		Port,
		User,
		DB,
		Pass,
	)

	psql, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	dbConn = psql
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
