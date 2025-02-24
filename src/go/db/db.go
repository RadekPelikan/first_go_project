package db

import (
	"FirstProject/src/go/helpers"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
)

type DbParams struct {
	username string
	password string
	host     string
	port     string
	database string
}

func GetDB(dbParams DbParams) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		dbParams.username,
		dbParams.password,
		dbParams.host,
		dbParams.port,
		dbParams.database,
	)
	fmt.Printf("dsn: %s\n", dsn)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
	return conn
}

func GetDefaultDbParams() DbParams {
	database_flag := flag.String("db", "", "Database name")

	flag.Parse()
	database := *database_flag
	if database == "" {
		database = os.Getenv("DB_DATABASE")
	}

	if database == "" {
		log.Fatal("db flag or DB_DATABASE env is not set")
		panic("Database name is empty")
	}
	dbParams := DbParams{
		username: helpers.WithDefault(os.Getenv("DB_USER"), "root"),
		password: helpers.WithDefault(os.Getenv("DB_PASS"), ""),
		host:     helpers.WithDefault(os.Getenv("DB_HOST"), "localhost"),
		port:     helpers.WithDefault(os.Getenv("DB_PORT"), "3306"),
		database: database,
	}
	return dbParams
}
