package tasks

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "taskID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	task, extErr := services.Get().Tasks().GetByID(id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, task)
}
