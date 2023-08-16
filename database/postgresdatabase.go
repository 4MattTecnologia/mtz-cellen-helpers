package database

import (
    "database/sql"
    "fmt"
    "log"

    "cloud.google.com/go/cloudsqlconn"
    "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
    _ "github.com/lib/pq"
)

type PostgreSQLDatabase struct {
    DBConn *sql.DB
}

func (p *PostgreSQLDatabase) Close() {
    if p.DBConn != nil {
        p.DBConn.Close()
    }
}

func (p *PostgreSQLDatabase) Timestamp() (string, error) {
    rows, err := p.DBConn.Query(
        "SELECT now()")
    defer rows.Close()
    if err != nil {
        log.Printf("Error executing query in Timestamp(): %v", err.Error())
        return "", err
    }

    var time string
    rows.Next()
    if err := rows.Scan(&time); err != nil {
        log.Printf("Error  scanning query results in Timestamp(): %v", err.Error())
        return "", err
    }
    return time, nil
}

// gcp connectors described in:
// https://github.com/GoogleCloudPlatform/cloud-sql-go-connector#connecting-to-a-database
func (p *PostgreSQLDatabase) ConnectCloud(dbName string,
        dbHost string,
        dbUser string,
        dbPwd string,
        instanceName string,
        credentialsJSON []byte) error {
    if _, exists := sql.Drivers()["cloudsql-postgres"]; !exists {
        cleanup, err := pgxv4.RegisterDriver(
            "cloudsql-postgres", cloudsqlconn.WithCredentialsJSON(credentialsJSON))
        if err != nil {
            return err
        }
        // call cleanup when you're done with the database connection
        defer cleanup()
    }

    connstring := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable",
        instanceName,
        dbUser,
        dbPwd,
        dbName)
    db, err := sql.Open(
        "cloudsql-postgres",
        connstring,
    )
    if err != nil {
        return err
    }
    p.DBConn = db
    return nil
}

func (p *PostgreSQLDatabase) Connect(dbName string,
    dbHost string,
    dbPort string,
    dbUser string,
    dbPwd string) error {
    connstring := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
        dbHost,
        dbPort,
        dbUser,
        dbPwd,
        dbName)
    db, err := sql.Open(
        "postgres",
        connstring,
    )
    if err != nil {
        return err
    }
    p.DBConn = db
    return nil
}
