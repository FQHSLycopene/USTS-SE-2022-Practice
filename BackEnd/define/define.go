package define

import "time"

// 成功代码200，数据库错误401，token错误402

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	DefaultEmailCodeDuration   = time.Minute * 5
	DefaultTeacherRegisterCode = "123456"
	DefaultTeacherCodeStr      = "0"
	DefaultTeacherCode         = 0
	DefaultStudentCodeStr      = "1"
	DefaultStudentCode         = 1
	DefaultUserStatusStr       = DefaultStudentCodeStr
	DefaultPage                = "1"
	DefaultPageSize            = "10"
)
