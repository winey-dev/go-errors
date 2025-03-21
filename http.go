package errors

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/winey-dev/go-errors/codes"
)

type HTTPErrorForm struct {
	ErrCode  codes.Code `json:"err_code"`
	Message  string     `json:"message"`
	Detailed []string   `json:"detailed"`
}

func HTTPError(err error) *HTTPErrorForm {
	var httpError *HTTPErrorForm
	var Error *Error
	if errors.As(err, &Error) { //
		httpError = &HTTPErrorForm{
			ErrCode:  Error.code,
			Message:  Error.message,
			Detailed: Error.Details(),
		}
	} else {
		httpError = &HTTPErrorForm{
			ErrCode: codes.Internal,
			Message: err.Error(),
		}
	}
	return httpError
}

func JSONHTTPErrorHandle(w http.ResponseWriter, err error) {
	httpError := HTTPError(err)
	dat, err := json.Marshal(httpError)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(dat), codes.ToHTTPStatus(httpError.ErrCode))
}
