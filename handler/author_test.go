package handler

import(
	"context"
	"testing"
	"github.com/gorilla/mux"
	util "github.com/vod/utils"
	db "github.com/vod/db/sqlc"
	"github.com/stretchr/testify/require"
)

// var server *Server

type Server struct {
	config util.Config
	store  db.Querier
	//tokenMaker token.Maker
	router *mux.Router
}

func NewServer(store db.Querier, config util.Config) (*Server, error) {

	server := &Server{
		store:  store,
		config: config,
	}

	// server.setupRouter()
	return server, nil
}


var server *Server


func TestCreateAuthor(t *testing.T){
	arg:= db.CreateAuthorParams{
		Name: "Keshav",
	}

	author, err := server.store.CreateAuthor(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t,author)
	require.Equal(t, arg.Name, author.Name)

	require.NotZero(t,author.ID)
	require.NotZero(t,author.Isdeleted)
	require.NotZero(t,author.created_at)
	require.NotZero(t,author.updated_at)
	// require.NotZero(t,brand.created_at)	
}