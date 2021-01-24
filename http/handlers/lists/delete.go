package lists

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "listID")
	if extErr != nil {
		return
	}

	projectID, extErr := common.GetID(r, "projectID")
	if extErr != nil {
		return
	}

	extErr = services.Get().Lists().Delete(projectID, id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendRawResponse(w, http.StatusNoContent, []byte{})
}
