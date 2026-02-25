package handler

import (
	"errors"
	"net/http"

	"github.com/SaikatDeb12/storeX/internal/database/dbhelper"
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
	case "Mobile":
		err = dbhelper.InsertMobileDetails(assetID, *req.Mobile)
	default:
		err = errors.New("invalid asset type")
	}

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "")
		return
	}
}
