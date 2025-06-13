package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/arianaw15/ip-sentinel/service/country"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	countryStore := country.NewStore(s.db)
	countryHandler := country.NewHandler(countryStore)
	countryHandler.RegisterRoutes(subrouter)

	log.Println("Starting API server on", s.addr)
	return http.ListenAndServe(s.addr, subrouter)
}
