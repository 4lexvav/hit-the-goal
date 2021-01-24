package tasks

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/tasks/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var req requests.TaskRequest
	id, extErr := common.GetID(r, "listID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	task, extErr := services.Get().Tasks().Create(int64(id), req.Name, req.Description, req.Position)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, task)
}
