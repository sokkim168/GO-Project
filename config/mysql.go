package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySql(user, password, host, port, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + ";")
	dsn2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	database, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil
}

func GetDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbname := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	db, err := ConnectMySql(user, pass, host, port, dbname)
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	return db
}
