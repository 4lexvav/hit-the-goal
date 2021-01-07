package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/4lexvav/hit-the-goal/config"
	"github.com/4lexvav/hit-the-goal/logger"
	"github.com/4lexvav/hit-the-goal/validator"
	exterrors "github.com/eugeneradionov/ext-errors"
	v "github.com/go-playground/validator/v10"
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

func ProcessRequestBody(w http.ResponseWriter, r *http.Request, body interface{}) error {
	if extError := UnmarshalRequestBody(r, body); extError != nil {
		SendExtError(w, extError)
		return extError
	}

	if httpErrs := ValidateRequestBody(r, body); httpErrs != nil {
		SendExtErrors(w, http.StatusUnprocessableEntity, httpErrs)
		return httpErrs
	}

	return nil
}

func UnmarshalRequestBody(r *http.Request, body interface{}) exterrors.ExtError {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serverError := exterrors.NewInternalServerErrorError(errors.New("Failed to read JSON body"))
		logger.WithCtxValue(r.Context()).Errorw("Invalid request body", "error", err)
		return serverError
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, body)
	if err != nil {
		serverError := exterrors.NewInternalServerErrorError(errors.New("Failed to parse JSON body"))
		logger.WithCtxValue(r.Context()).Errorw("Invalid JSON request body", "error", err, "json", string(reqBody))
		return serverError
	}

	return nil
}

func ValidateRequestBody(r *http.Request, body interface{}) exterrors.ExtErrors {
	err := validator.Get().Struct(body)
	if err != nil {
		validationErrors := err.(v.ValidationErrors)
		serverError := validator.FormatErrors(validationErrors)
		logger.WithCtxValue(r.Context()).Infow("Request body validation failed", "error", err)
		return serverError
	}

	return nil
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
