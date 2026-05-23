package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	GetInstance() *sql.DB
	Close() error
}

type service struct {
	db *sql.DB
}

func (s *service) GetInstance() *sql.DB {
	return s.db
}

var (
	GOWAY_DIALECT = os.Getenv("GOWAY_DIALECT")
	database      = os.Getenv("GOWAY_DB_DATABASE")
	password      = os.Getenv("GOWAY_DB_PASSWORD")
	username      = os.Getenv("GOWAY_DB_USERNAME")
	port          = os.Getenv("GOWAY_DB_PORT")
	host          = os.Getenv("GOWAY_DB_HOST")
	schema        = os.Getenv("GOWAY_DB_SCHEMA")
	dbInstance    *service
)

func New() *service {
	if dbInstance != nil {
		return dbInstance
	}

	db, err := getSqlConnection(GOWAY_DIALECT)

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
