package lists

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/lists/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var req requests.ListRequest
	id, extErr := common.GetID(r, "projectID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	proj, extErr := services.Get().Projects().GetByID(id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	position := req.Position
	if position == 0 {
		position, extErr = services.Get().Lists().GetMaxPosition(proj.ID)
		if extErr != nil {
			common.SendExtError(w, extErr)
			return
		}

		position = position + 1
	}

	list, extErr := services.Get().Lists().Create(proj.ID, req.Name, req.Status, position)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, list)
}
