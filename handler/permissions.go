package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)
type Permission struct {
	ID         int32            `json:"id"`
	Name       string           `json:"name" validate:"required"`
	Permission pgtype.Text      `json:"permission" validate:"required"`
	IsDeleted  pgtype.Bool      `json:"is_deleted"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}


func (server *Server) handlerCreatePermission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	permission := Permission{}
	err := json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		fmt.Println("------------",err)
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
	err = validate.Struct(permission)
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

	arg := db.CreatePermissionParams{
		Name:       permission.Name,
		Permission: permission.Permission,
		IsDeleted:     permission.IsDeleted,
	}

	permissionsInfo, err := server.store.CreatePermission(ctx, arg)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Failed to create permission")
		return
	}

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.Permission `json:"data"`
	}{
		Status:  true,
		Message: "Permission created successfully",
		Data:    []db.Permission{permissionsInfo},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetPermissionById(w http.ResponseWriter, r *http.Request) {
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

	permissionsInfo, err := server.store.GetPermission(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch permission",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.Permission `json:"data"`
	}{
		Status:  true,
		Message: "Permission retrieved successfully",
		Data:    []db.Permission{permissionsInfo},
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

func (server *Server) handlerGetAllPermission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	permissionsInfo, err := server.store.GetAllPermissions(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch permission",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.Permission `json:"data"` // Use []db.BrandsLanguage
	}{
		Status:  true,
		Message: "Permission retrieved successfully",
		Data:    permissionsInfo,
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

func (server *Server) handlerUpdatePermission(w http.ResponseWriter, r *http.Request){
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

	permission := db.Permission{}
	err = json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON request")
		return
	}

	arg := db.UpdatePermissionParams{
		ID:       	 int32(id),
		Name:    	permission.Name,
		Permission: permission.Permission,
		IsDeleted: 	permission.IsDeleted,
	}

	permissionsInfo,err:= server.store.UpdatePermission(ctx, arg)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Failed to fetch permission")
		return
	}

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.Permission `json:"data"`
	}{
		Status:  true,
		Message: "permission updated successfully",
		Data:    []db.Permission{permissionsInfo},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerDeletePermission(w http.ResponseWriter, r *http.Request) {
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

	permissionsInfo, err:= server.store.DeletePermission(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch permission",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.Permission `json:"data"`
	}{
		Status:  true,
		Message: "permission deleted successfully",
		Data:    []db.Permission{permissionsInfo},
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


