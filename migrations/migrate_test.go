package migrations

import (
	"fmt"
	"log"
	"testing"
)

func TestMigrate(t *testing.T) {
	res, err := Migrate()
	if err != nil {
		log.Fatalf("Migration Messages: %v", err)
	}
	fmt.Printf("Migrations Messages: %v", res)
}
