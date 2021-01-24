package comments

import (
	"net/http"

	"github.com/4lexvav/hit-the-goal/http/common"
	"github.com/4lexvav/hit-the-goal/services"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, extErr := common.GetID(r, "commentID")
	if extErr != nil {
		return
	}

	extErr = services.Get().Comments().Delete(id)
	if extErr != nil {
		common.SendExtError(w, extErr)
		return
	}

	common.SendRawResponse(w, http.StatusNoContent, []byte{})
}
