package define

import "time"

// 成功代码200，数据库错误401，token错误402,接受数据错误403

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Accept struct {
}

const (
	DefaultEmailCodeDuration   = time.Minute * 5
	DefaultTeacherRegisterCode = "123456"
	DefaultTeacherCodeStr      = "2"
	DefaultTeacherCode         = 2
	DefaultStudentCodeStr      = "1"
	DefaultStudentCode         = 1
	DefaultUserStatusStr       = DefaultStudentCodeStr
	DefaultPage                = "1"
	DefaultPageSize            = "10"
)
