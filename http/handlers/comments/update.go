package comments

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/comments/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var req requests.CommentRequest
	id, extErr := common.GetID(r, "commentID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	task, extErr := services.Get().Comments().Update(id, req.Text)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, task)
}
