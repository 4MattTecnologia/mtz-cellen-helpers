package database

import (
    "database/sql"
    "context"
    "fmt"
    "log"
    "net"

    "cloud.google.com/go/cloudsqlconn"
//    "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/stdlib"
    _ "github.com/lib/pq"
)

type PostgreSQLDatabase struct {
    DBConn *sql.DB
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
// https://github.com/GoogleCloudPlatform/golang-samples/blob/a09ece5d45a42a15d79e6d7c5afba77b088942ac/cloudsql/postgres/database-sql/connect_connector_iam_authn.go#L31
func (p *PostgreSQLDatabase) ConnectCloud(dbName string,
        dbHost string,
        dbUser string,
        dbPwd string,
        instanceName string,
        credentialsJSON []byte) error {
    d, err := cloudsqlconn.NewDialer(
        context.Background(),
        cloudsqlconn.WithCredentialsJSON(credentialsJSON))
    if err != nil {
        return fmt.Errorf("cloudsqlconn.NewDialer: %v", err)
    }

    dsn := fmt.Sprintf("user=%s database=%s", dbUser, dbName)
    config, err := pgx.ParseConfig(dsn)
    if err != nil {
        return err
    }

    config.DialFunc = func(ctx context.Context,
                           network string,
                           instance string) (
                           net.Conn, error) {
        return d.Dial(ctx, instanceName, cloudsqlconn.WithPrivateIP())
    }
    dbURI := stdlib.RegisterConnConfig(config)
    db, err := sql.Open("pgx", dbURI)
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
