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
	mux.Handle("/alumni", AlumnusMux(mux))
}

func AlumnusMux(mux *http.ServeMux) http.Handler {
	db, err := NewDatabase(os.Getenv("DB_USED"))
	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}
	repo := NewAlumnusRepository(db)
	serv := NewAlumnusService(repo)

	mux.HandleFunc("/count", http.HandlerFunc(serv.Count))
	mux.HandleFunc("", http.HandlerFunc(serv.List))
	mux.HandleFunc(":filter", http.HandlerFunc(serv.List))
	mux.HandleFunc("/count/:filter", http.HandlerFunc(serv.List))
	mux.HandleFunc("/get", http.HandlerFunc(serv.List))
	mux.HandleFunc("/save", http.HandlerFunc(serv.List))
	mux.HandleFunc("/get/:id", http.HandlerFunc(serv.List))
	mux.HandleFunc("/delete/:id", http.HandlerFunc(serv.List))
	return http.StripPrefix("/alumni", mux)
}

func GroupMux(mux *http.ServeMux) http.Handler {
	db, err := NewDatabase(os.Getenv("DB_USED"))
	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}
	repo := NewGroupRepository(db)
	serv := NewGroupService(repo)

	mux.HandleFunc("/count", http.HandlerFunc(serv.Count))
	mux.HandleFunc("", http.HandlerFunc(serv.List))
	mux.HandleFunc(":filter", http.HandlerFunc(serv.List))
	mux.HandleFunc("/count/:filter", http.HandlerFunc(serv.List))
	mux.HandleFunc("/get", http.HandlerFunc(serv.List))
	mux.HandleFunc("/save", http.HandlerFunc(serv.List))
	mux.HandleFunc("/get/:id", http.HandlerFunc(serv.List))
	mux.HandleFunc("/delete/:id", http.HandlerFunc(serv.List))
	return http.StripPrefix("/alumni", mux)
}
