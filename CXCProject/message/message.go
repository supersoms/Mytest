package message

const (
	USER_OAUTH_LOGIN_SUCCESS_CODE          = 200 //用户授权登陆成功状态码
	USER_OAUTH_LOGIN_FAIL_CODE             = 201 //用户授权登陆失败状态码
	USER_BETTING_FAIL_CODE                 = 201 //用户投注失败状态码
	USER_BETTING_SUCCESS_CODE              = 200 //用户投注成功状态码
	USER_LOGIN_FAIL_CODE                   = 202 //用户登陆失败
	USER_OAUTH_LOGIN_SUCCESS_MSG           = "用户授权登陆成功!"
	USER_REGISTER_FAIL_MSG                 = "用户注册失败!"
	USER_REGISTER_FAIL_EXISTED_MSG         = "用户注册失败，该用户已存在!"
	USER_NOT_EXIST_MSG                     = "用户不存在，请注册!"
	LIVEBETTING_SUCCESS_MSG                = "用户实时投注成功!"
	LIVEBETTING_FAILED_MSG                 = "用户实时投注失败!"
	USER_SELECT_BET_FAIL_MSG               = "用户选号下注失败!"
	GET_USER_BETTING_RECORDING_SUCCESS_MSG = "获取用户投注记录数据成功!"
	GET_USER_BETTING_RECORDING_FAILED_MSG  = "获取用户投注记录数据失败!"
	USER_SELECT_BET_SUCCESS_MSG            = "用户选号下注成功!"
)

type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

type LoginInfo struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResult struct {
	Code   int `json:"code"` //状态码 500表示：该用户未注册，200表示：登陆成功
	UserId []int
	Error  string `json:"error"` //错误信息
}

type RegisterResult struct {
	Code  int    `json:"code"` //状态码 400表示：该用户已存在，200表示：注册成功
	Error string `json:"error"`
}
