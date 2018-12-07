package commons

import (
	"github.com/jinzhu/gorm"

	// Database Driver for Postgres and SQL Server
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Declare DB as type of gorm DB
var dbPostgreSQL *gorm.DB
var dbSQLServer *gorm.DB

// InitPostgreSQL : Function Init for InitPostgreSQL.
func InitPostgreSQL() *gorm.DB {

	Logger("info", "Initialize Connection to PostgreSQL", "InitPostgreSQL")

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=admin123 dbname=KreditPlus sslmode=disable")

	if err != nil {
		Logger("error", err.Error(), "InitPostgreSQL")
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	db.SingularTable(true)

	Logger("info", "PostgreSQL : Successfully connected!", "InitPostgreSQL")
	dbPostgreSQL = db

	return dbPostgreSQL
}

// InitSQLServer : Function Init for InitSQLServer.
func InitSQLServer() *gorm.DB {

	Logger("info", "Initialize Connection to SQL Server", "InitSQLServer")

	db, err := gorm.Open("mssql", "server=localhost;user id=sa;password=p@ndu4n4nd4;database = KreditPlus")

	if err != nil {
		Logger("error", "Error open db: "+err.Error(), "InitSQLServer")
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	db.SingularTable(true)

	Logger("info", "SQL Server : Successfully connected!", "InitSQLServer")
	dbSQLServer = db
	return dbSQLServer
}

// GetPostgreSQLDB : Function Get Database Connection for PostgreSQL.
func GetPostgreSQLDB() *gorm.DB {
	return dbPostgreSQL
}

// GetSQLServerDB : Function Get Database Connection for SQL Server.
func GetSQLServerDB() *gorm.DB {
	return dbSQLServer
}
