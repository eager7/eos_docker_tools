package code

type Resp struct {
	ErrCode ErrCode     `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Payload interface{} `json:"payload"`
}
