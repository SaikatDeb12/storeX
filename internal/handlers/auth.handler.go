package handler

import (
	"net/http"

	"github.com/SaikatDeb12/storeX/internal/database/dbhelper"
	"github.com/SaikatDeb12/storeX/internal/models"
	"github.com/SaikatDeb12/storeX/internal/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := utils.ParseBody(r.Body, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	if err := utils.ValidateStruct(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "payload validation error")
		return
	}

	if err := dbhelper.CheckUserExistsByEmail(req.Email); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "user already exists")
		return
	}

	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "password hashing failed")
		return
	}

	userID, err := dbhelper.CreateUser(req.Name, req.Email, req.PhoneNumber, req.UserRole, req.UserType, hashedPassword)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}

	sessionID, err := dbhelper.CreateSession(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create session")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, map[string]string{
		"message": "user register successfully",
		"token":   sessionID,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := utils.ParseBody(r.Body, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	if err := utils.ValidateStruct(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	user, err := dbhelper.GetUserAuthByEmail(req.Email)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid credentials")
		return
	}

	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid credentials")
		return
	}

	userID := user.ID
	sessionID, err := dbhelper.CreateSession(userID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create session")
		return
	}

	token, err := utils.GenerateJWT(userID, sessionID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "error while generating token")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "login successfull",
		"token":   token,
	})
}
