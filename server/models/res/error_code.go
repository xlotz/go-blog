package res

type ErrorCode int

/**
自定义错误码
*/

const (
	SettingsError ErrorCode = 1001 // 系统错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
	}
)
