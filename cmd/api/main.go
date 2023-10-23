package main

import (
	"context"

	"log/slog"

	conn "github.com/manoj-negi/bookshop-adminapi/config/database"
	db "github.com/manoj-negi/bookshop-adminapi/db/sqlc"
	"github.com/manoj-negi/bookshop-adminapi/handler"
	util "github.com/manoj-negi/bookshop-adminapi/utils"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Info("cannot load config", err)
	}

	dbConn, err := conn.NewPostgres(context.Background(), config.DB_URI)
	if err != nil {
		slog.Info("cannot connect to db", err)
	}
	store := db.NewStore(dbConn.DB) //db.New(dbConn.DB)

	server, err := handler.NewServer(store, config)
	if err != nil {
		slog.Info("cannot create server")
	}

	err = server.Start(":8080")
	if err != nil {
		slog.Info("=======", err)
	}

}
