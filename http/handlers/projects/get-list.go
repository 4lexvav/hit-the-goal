package projects

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var query = r.URL.Query()

	size, page, extError := common.GetSizeAndPage(query)
	if extError != nil {
		common.SendExtError(w, extError)
		return
	}

	projects, extError := services.Get().Projects().Get(size, page)
	if extError != nil {
		common.SendExtError(w, extError)
		return
	}

	if len(projects) == 0 {
		common.SendResponse(w, http.StatusOK, []string{})
		return
	}

	common.SendResponse(w, http.StatusOK, projects)
}
