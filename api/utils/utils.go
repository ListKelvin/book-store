package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
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

var errorMessages = make(map[string]string)

var err error

func FormatError(errString string) map[string]string {

	if strings.Contains(errString, "username") {
		errorMessages["Taken_username"] = "Username Already Taken"
	}

	if strings.Contains(errString, "email") {
		errorMessages["Taken_email"] = "Email Already Taken"

	}
	if strings.Contains(errString, "title") {
		errorMessages["Taken_title"] = "Title Already Taken"

	}
	if strings.Contains(errString, "hashedPassword") {
		errorMessages["Incorrect_password"] = "Incorrect Password"
	}
	if strings.Contains(errString, "record not found") {
		errorMessages["No_record"] = "No Record Found"
	}

	if strings.Contains(errString, "double like") {
		errorMessages["Double_like"] = "You cannot like this post twice"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	if len(errorMessages) == 0 {
		errorMessages["Incorrect_details"] = "Incorrect Details"
		return errorMessages
	}

	return nil
}