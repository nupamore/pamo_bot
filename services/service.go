package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver" // mysql driver
)

// Service : service
type Service struct{}

// DB : db
var DB *sql.DB

// DBsetup : db init
func DBsetup() error {
	// db setup
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		"charset=utf8&parseTime=True&loc=Local",
	)
	db, err := sql.Open("mysql", dsn)
	DB = db

	if err != nil {
		return err
	}
	return nil
}
