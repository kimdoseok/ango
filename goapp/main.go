package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
)

// Define a struct to represent the data
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Slice to store items
var items []Item
var nextID = 1

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
	} else {
		db, err = gorm.Open(sqlite.Open(os.Getenv("DB_SQLITE_ALUMNI")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //.Silent
		})
	}

	if err != nil {
		log.Println("Gorm connection error: ", err)
		return nil, err
	}
	return db, nil
}

/*
// Handlers
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db,err :=NewDatabase()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
*/

func main() {
	db, err := NewDatabase("mysql")
	if err != nil {
		log.Println("Gorm connection error: ", err)
	}
	log.Printf("%+v", db)

	repo_alumni := NewAlumniRepository(db)
	serv_alumni := NewAlumniService(repo_alumni)

	// Define routes
	http.HandleFunc("/alumni", serv_alumni.List)
	http.HandleFunc("/alumni/:id", serv_alumni.Get)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
