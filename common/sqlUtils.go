package common
/*


import (
	"github.com/jmoiron/sqlx"
	//_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

// this Pings the database trying to connect, panics on error
// use sqlx.Open() for sql.Open() semantics

var (
	session            *sqlx.DB
	driverName         = "mysql"
	dataSourceTemplate = "%s:%s@tcp(%s:3306)/%s"
)

func *GetSession() *sqlx.DB {
	if session == nil {
		var err error
		dataSource := fmt.Sprintf(dataSourceTemplate, AppConfig.DatabaseUser, AppConfig.DatabasePassword, AppConfig.SqlDatabaseHost, AppConfig.Database)
		//fmt.Println(dataSource)
		session, err = sqlx.Open(
			driverName,

			dataSource)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return session
}
*/