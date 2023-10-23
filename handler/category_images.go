package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/manoj-negi/bookshop-adminapi/db/sqlc"
	util "github.com/manoj-negi/bookshop-adminapi/utils"
)

type CategoriesImage struct {
	ID         int32            `json:"id"`
	CategoryID int32            `json:"category_id" validate:"required"`
	Image      pgtype.Text      `json:"image" validate:"required"`
	IsDeleted  pgtype.Bool      `json:"is_deleted"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateCategoryImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	categoryImage := CategoriesImage{}
	err := json.NewDecoder(r.Body).Decode(&categoryImage)

	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusNotAcceptable,
		}

		util.WriteJSONResponse(w, http.StatusNotAcceptable, jsonResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(categoryImage)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				jsonResponse := JsonResponse{
					Status:     false,
					Message:    "Invalid value for " + err.Field(),
					StatusCode: http.StatusNotAcceptable,
				}

				json.NewEncoder(w).Encode(jsonResponse)
				return

			}
		}
	}

	arg := db.CreateCategoryImageParams{
		CategoryID: categoryImage.CategoryID,
		Image:      categoryImage.Image,
		IsDeleted:  categoryImage.IsDeleted,
	}

	categoryInfo, err := server.store.CreateCategoryImage(ctx, arg)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request1",
			StatusCode: http.StatusNotAcceptable,
		}
		util.WriteJSONResponse(w, http.StatusNotAcceptable, jsonResponse)
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CategoriesImage `json:"data"`
	}{
		Status:  true,
		Message: "Category Image created successfully",
		Data:    []db.CategoriesImage{categoryInfo},
	}

	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetCategoryImageById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()
	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		util.ErrorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}
	categoryInfo, err := server.store.GetCategoryImage(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch category image",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CategoriesImage `json:"data"`
	}{
		Status:  true,
		Message: "Category Image retrieved successfully",
		Data:    []db.CategoriesImage{categoryInfo},
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

func (server *Server) handlerGetAllCategoryImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	categoryInfo, err := server.store.GetAllCategoryImages(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch category image",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CategoriesImage `json:"data"`
	}{
		Status:  true,
		Message: "Category Image retrieved successfully",
		Data:    categoryInfo,
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

func (server *Server) handlerUpdateCategoryImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only PUT requests are allowed")
		return
	}

	ctx := r.Context()

	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		util.ErrorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}

	category := CategoriesImage{}
	err = json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON request")
		return
	}

	arg := db.UpdateCategoryImageParams{
		ID: int32(id),
	}

	if category.CategoryID > 0 {
		arg.SetCategoryID = true
		arg.CategoryID = category.CategoryID
	}

	if category.Image != emptyText {
		arg.SetImage = true
		arg.Image = category.Image
	}

	if category.IsDeleted.Valid && category.IsDeleted.Bool {
		arg.SetIsDeleted = true
		arg.IsDeleted = category.IsDeleted
	}

	categoryInfo, err := server.store.UpdateCategoryImage(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch category image")
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CategoriesImage `json:"data"`
	}{
		Status:  true,
		Message: "Category Image updated successfully",
		Data:    []db.CategoriesImage{categoryInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerDeleteCategoryImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only DELETE requests are allowed")
		return
	}
	ctx := r.Context()

	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		util.ErrorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}

	categoryInfo, err := server.store.DeleteCategoryImage(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch category image",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CategoriesImage `json:"data"`
	}{
		Status:  true,
		Message: "category image deleted successfully",
		Data:    []db.CategoriesImage{categoryInfo},
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
