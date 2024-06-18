package config

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectMySql(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbname := os.Getenv("DATABASE")
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	db, err := ConnectMySql(user, pass, host, port, dbname)
	// db, err := ConnectMySql("root", "MYSQL@root2024", "localhost", "3309", "room_booking")
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	fmt.Printf("DB Connected: %v", db)
}
