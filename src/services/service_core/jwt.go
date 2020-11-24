package service_core

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/types"

	"github.com/spf13/cast"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type JWTService struct {
	jwtHeader        string
	jwtSecretKey     string
	jwtHeaderType    string
	jwtMaxAgeSeconds int
	jwtSigner        *jwt.Signer
	jwtVerifier      *jwt.Verifier
}

func NewJWT() *JWTService {
	return NewCustomJWT("")
}

func NewCustomJWT(header string) *JWTService {
	instance := new(JWTService)
	if header == "" {
		header = os.Getenv("jwt-header")
	}
	instance.jwtHeader = header
	instance.jwtSecretKey = os.Getenv("jwt-secretkey")
	instance.jwtHeaderType = os.Getenv("jwt-headertype")
	instance.jwtMaxAgeSeconds = cast.ToInt(os.Getenv("jwt-max-age-seconds"))
	instance.jwtSigner = jwt.NewSigner(jwt.HS256, instance.jwtSecretKey, time.Second*time.Duration(instance.jwtMaxAgeSeconds))
	instance.jwtVerifier = jwt.NewVerifier(jwt.HS256, instance.jwtSecretKey, jwt.Expected{
		Issuer: lib.APP_NAME,
	})
	instance.jwtVerifier.WithDefaultBlocklist()
	instance.jwtVerifier.Extractors = append(instance.jwtVerifier.Extractors, instance.FromCookie)
	return instance
}

// CreateToken 生成Token，默认存入Cookie中
func (m *JWTService) CreateToken(ctx iris.Context, claims interface{}) (token string, err error) {
	/* 这里省去了对用户的验证，在实际使用过程中需要验证用户是否存在，密码是否正确 */
	if tokenData, err := m.jwtSigner.Sign(claims, jwt.Claims{
		ID:     lib.NewGuidString(),
		Issuer: lib.APP_NAME,
	}); err == nil {
		token = string(tokenData)
		ctx.SetCookieKV(m.jwtHeader, fmt.Sprintf("%s %s", m.jwtHeaderType, token), iris.CookieHTTPOnly(false), iris.CookieExpires(time.Second*time.Duration(m.jwtMaxAgeSeconds)))
	}
	return
}

// VerifyToken 从默认渠道(Cookie)中获取Token信息并验证，验证成功后返回VerifiedToken对象，仅用于手动验证，不建议使用
func (m *JWTService) VerifyToken(ctx iris.Context) (result *jwt.VerifiedToken, err error) {
	token_str := m.FromCookie(ctx)
	result, err = m.jwtVerifier.VerifyToken([]byte(token_str))
	return
}

// RemoveToken 从Cookie中移除Token信息，一般用于"注销"功能
func (m *JWTService) RemoveToken(ctx iris.Context) {
	ctx.Logout()
	ctx.RemoveCookie(m.jwtHeader)
}

// GetClaims 验证Token并返回其中包含的Claims对象指针，需要在Middleware处理后才能获得，如需手动验证，请使用VerifyToken方法，强烈建议使用此方法
func (m *JWTService) GetClaims(ctx iris.Context) (claims interface{}) {
	claims = jwt.Get(ctx)
	return
}

// GetMiddleware 获取iris路由中间件，用于验证请求合法性
func (m *JWTService) GetMiddleware() iris.Handler {
	return m.jwtVerifier.Verify(func() interface{} {
		return new(types.BaseUserInfoClaims)
	})
}

func (m *JWTService) FromCookie(ctx iris.Context) string {
	tokenCookie := ctx.GetCookie(m.jwtHeader)
	if tokenCookie == "" {
		return ""
	}

	tokenCookieParts := strings.Split(tokenCookie, " ")
	if len(tokenCookieParts) != 2 || !strings.EqualFold(tokenCookieParts[0], m.jwtHeaderType) {
		return ""
	}

	return tokenCookieParts[1]
}
