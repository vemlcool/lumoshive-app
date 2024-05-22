package respond

type ErrorPost struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
