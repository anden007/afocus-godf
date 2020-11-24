package interface_core

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/jwt"
)

type IJwtService interface {
	// CreateToken 生成Token，默认存入Cookie中
	CreateToken(ctx iris.Context, claims interface{}) (token string, err error)
	// GetClaims 验证Token并返回其中包含的Claims对象指针，需要在Middleware处理后才能获得，如需手动验证，请使用VerifyToken方法，强烈建议使用此方法
	GetClaims(ctx iris.Context) (claims interface{})
	// VerifyToken 从默认渠道(Cookie)中获取Token信息并验证，验证成功后返回VerifiedToken对象，仅用于手动验证，不建议使用
	VerifyToken(ctx iris.Context) (result *jwt.VerifiedToken, err error)
	// RemoveToken 从Cookie中移除Token信息，一般用于"注销"功能
	RemoveToken(ctx iris.Context)
	// GetMiddleware 获取iris路由中间件，用于验证请求合法性
	GetMiddleware() iris.Handler
}
