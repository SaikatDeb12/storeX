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
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := utils.ParseBody(r.Body, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	// validator
}
