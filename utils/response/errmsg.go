package response

const (
	SUCCESS = 0
	ERROR   = 500

	NO_TOKEN            = 300
	TOEKEN_FORMAT_ERROR = 301
	TOKEN_ERROR         = 302
	TOKEN_TIMEOUT       = 303

	PERMISSION_DENIED = 350

	USER_NOT_EXIST = 400
)

var codeMsg = map[int]string{
	SUCCESS: "成功",
	ERROR:   "失败",

	NO_TOKEN:            "没有Token",
	TOEKEN_FORMAT_ERROR: "Token格式错误",
	TOKEN_ERROR:         "Token错误",
	TOKEN_TIMEOUT:       "Token过期",

	PERMISSION_DENIED: "权限不足",
	USER_NOT_EXIST:    "用户不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
