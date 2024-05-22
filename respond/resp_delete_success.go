package respond

type RespDeleteSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
