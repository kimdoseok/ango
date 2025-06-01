package main

import (
	"net/http"
	"os"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Routes(mux *http.ServeMux) {
	db, err := NewDatabase(os.Getenv("DB_USED"))
	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}

	//mux.Handle("/alumnus", AlumnusMux(mux))
	//mux.Handle("/group", GroupMux(mux))
	AlumnusMux(mux, db)
	GroupMux(mux, db)
}

func AlumnusMux(mux *http.ServeMux, db *gorm.DB) {
	submux := http.NewServeMux()

	repo := NewAlumnusRepository(db)
	serv := NewAlumnusService(repo)

	submux.HandleFunc("GET /count", http.HandlerFunc(serv.Count))
	submux.HandleFunc("GET /", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /:filter", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /count/:filter", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /get", http.HandlerFunc(serv.List))
	submux.HandleFunc("POST /save", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /get/:id", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /delete/:id", http.HandlerFunc(serv.List))
	mux.Handle("/alumnus/", http.StripPrefix("/alumnus", submux))

}

func GroupMux(mux *http.ServeMux, db *gorm.DB) {
	submux := http.NewServeMux()

	repo := NewGroupRepository(db)
	serv := NewGroupService(repo)

	submux.HandleFunc("GET /count", http.HandlerFunc(serv.Count))
	submux.HandleFunc("GET /", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /:filter", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /count/:filter", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /get", http.HandlerFunc(serv.List))
	submux.HandleFunc("POST /save", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /get/:id", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /delete/:id", http.HandlerFunc(serv.List))
	mux.Handle("/group/", http.StripPrefix("/group", submux))
}
