package types

type Response struct {
  Status  int     `json:"status"`
  Message string  `json:"message"`
  Body    interface{} `json:"body,omitempty"` // Empty if no body needed
}