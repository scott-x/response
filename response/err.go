package response

type Err struct {
	Msg       string `json:"error_msg"`
	InnerCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Err
}

var (
	ErrorParseBodyFailed  = ErrorResponse{400, Err{"Request body is not correct", "001"}}
	ErrorAuthUserFailed   = ErrorResponse{401, Err{"User authentication failed.", "002"}}
	ErrorPermissionDenied = ErrorResponse{403, Err{"permission denied", "003"}}
	ErrorDBError          = ErrorResponse{500, Err{"DB error", "004"}}
	ErrorInternalFaults   = ErrorResponse{500, Err{"Internal service error", "005"}}
)
