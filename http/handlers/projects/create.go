package projects

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/projects/requests"
	"github.com/4lexvav/hit-the-goal/logger"
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/services"
)

const defaultListName string = "TODO"

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

	if _, extErrList := services.Get().Lists().Create(project.ID, defaultListName, models.ListStatusActive, 0); extErrList != nil {
		logger.Get().Errorw("Couldn't create default list while creating project", "project_id", project.ID, "error", extErrList)
	}

	common.SendResponse(w, http.StatusOK, project)
}
