package config

import (
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"time"
)

func InitDB() *sqlx.DB {
	conn := connectToDB(Viper)
	if conn == nil {
		log.Fatal("can't connect to database")
		return nil
	}
	return conn
}

// connectToDB tries to connect to PostgreSQL, and backs off until a connection
// is made, or we have not connected after 10 tries
func connectToDB(v *viper.Viper) *sqlx.DB {
	counts := 0

	// Default to disable if the development environment
	sslMode := "disable"
	if v.GetString("APP_ENV") == "production" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		v.GetString("DB_HOST"),
		v.GetString("DB_PORT"),
		v.GetString("DB_USER"),
		v.GetString("DB_PASSWORD"),
		v.GetString("DB_NAME"),
		sslMode,
	)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("PostgreSQL not yet ready...")
			log.Printf("Open DB Error: %v", err)
		} else {
			log.Print("connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Printf("Backing off for 1 second : %v", counts)
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

// openDB opens a connection to PostgreSQL using a DSN read
// from the environment variable DSN
func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
