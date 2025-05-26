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
func SetupPgsql() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=America/New_York",
		os.Getenv("DB_PGSQL_HOST"),
		os.Getenv("DB_PGSQL_USER"),
		os.Getenv("DB_PGSQL_PASSWORD"),
		os.Getenv("DB_PGSQL_PORT"),
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	result := map[string]interface{}{}
	db.Table("pg_user").First(&result, "username = ?", os.Getenv("DB_PGSQL_USER"))
	_, ok := result["usename"]
	if !ok {
		_ = db.Exec(fmt.Sprintf("CREATE USER IF NOT EXISTS %s WITH PASSWORD '%s';", os.Getenv("DB_PGSQL_USER"), os.Getenv("DB_PGSQL_PASSWORD")))
		_ = db.Exec(fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", os.Getenv("DB_PGSQL_DBNAME"), os.Getenv("DB_PGSQL_USER")))
		_ = db.Exec(fmt.Sprintf("ALTER USER %s WITH SUPERUSER;", os.Getenv("DB_PGSQL_USER")))
		_ = db.Exec(fmt.Sprintf("ALTER USER %s WITH PASSWORD '%s';", os.Getenv("DB_PGSQL_USER"), os.Getenv("DB_PGSQL_PASSWORD")))
	}
	_ = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s; ", os.Getenv("DB_PGSQL_DBNAME")))
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	sqlDB.Close()
}

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
		SetupPgsql()
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
	db, err = NewDatabase(os.Getenv("DB_USED"))
	if err != nil {
		log.Println(err)
	}
	//db.AutoMigrate(Alumni{})
	repo_alumni := NewAlumniRepository(db)
	serv_alumni := NewAlumniService(repo_alumni)

	http.HandleFunc("/", http.HandlerFunc(serv_alumni.List))
	http.HandleFunc("/count", http.HandlerFunc(serv_alumni.Count))
	http.HandleFunc("/:filter", http.HandlerFunc(serv_alumni.List))
	http.HandleFunc("/count/:filter", http.HandlerFunc(serv_alumni.List))
	http.HandleFunc("/get", http.HandlerFunc(serv_alumni.List))
	http.HandleFunc("/save", http.HandlerFunc(serv_alumni.List))
	http.HandleFunc("/get/:id", http.HandlerFunc(serv_alumni.List))
	http.HandleFunc("/delete/:id", http.HandlerFunc(serv_alumni.List))

	// Define routes
	//Routes()

	// Start the server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":81", nil))
}
