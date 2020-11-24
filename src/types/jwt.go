package types

import "github.com/kataras/iris/v12/middleware/jwt"

type WxAuthInfoClaims struct {
	jwt.Claims
	WxAuthInfo
}

type BaseUserInfo struct {
	Id       string
	NickName string
	Avatar   string
	UserName string
	Sex      string
	Mobile   string
	WeiXin   string
	QQ       string
	EMail    string
	Address  string
	Street   string
	Roles    []string
}

type BaseUserInfoClaims struct {
	jwt.Claims
	BaseUserInfo
}
