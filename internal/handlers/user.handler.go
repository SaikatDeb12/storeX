package handler

import (
	"net/http"
	"strconv"

	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/database/dbhelper"
	"github.com/SaikatDeb12/storeX/internal/models"
	"github.com/SaikatDeb12/storeX/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	role := query.Get("role")
	employment := query.Get("employment")
	assetStatus := query.Get("status")

	limit := 10
	page := 1
	var err error

	if limitInput := query.Get("limit"); limitInput != "" {
		limit, err = strconv.Atoi(limitInput)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid limit input")
		}
	}

	if pageInput := query.Get("page"); pageInput != "" {
		page, err = strconv.Atoi(pageInput)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid page input")
		}
	}

	page = max(page, 1)
	offset := (page - 1) * limit

	userDetails, err := dbhelper.FetchUsers(name, role, employment, assetStatus, limit, offset)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch users")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string][]models.UserInfoRequest{
		"users": userDetails,
	})
}

func GetUserInfoByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	userDetails, err := dbhelper.FetchUserByID(userID)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, err, "failed to fetch user details")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]models.UserInfoRequest{
		"user": userDetails,
	})
}

func AssignRole(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "user id not provided")
		return
	}

	var req models.AssignRoleRequest
	if err := utils.ParseBody(r.Body, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	role := req.Role
	err := dbhelper.AssignUserRole(userID, role)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "error while assigning role")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "user role changed",
	})
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "user id not provided")
		return
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		err := dbhelper.UnassignAssets(tx, userID)
		if err != nil {
			return err
		}

		err = dbhelper.DeleteUserSession(tx, userID)
		if err != nil {
			return err
		}

		err = dbhelper.DeleteUser(tx, userID)
		if err != nil {
			return err
		}

		return err
	})
	if txErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, txErr, "failed to delete user")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"status": "user deleted successfully",
	})
}
