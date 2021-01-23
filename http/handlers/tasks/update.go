package tasks

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/tasks/requests"
	"github.com/4lexvav/hit-the-goal/services"
	exterrors "github.com/eugeneradionov/ext-errors"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var req requests.TaskRequest
	id, extErr := common.GetID(r, "taskID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	listID, extErr := getListID(r, req)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	task, extErr := services.Get().Tasks().Update(id, req.Name, req.Description, req.Position, listID)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, task)
}

func getListID(r *http.Request, req requests.TaskRequest) (listID int64, extErr exterrors.ExtError) {
	id, extErr := common.GetID(r, "listID")
	if extErr != nil {
		return
	}

	listID = req.NewListID
	if req.NewListID == 0 {
		listID = int64(id)
	}

	return listID, nil
}
