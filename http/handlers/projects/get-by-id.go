package projects

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "projectID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	project, extErr := services.Get().Projects().GetByID(id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, project)
}
