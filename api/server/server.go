package server

import (
	"database/sql"
	"net/http"
)

// Handlers --
type handlers interface {
	Home(w http.ResponseWriter, req *http.Request)
}

// ConfigServer --
type ConfigServer struct {
	Addr string
	Mux  *http.ServeMux
}

// Svr --
type Svr struct {
	HTTPSrv http.Server
	db      *sql.DB
	*Store
	handlers
}

// NewServer --
func NewServer(client *sql.DB, ConfigServer ConfigServer) *Svr {
	return &Svr{
		db: client,
		Store: &Store{
			client: client,
		},
		HTTPSrv: http.Server{
			Addr:    ConfigServer.Addr,
			Handler: ConfigServer.Mux,
		},
	}
}
