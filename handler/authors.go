package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/gorilla/mux"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type Author struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name" validate:"required"`
	IsDeleted pgtype.Bool      `json:"is_deleted"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	author := Author{}
	err := json.NewDecoder(r.Body).Decode(&author)

	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusNotAcceptable,
		}
		w.Header().Set("Content-Type", "application/json")
		util.WriteJSONResponse(w, http.StatusNotAcceptable, jsonResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(author)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				jsonResponse := JsonResponse{
					Status:     false,
					Message:    "Invalid value for " + err.Field(),
					StatusCode: http.StatusNotAcceptable,
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(jsonResponse)
				return

			}
		}
	}

	arg := db.CreateAuthorParams{
		Name:    author.Name,
		IsDeleted: author.IsDeleted,
	}

	authorInfo, err := server.store.CreateAuthor(ctx, arg)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request1",
			StatusCode: http.StatusNotAcceptable,
		}
		util.WriteJSONResponse(w, http.StatusNotAcceptable, jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.Author `json:"data"`
	}{
		Status:  true,
		Message: "Author created successfully",
		Data:    []db.Author{authorInfo},
	}

	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetAuthorById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()
	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		errorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}
	authorInfo, err:= server.store.GetAuthor(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch author",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    []db.Author `json:"data"`
	}{
		Status:  true,
		Message: "Author retrieved successfully",
		Data:    []db.Author{authorInfo},
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to encode response",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}
}

func (server *Server) handlerGetAllAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	authorInfo, err := server.store.GetAllAuthors(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch author",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    []db.Author `json:"data"`
	}{
		Status:  true,
		Message: "author retrieved successfully",
		Data:    authorInfo,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to encode response",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}
}


func (server *Server) handlerUpdateAuthor(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut {
		errorResponse(w, http.StatusMethodNotAllowed, "Only PUT requests are allowed")
		return
	}

	ctx := r.Context()

	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		errorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}

	author := db.Author{}
	err = json.NewDecoder(r.Body).Decode(&author)

	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON request")
		return
	}

	arg := db.UpdateAuthorParams{
		ID:       	 int32(id),
		Name:    author.Name,
		IsDeleted: author.IsDeleted,
	}

	authorInfo,err:= server.store.UpdateAuthor(ctx, arg)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Failed to fetch author")
		return
	}

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.Author `json:"data"`
	}{
		Status:  true,
		Message: "author updated successfully",
		Data:     []db.Author{authorInfo},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		errorResponse(w, http.StatusMethodNotAllowed, "Only DELETE requests are allowed")
		return
	}
	ctx := r.Context()

	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		errorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}

	authorInfo, err:= server.store.DeleteAuthor(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch author",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.Author `json:"data"`
	}{
		Status:  true,
		Message: "author deleted successfully",
		Data:     []db.Author{authorInfo},
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to encode response",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}
}

