package models

type ApiReturn struct {
	ErrCode int `json:"err_code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}


