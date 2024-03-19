package api

const (
	StatusOK    = "OK"
	StatusError = "error"
)

// CommonResponse used to indicate Ok/Error status and contains additional error details
// to inform a user.
type CommonResponse struct {
	Status    string `json:"status,omitempty"`
	ErrorCode int    `json:"code,omitempty"`
	ErrorMsg  string `json:"error,omitempty"`
}

// GetCommonResponseOk rerurns a successful CommonResponse
func GetCommonResponseOk() *CommonResponse {
	return &CommonResponse{
		Status: StatusOK,
	}
}

// GetCommonResponseError rerurns a failed CommonResponse
func GetCommonResponseError(errCode int, errMsg string) *CommonResponse {
	return &CommonResponse{
		Status:    StatusError,
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
	}
}
