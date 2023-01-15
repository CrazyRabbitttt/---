package Common

import "errors"

var (
	ErrorUserNameNull     = errors.New("用户名为空")
	ErrorUserNameNotValid = errors.New("用户名格式不符合规范")
	ErrorPasswordNull     = errors.New("密码为空")
	ErrorPasswordNotValid = errors.New("密码格式不符合规范")
	ErrorCreateUser       = errors.New("创建新用户失败")
	ErrorUserExits        = errors.New("注册的用户已经是存在的")
	ErrorUserNotExit      = errors.New("登陆的用户不存在")
)
