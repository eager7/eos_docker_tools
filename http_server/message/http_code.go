package message

type Resp struct {
	Code    Code        `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}
