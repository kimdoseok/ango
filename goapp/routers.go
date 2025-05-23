package main

import (
	"net/http"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	repo *AlumniRepository
	serv *AlumniService
)

func Routes() {
	db, err = NewDatabase("mysql")
	if err != nil {
		panic("Gorm connection error: " + err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/alumni", AlumniMux())
}

func AlumniMux() http.Handler {
	db.AutoMigrate(Alumni{})
	repo = NewAlumniRepository(db)
	serv = NewAlumniService(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/count", http.HandlerFunc(serv.Count))
	mux.HandleFunc("/", http.HandlerFunc(serv.List))
	mux.HandleFunc("/:filter", http.HandlerFunc(serv.List))
	mux.HandleFunc("/count/:filter", http.HandlerFunc(serv.List))
	mux.HandleFunc("/get", http.HandlerFunc(serv.List))
	mux.HandleFunc("/save", http.HandlerFunc(serv.List))
	mux.HandleFunc("/get/:id", http.HandlerFunc(serv.List))
	mux.HandleFunc("/delete/:id", http.HandlerFunc(serv.List))
	return http.StripPrefix("/alumni", mux)
}
