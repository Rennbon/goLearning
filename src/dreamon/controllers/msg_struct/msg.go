package msg_struct

type msg struct {
	Status  ResultCode
	Message string
	Data    interface{}
}

// 请求结果
type ResultCode int

const (
	//成功
	Success ResultCode = 1
	Error   ResultCode = 2
)

func NewMsg(status ResultCode, message string, data interface{}) *msg {
	return &msg{status, message, data}
}
