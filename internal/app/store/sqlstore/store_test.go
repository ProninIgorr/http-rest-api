package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURl string
)

// TestMain ...
func TestMain(m *testing.M) {
	databaseURl = os.Getenv("DATABASE_URL")
	if databaseURl == "" {
		databaseURl = "host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
