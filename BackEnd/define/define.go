package define

import "time"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	DefaultEmailCodeDuration   = time.Minute * 5
	DefaultTeacherRegisterCode = "123456"
	DefaultTeacherCode         = "0"
	DefaultStudentCode         = "1"
	DefaultUserStatus          = DefaultStudentCode
)
