package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

func ConnectSQL() (*DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	hostName := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USER")
	hostPort := os.Getenv("DB_PORT")

	// pgConnStrings := fmt.Sprintf("port=%s host=%s user=%s "+"password=%s dbname=%s sslmode=disable", hostPort, hostName, userName, password, dbName)
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		userName,
		password,
		hostName,
		hostPort,
		dbName)

	fmt.Println("Datasource", url)

	d, err := sqlx.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(10)
	d.SetConnMaxLifetime(5 * time.Minute)

	dbConn.SQL = d
	return dbConn, err
}
