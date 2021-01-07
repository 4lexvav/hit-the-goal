package projects

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/handlers/common"
	"github.com/4lexvav/hit-the-goal/handlers/projects/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var req requests.ProjectRequest
	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	project, extErr := services.Get().Projects().Create(req.Name, req.Description)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, project)
}
