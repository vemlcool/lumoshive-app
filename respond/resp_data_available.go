package respond

import (
	"myproject/model"
)

type RespDataAvailable struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       struct {
		model.User
	}
}

//belum selesai
