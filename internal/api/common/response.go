package common

// Response is the standard response format for all API endpoints.
type Response struct {
	Status  string      `json:"status"`            // "success", "error", or "fail"
	Code    int         `json:"code"`              // HTTP status code, e.g., 200, 400, etc.
	Message string      `json:"message,omitempty"` // Optional human-readable message
	Data    interface{} `json:"data,omitempty"`    // Payload containing the data (if any)
	Error   string      `json:"error,omitempty"`   // Error details (if any)
}
