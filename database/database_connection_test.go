package database

import (
    "os"
    "testing"
    "fmt"
)

var db = PostgreSQLDatabase{}
var dbPtr = &db

func TestDatabaseConnectCloud(t *testing.T) {
    err := db.ConnectCloud(os.Getenv("DB_NAME"),
                      os.Getenv("DB_HOST"),
                      os.Getenv("DB_USER"),
                      os.Getenv("DB_PASS"),
                      os.Getenv("INSTANCE_NAME"),
                      []byte(os.Getenv("CREDENTIALS_JSON")))
    if err != nil {
        t.Fatalf("Connect failed: %v", err)
    }
    stamp, err := db.Timestamp()
    if err != nil {
        t.Fatalf("Timestamp failed: %v", err)
    }
    fmt.Println(stamp)
}
