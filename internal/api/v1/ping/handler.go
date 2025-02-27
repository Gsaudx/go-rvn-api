package ping

import (
	"encoding/json"
	"net/http"

	"github.com/gsaudx/go-rvn-api/internal/api/common"
)

// PingHandler handles the /ping endpoint request.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//Using standardized response
	resp := common.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Request successful",
		Data: map[string]string{
			"message": "pong",
		},
	}

	json.NewEncoder(w).Encode(resp)
}
