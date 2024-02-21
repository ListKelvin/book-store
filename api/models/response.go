package models

type Response struct {
	Data    interface{} `json:"data"`    // Can hold any data type
	Message string      `json:"message"` // Optional message
	Status  string      `json:"status"`  // Success, error, etc.
}