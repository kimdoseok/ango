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

	submux.HandleFunc("GET /alumnus/count", http.HandlerFunc(serv.Count))
	submux.HandleFunc("GET /alumnus/", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /alumnus/{filter}", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /alumnus/count/{filter}", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /alumnus/get", http.HandlerFunc(serv.List))
	submux.HandleFunc("POST /alumnus/save", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /galumnus/het/{id}", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /alumnus/delete/{id}", http.HandlerFunc(serv.List))
	mux.Handle("/alumnus", http.StripPrefix("/alumnus", submux))

}

func GroupMux(mux *http.ServeMux, db *gorm.DB) {
	submux := http.NewServeMux()

	repo := NewGroupRepository(db)
	serv := NewGroupService(repo)

	submux.HandleFunc("GET /group/count", http.HandlerFunc(serv.Count))
	submux.HandleFunc("GET /group", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /group/{filter}", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /group/count/{filter}", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /group/get", http.HandlerFunc(serv.List))
	submux.HandleFunc("POST /group/save", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /group/get/{id}", http.HandlerFunc(serv.List))
	submux.HandleFunc("GET /group/delete/{id}", http.HandlerFunc(serv.List))
	mux.Handle("/group", http.StripPrefix("/group", submux))
}
