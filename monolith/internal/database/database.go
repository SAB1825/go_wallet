package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB(dsn string) (*sql.DB, error) {
	var db *sql.DB

	var err error
	maxRetries := 3
	backoff := 2 * time.Second

	for range maxRetries {
		log.Println("Connecting to database!!")

		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Successfully connected to the database")

				db.SetMaxOpenConns(25)
				db.SetMaxIdleConns(25)
				db.SetConnMaxLifetime(5 * time.Minute)

				return db, nil
			}

			log.Println("Failed to connect to the database" + err.Error())
			time.Sleep(backoff)
			backoff *= 2
		}
	}

	return nil, err
}
