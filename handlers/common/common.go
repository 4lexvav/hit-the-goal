package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/logger"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func SendResponse(w http.ResponseWriter, statusCode int, respBody interface{}) {
	binRespBody, err := json.Marshal(respBody)
	if err != nil {
		logger.Get().Errorw("Failed to marshal response body to json", "error", err)
		statusCode = http.StatusInternalServerError
	}

	SendRawResponse(w, statusCode, binRespBody)
}

func SendRawResponse(w http.ResponseWriter, statusCode int, binBody []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(statusCode)

	if _, err := w.Write(binBody); err != nil {
		logger.Get().Errorw("Failed to write response body", "error", err)
	}
}

func SendExtError(w http.ResponseWriter, extError exterrors.ExtError) {
	errs := exterrors.NewExtErrors()
	errs.Add(extError)

	SendExtErrors(w, extError.HTTPCode(), errs)
}

func SendExtErrors(w http.ResponseWriter, httpCode int, httpErrors exterrors.ExtErrors) {
	SendResponse(w, httpCode, httpErrors)
}

func GetSizeAndPage(urlQuery url.Values) (_, _ int, extErr exterrors.ExtError) {
	sizeStr := urlQuery.Get("size")
	size, err := strconv.ParseInt(sizeStr, 0, 64)

	if err != nil {
		extErr = exterrors.NewBadRequestError(errors.Wrap(err, "failed to parse 'size' from url"))
		return
	}

	if size <= 0 {
		extErr = exterrors.NewBadRequestError(errors.New("'size' must be greater than 0"))
		return
	}

	if size > config.Get().PaginationMaxLimit {
		extErr = exterrors.NewBadRequestError(fmt.Errorf("'size' must be less than %d", config.Get().PaginationMaxLimit))
		return
	}

	pageStr := urlQuery.Get("page")
	page, err := strconv.ParseInt(pageStr, 0, 64)
	if err != nil {
		extErr = exterrors.NewBadRequestError(errors.Wrap(err, "failed to parse 'page' from url"))
		return
	}

	if page < 0 {
		extErr = exterrors.NewBadRequestError(errors.New("'page' must be greater or equal to 0"))
		return
	}

	return int(size), int(page), nil
}
