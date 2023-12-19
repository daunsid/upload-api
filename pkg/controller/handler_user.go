package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daunsid/upload-api/internal/db"
	"github.com/daunsid/upload-api/pkg/util"
	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		util.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	_, err = apiConfig.DB.GetUser(r.Context(), params.UserName)
	if err == nil {
		util.RespondWithError(w, 400, "Username already exists")
		return
	}

	hashedPassword, err := util.HashPassword(params.PassWord)
	if err != nil {
		util.RespondWithError(w, 500, fmt.Sprintf("Error hashing password: %v\n", err))
	}
	users, err := apiConfig.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:           uuid.New(),
		UserName:     params.UserName,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		util.RespondWithError(w, 400, fmt.Sprintf("Could'nt create user: %v", err))
		return
	}
	UserResponse := struct {
		UserID   uuid.UUID `json:"user_id"`
		UserName string    `json:"username"`
	}{

		UserID:   users.ID,
		UserName: users.UserName,
	}
	util.RespondWithJson(w, 200, UserResponse)
}

func (apiConfig *ApiConfig) HandlerLogin() {

}
