package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SaikatDeb12/storeX/internal/database/dbhelper"
	"github.com/SaikatDeb12/storeX/internal/middleware"
	"github.com/SaikatDeb12/storeX/internal/models"
	"github.com/SaikatDeb12/storeX/internal/utils"
)

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAssetRequest
	if err := utils.ParseBody(r.Body, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	if err := utils.ValidateStruct(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "payload validation failed")
		return
	}

	assetID, err := dbhelper.CreateAsset(req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create asset")
		return
	}

	assetType := req.Type
	switch assetType {
	case "laptop":
		err = dbhelper.InsertLaptopDetails(assetID, *req.Laptop)
	case "keyboard":
		err = dbhelper.InsertKeyboardDetails(assetID, *req.Keyboard)
	case "mouse":
		err = dbhelper.InsertMouseDetails(assetID, *req.Mouse)
	case "mobile":
		err = dbhelper.InsertMobileDetails(assetID, *req.Mobile)
	default:
		err = errors.New("invalid asset type")
	}

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "")
		return
	}
}

func ShowAssets(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	brand := query.Get("brand")
	model := query.Get("model")
	assetType := query.Get("type")
	status := query.Get("status")
	owner := query.Get("owner")
	serialNumber := query.Get("serial_number")

	limit := 10
	page := 1
	var err error

	if limitInput := query.Get("limit"); limitInput != "" {
		limit, err = strconv.Atoi(limitInput)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid limit")
			return
		}
	}

	if pageInput := query.Get("page"); pageInput != "" {
		page, err = strconv.Atoi(pageInput)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err, "invalid limit")
			return
		}
	}

	offset := (page - 1) * limit

	assets, err := dbhelper.ShowAssets(brand, model, assetType, serialNumber, status, owner, limit, offset)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch assets")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"assets": assets,
	})
}

func AssignedAssets(w http.ResponseWriter, r *http.Request) {
	var req models.AssetAssignRequest

	if parseErr := utils.ParseBody(r.Body, &req); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed parsing body")
		return
	}
	userCtx, _ := middleware.UserContext(r)
	currectUserID := userCtx.UserID

	err := dbhelper.AssignedAssets(req.AssetID, currectUserID, req.UserID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to assigned assets")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "successfully assigned",
	})
}
