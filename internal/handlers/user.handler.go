package handler

import (
	"net/http"

	"github.com/SaikatDeb12/storeX/internal/database/dbhelper"
	"github.com/SaikatDeb12/storeX/internal/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	role := query.Get("role")
	employment := query.Get("employment")
	assetStatus := query.Get("status")

	// userCtx, _ := middleware.UserContext(r)
	// userID := userCtx.UserID

	userDetails, err := dbhelper.GetUserInfo(name, role, employment, assetStatus)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch users")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"users": userDetails,
	})
}
