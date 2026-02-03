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
	a := Alumnus{}
	log.Printf("Alumnus Struct: %+v", a)

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
		db.AutoMigrate(&Alumnus{})
		db.AutoMigrate(&Group{})
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
	submux := http.NewServeMux()
	//Routes(mux)
	db, err := NewDatabase(os.Getenv("DB_USED"))
	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}

	auth := NewAuthService(db)
	submux.HandleFunc("GET /api/auth/login", http.HandlerFunc(auth.Login))
	submux.HandleFunc("GET /api/auth/logout", http.HandlerFunc(auth.Logout))

	repo_alumnus := NewAlumnusRepository(db)
	serv_alumnus := NewAlumnusService(repo_alumnus)
	submux.HandleFunc("GET /api/alumnus/count", http.HandlerFunc(serv_alumnus.Count))
	submux.HandleFunc("GET /api/alumnus/", http.HandlerFunc(serv_alumnus.List))
	submux.HandleFunc("GET /api/alumnus/{filter}", http.HandlerFunc(serv_alumnus.List))
	submux.HandleFunc("GET /api/alumnus/count/{filter}", http.HandlerFunc(serv_alumnus.List))
	submux.HandleFunc("GET /api/alumnus/get", http.HandlerFunc(serv_alumnus.List))
	submux.HandleFunc("POST /api/alumnus/save", http.HandlerFunc(serv_alumnus.List))
	submux.HandleFunc("GET /api/alumnus/het/{id}", http.HandlerFunc(serv_alumnus.List))
	submux.HandleFunc("POST /api/alumnus/import", http.HandlerFunc(serv_alumnus.Import))
	submux.HandleFunc("GET /api/alumnus/export", http.HandlerFunc(serv_alumnus.List))

	repo_group := NewGroupRepository(db)
	serv_group := NewGroupService(repo_group)
	submux.HandleFunc("GET /api/group/count", http.HandlerFunc(serv_group.Count))
	submux.HandleFunc("GET /api/group", http.HandlerFunc(serv_group.List))
	submux.HandleFunc("GET /api/group/{filter}", http.HandlerFunc(serv_group.List))
	submux.HandleFunc("GET /api/group/count/{filter}", http.HandlerFunc(serv_group.List))
	submux.HandleFunc("GET /api/group/get", http.HandlerFunc(serv_group.List))
	submux.HandleFunc("POST /api/group/save", http.HandlerFunc(serv_group.List))
	submux.HandleFunc("GET /api/group/get/{id}", http.HandlerFunc(serv_group.List))
	submux.HandleFunc("GET /api/group/delete/{id}", http.HandlerFunc(serv_group.List))

	// Start the server
	server := &http.Server{
		Addr:    ":81",
		Handler: submux,
	}

	fmt.Println("Server listening on port 81...")
	server.ListenAndServe()

}
