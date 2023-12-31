package handler

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

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

func SetContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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
	router.Use(SetContentTypeJSON)
	router.HandleFunc("/book/insert", server.handlerInsertBook)


	//router.HandleFunc("/upload/video", server.uploadVideoToS3).Methods("POST")
	//router.HandleFunc("/videos", server.listAllVideos).Methods("GET")
	
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
  
	server.router = router 
}


// Start runs the HTTP server on a specific address.
func (server *Server) Start(port string) error {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	addr := "0.0.0.0"

	srv := &http.Server{
		Addr: addr + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      server.router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			slog.Info(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	slog.Info("shutting down")
	os.Exit(0)
	return nil

}

