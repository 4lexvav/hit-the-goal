package lists

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "listID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	list, extErr := services.Get().Lists().GetByID(id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, list)
}
