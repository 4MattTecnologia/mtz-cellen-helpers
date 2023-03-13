package database

type AbstractDatabase interface {
    Connect(db_name string, db_host string, db_port string,
       db_user string, db_pass string) error
    ConnectCloud(dbName string,
       dbHost string,
       dbUser string,
       dbPwd string,
       instanceName string,
       credentialsJSON []byte) error
}
