package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Store
	//tokenMaker token.Maker
	router *mux.Router
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store, config util.Config) (*Server, error) {

	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/book/insert", server.handlerInsertBook)
	//router.HandleFunc("/users/login", server.loginUser)
	//router.HandleFunc("/tokens/renew_access", server.refreshToken)

	//router.HandleFunc("/upload/video", server.uploadVideoToS3).Methods("POST")
	//router.HandleFunc("/videos", server.listAllVideos).Methods("GET")
	server.router = router 
	router.HandleFunc("/role", server.handlerCreateRole).Methods("POST")
	router.HandleFunc("/role/{id}", server.handlerGetRoleById).Methods("GET")
	router.HandleFunc("/role", server.handlerGetAllRole).Methods("GET")
	router.HandleFunc("/role/{id}", server.handlerUpdateRole).Methods("PUT")
	router.HandleFunc("/role/{id}", server.handlerDeleteRole).Methods("DELETE")


	router.HandleFunc("/permission", server.handlerCreatePermission).Methods("POST")
	router.HandleFunc("/permission/{id}", server.handlerGetPermissionById).Methods("GET")
	router.HandleFunc("/permission", server.handlerGetAllPermission).Methods("GET")
	router.HandleFunc("/permission/{id}", server.handlerUpdatePermission).Methods("PUT")
	router.HandleFunc("/permission/{id}", server.handlerDeletePermission).Methods("DELETE")


	router.HandleFunc("/rolepermission", server.handlerCreateRolePermission).Methods("POST")
	router.HandleFunc("/rolepermission/{id}", server.handlerGetRolePermissionById).Methods("GET")
	router.HandleFunc("/rolepermission", server.handlerGetAllRolePermission).Methods("GET")
	router.HandleFunc("/rolepermission/{id}", server.handlerUpdateRolePermission).Methods("PUT")
	router.HandleFunc("/rolepermission/{id}", server.handlerDeleteRolePermission).Methods("DELETE")

	router.HandleFunc("/country", server.handlerCreateCountry).Methods("POST")
	router.HandleFunc("/country/{id}", server.handlerGetCountryById).Methods("GET")
	router.HandleFunc("/country", server.handlerGetAllCountry).Methods("GET")
	router.HandleFunc("/country/{id}", server.handlerUpdateCountry).Methods("PUT")
	router.HandleFunc("/country/{id}", server.handlerDeleteCountry).Methods("DELETE")

	router.HandleFunc("/author", server.handlerCreateAuthor).Methods("POST")
	router.HandleFunc("/author/{id}", server.handlerGetAuthorById).Methods("GET")
	router.HandleFunc("/author", server.handlerGetAllAuthor).Methods("GET")
	router.HandleFunc("/author/{id}", server.handlerUpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{id}", server.handlerDeleteAuthor).Methods("DELETE")
}


// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	err := http.ListenAndServe(address, server.router)
	if err != nil {
		return err
	}
	return nil
}

