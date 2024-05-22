package respond

type ErrorAuth struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      struct {
		Detail string `json:"detail"`
	} `json:"error"`
}
