package respond

type RespStoreSuccess struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
