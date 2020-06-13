package response

import (
	"encoding/json"
	"io"
	"net/http"
)

// send normal response to client
func SendNormalResponse(w http.ResponseWriter, r *http.Request, json_str string, status_code int) {
	w.WriteHeader(status_code)
	io.WriteString(w, json_str)
}

//send error response to client
func SendErrorResponse(w http.ResponseWriter, ep ErrorResponse) {
	w.WriteHeader(ep.HttpSC)

	resStr, _ := json.Marshal(&ep.Err)
	io.WriteString(w, string(resStr))
}
