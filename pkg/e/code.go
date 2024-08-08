package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
)

var MsgFalgs = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFalgs[code]
	if ok {
		return msg
	}
	return MsgFalgs[Error]
}
