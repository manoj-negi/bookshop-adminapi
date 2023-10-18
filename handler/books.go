package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	// db "github.com/vod/db/sqlc"
	// util "github.com/vod/utils"
)

type CreateBookRequest struct {
	Title         string `json:"title" validate:"required"`
	AuthorID      int32  `json:"author_id" validate:"required"`
	Price         int32  `json:"price" validate:"required"`
	StockQuantity int32  `json:"stock_quantity" validate:"required"`
}

type BookResponse struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"status_code"`
}

/* type SuceessResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	UserToken interface{} `json:"userToken"`
} */

/*
	func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
			return
		}
		ctx := r.Context()

		user := User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

			jsonResponse := JsonResponse{
				Status:     false,
				Message:    "invalid JSON request",
				StatusCode: http.StatusTeapot,
			}
			util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
			return
		}

		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {

			jsonResponse := JsonResponse{
				Status:     false,
				Message:    "invalid JSON request",
				StatusCode: http.StatusTeapot,
			}
			util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
			return
		}

		arg := db.CreateUserParams{
			Username:       user.Username,
			HashedPassword: hashedPassword,
			FullName:       user.FullName,
			Email:          user.Email,
		}

		userInfo, err := server.store.CreateUser(ctx, arg)
		if err != nil {

			jsonResponse := JsonResponse{
				Status:     false,
				Message:    "invalid JSON request",
				StatusCode: http.StatusTeapot,
			}
			util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userInfo)
	}
*/
func (server *Server) handlerInsertBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}

	book := CreateBookRequest{}
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Something went wrong",
			StatusCode: http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(book)
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

	// dummy JSON response
	bookInfo := struct {
		Message string `json:"message"`
		UserID  int    `json:"user_id"`
	}{
		Message: "Book Added",
		UserID:  123, // Replace with an actual user ID if needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookInfo)
}

/* func (server *Server) refreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}

	// dummy JSON response
	userInfo := struct {
		Message string `json:"message"`
		UserID  int    `json:"user_id"`
	}{
		Message: "refresh token",
		UserID:  123, // Replace with an actual user ID if needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
} */
