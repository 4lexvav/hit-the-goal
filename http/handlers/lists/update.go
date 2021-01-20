package lists

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/lists/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var req requests.ListRequest
	id, extErr := common.GetID(r, "listID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	list, extErr := services.Get().Lists().Update(id, req.Name, req.Status, req.Position)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, list)
}
