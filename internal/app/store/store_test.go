package store_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// url configuration for local testing
		databaseURL = "host=localhost dbname=golang_db sslmode=disable user=postgres password=1234"
	}
	os.Exit(m.Run())
}
