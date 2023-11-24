package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var defaultDialTimeout = 10 * time.Second
var mySQLDriverName = "mysql"

func NewMySQLDatabaseConn(conf Config) *sql.DB {
	return newDatabaseConn(conf, mySQLDriverName)
}

func newDatabaseConn(conf Config, driverName string) *sql.DB {
	db, err := sql.Open(driverName, buildDSN(conf))
	if err != nil {
		log.Fatalf("Failed to establish connection %v", err)
	}
	db.SetMaxOpenConns(conf.MaxOpenConn)
	db.SetMaxIdleConns(conf.MaxIdleConn)
	db.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	return db
}

func buildDSN(conf Config) string {
	mysqlConf := mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Passwd,
		Net:                  "tcp",
		Addr:                 conf.Address,
		DBName:               conf.DatabaseName,
		Loc:                  time.Now().Location(),
		Timeout:              defaultDialTimeout,
		ParseTime:            true,
		AllowNativePasswords: conf.AllowNativePasswords,
	}
	return mysqlConf.FormatDSN()
}
