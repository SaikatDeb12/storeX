package handler

import (
	"net/http"

	"github.com/SaikatDeb12/storeX/internal/database/dbhelper"
	"github.com/SaikatDeb12/storeX/internal/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// userCtx, _ := middleware.UserContext(r)
	// userID := userCtx.UserID
	userDetails, err := dbhelper.GetUserInfo()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch users")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"users": userDetails,
	})
}
