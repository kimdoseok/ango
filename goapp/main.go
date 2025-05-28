package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
)

// Define a struct to represent the data
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Slice to store items
func NewDatabase(usedb string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if usedb == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_MYSQL_USER"),
			os.Getenv("DB_MYSQL_PASSWORD"),
			os.Getenv("DB_MYSQL_HOST"),
			os.Getenv("DB_MYSQL_PORT"),
			os.Getenv("DB_MYSQL_DBNAME"))
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //.Silent
		})
	} else if usedb == "pgsql" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/New_York",
			os.Getenv("DB_PGSQL_HOST"),
			os.Getenv("DB_PGSQL_USER"),
			os.Getenv("DB_PGSQL_PASSWORD"),
			os.Getenv("DB_PGSQL_DBNAME"),
			os.Getenv("DB_PGSQL_PORT"),
		)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //.Silent
		})
	} else {
		db, err = gorm.Open(sqlite.Open(os.Getenv("DB_SQLITE_FILES")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //.Silent
		})
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	// Define routes
	mux := http.NewServeMux()
	Routes(mux)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":81", nil))
}
