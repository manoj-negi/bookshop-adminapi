package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)
type RolesPermission struct {
	ID           int32            `json:"id"`
	RoleID       int32            `json:"role_id" validate:"required"`
	PermissionID int32            `json:"permission_id" validate:"required"`
	IsDeleted    pgtype.Bool      `json:"is_deleted"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateRolePermission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	rolepermission := RolesPermission{}
	err := json.NewDecoder(r.Body).Decode(&rolepermission)
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
	err = validate.Struct(rolepermission)
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

	arg := db.CreateRolePermissionParams{
		RoleID:    rolepermission.RoleID,
		PermissionID: rolepermission.PermissionID,
		IsDeleted: rolepermission.IsDeleted,
	}

	rolepermissionInfo, err := server.store.CreateRolePermission(ctx, arg)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Failed to create permission")
		return
	}

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.RolesPermission `json:"data"`
	}{
		Status:  true,
		Message: "Role Permission created successfully",
		Data:    []db.RolesPermission{rolepermissionInfo},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetRolePermissionById(w http.ResponseWriter, r *http.Request) {
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

	rolepermissionInfo, err := server.store.GetRolePermission(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch role permission",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.RolesPermission `json:"data"`
	}{
		Status:  true,
		Message: "Role Permission retrieved successfully",
		Data:    []db.RolesPermission{rolepermissionInfo},
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

func (server *Server) handlerGetAllRolePermission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	rolepermissionInfo, err := server.store.GetAllRolePermissions(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch role permission",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.RolesPermission `json:"data"` // Use []db.BrandsLanguage
	}{
		Status:  true,
		Message: "Role Permission retrieved successfully",
		Data:    rolepermissionInfo, // Assign the slice directly
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

func (server *Server) handlerUpdateRolePermission(w http.ResponseWriter, r *http.Request){
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

	permission := db.RolesPermission{}
	err = json.NewDecoder(r.Body).Decode(&permission)

	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid JSON request")
		return
	}

	arg := db.UpdateRolePermissionParams{
		ID:       	 int32(id),
		RoleID:      permission.RoleID,
		PermissionID: permission.PermissionID,
		IsDeleted: 	permission.IsDeleted,
	}

	rolepermissionsInfo,err:= server.store.UpdateRolePermission(ctx, arg)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Failed to fetch role permission")
		return
	}

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.RolesPermission `json:"data"`
	}{
		Status:  true,
		Message: "role permission updated successfully",
		Data:    []db.RolesPermission{rolepermissionsInfo},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerDeleteRolePermission(w http.ResponseWriter, r *http.Request) {
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

	rolepermissionsInfo, err:= server.store.DeleteRolePermission(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch role permission",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    []db.RolesPermission `json:"data"`
	}{
		Status:  true,
		Message: " role permission deleted successfully",
		Data:    []db.RolesPermission{rolepermissionsInfo},
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


