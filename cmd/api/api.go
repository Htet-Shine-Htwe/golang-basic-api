package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dede182/revesion/service/user"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func StartNewServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	user.NewHandler(userStore).RegisterRoutes(subRouter)

	log.Println("Listening on : ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
