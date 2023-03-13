package database

import (
	"os"
	"testing"
)

var db = PostgreSQLDatabase{}
var dbPtr = &db

func TestDatabaseConnect(t *testing.T) {
	err := db.Connect(os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
}
