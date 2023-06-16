package errors

var (
	ErrTokenType               = New(40002, "token 类型错误")
	ErrTokenParse              = New(40003, "token 解析失败")
	ErrTokenExp                = New(40004, "token 过期")
	ErrUsernameExist           = New(40005, "用户名已存在")
	ErrWrongUsernameOrPassword = New(40006, "用户名或密码错误")
)
