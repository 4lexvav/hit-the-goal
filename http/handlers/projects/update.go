package projects

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/projects/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "projectID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	var req requests.ProjectRequest
	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	project, extErr := services.Get().Projects().Update(id, req.Name, req.Description)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, project)
}
