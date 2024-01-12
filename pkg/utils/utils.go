package utils

import (
	"encoding/json"
	"io"
	"net/http"
)


func ParseBody(r *http.Request, x interface{}){

	 if r.Header.Get("Content-Type") != "application/json" {
        // Handle non-JSON content type (e.g., return an error or ignore the request)
        return
    }
	if body, err := io.ReadAll(r.Body); err == nil {
		if err:= json.Unmarshal([]byte(body),x); err != nil {
			return
		}
	
	}
}