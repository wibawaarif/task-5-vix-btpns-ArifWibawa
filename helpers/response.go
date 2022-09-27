package helpers

import (
	"net/http"
	"encoding/json"
)


func ResponseJSON(x http.ResponseWriter, code int, payload interface{}) {
	x.Header().Add("Content-Type", "application/json")
	x.WriteHeader(code)
	json.NewEncoder(x).Encode(payload)

}