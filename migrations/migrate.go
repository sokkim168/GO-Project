package migrations

import (
	"log"

	"example.com/booking/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() (string, error) {
	db, err := config.ConnectMySql("root", "MYSQL@root2024", "tcp(127.0.0.1:3309)", "room_booking")
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %v", err)
	}

	return "Successed", err
}
