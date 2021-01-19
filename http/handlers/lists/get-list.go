package lists

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "ProjectID")
	if extErr != nil {
		common.SendExtError(w, extErr)
	}

	size, page, extErr := common.GetSizeAndPage(r.URL.Query())
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	lists, extErr := services.Get().Lists().Get(id, size, page)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if len(lists) == 0 {
		common.SendResponse(w, http.StatusOK, []string{})
		return
	}

	common.SendResponse(w, http.StatusOK, lists)
}
