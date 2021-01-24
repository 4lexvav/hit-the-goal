package comments

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/http/handlers/comments/requests"
	"github.com/4lexvav/hit-the-goal/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var req requests.CommentRequest
	id, extErr := common.GetID(r, "taskID")
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	if err := common.ProcessRequestBody(w, r, &req); err != nil {
		return
	}

	comment, extErr := services.Get().Comments().Create(int64(id), req.Text)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendResponse(w, http.StatusOK, comment)
}
