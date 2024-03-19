package weberror

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/FurmanovD/postpackage/internal/app/service"
	api "github.com/FurmanovD/postpackage/pkg/api/v1"
)

// a case of error wrapping.

const (
	NoError      = 0
	UnknownError = -1

	NotImplemented = 10
	NotFound       = 20
	InvalidRequest = 30
	AlreadyExist   = 40
	CannotCreate   = 50

	DatabaseError = 80

	InternalServerError = 1001
)

// just internal structure to return an error code and a json response
type codeAndResponse struct {
	code int
	resp *api.CommonResponse
}

var (
	// NOTE: Add new service errors to return/wrap out here:
	serviceErrorToWebResponse = map[error]codeAndResponse{
		error(nil): {
			http.StatusOK,
			api.GetCommonResponseOk(),
		},

		service.ErrNotFound: {
			http.StatusNotFound,
			api.GetCommonResponseError(NotFound, "not found"),
		},

		service.ErrTimeout: {
			http.StatusOK,
			api.GetCommonResponseError(NoError, "request context is cancelled"),
		},

		service.ErrNotImplemented: {
			http.StatusNotImplemented,
			api.GetCommonResponseError(NotImplemented, "not implemented"),
		},

		service.ErrAlreadyExists: {
			http.StatusNotAcceptable,
			api.GetCommonResponseError(AlreadyExist, "already exists"),
		},

		service.ErrDBError: {
			http.StatusInternalServerError,
			api.GetCommonResponseError(DatabaseError, service.ErrDBError.Error()),
		},

		service.ErrInternalServerError: {
			http.StatusInternalServerError,
			api.GetCommonResponseError(InternalServerError, service.ErrInternalServerError.Error()),
		},
	}
)

func GetWebResponse(serviceError error, details string) (int, *api.CommonResponse) {
	if serviceError == nil {
		return http.StatusOK, api.GetCommonResponseOk()
	}

	if serviceError == service.ErrInvalidRequest {
		return http.StatusBadRequest, api.GetCommonResponseError(
			InvalidRequest, details,
		)
	}

	if retval, ok := serviceErrorToWebResponse[serviceError]; ok {
		if details != "" {
			retval.resp.ErrorMsg += ": " + details
		}
		return retval.code, retval.resp
	}

	// default case:
	logrus.Errorf(
		"unknown service error %+v has no string representation",
		serviceError,
	)
	return http.StatusInternalServerError, api.GetCommonResponseError(
		UnknownError, "Unknown error",
	)
}
