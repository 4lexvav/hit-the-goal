package projects

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/handlers/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "projectID")
	if extErr != nil {
		return
	}

	extErr = services.Get().Projects().Delete(id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendRawResponse(w, http.StatusNoContent, []byte{})
}
