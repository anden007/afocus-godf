package service_core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/anden007/afocus-godf/src/interfaces/interface_core"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/types"

	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	jsonTime "github.com/liamylian/jsontime/v2/v2"
	"github.com/spf13/cast"
)

type WxAuth struct {
	campaignId    string
	returnUrl     string
	authUrl       string
	authInfoUrl   string
	getInfoMode   string
	jwtInstance   interface_core.IJwtService
	cacheInstance interface_core.ICache
	authInfoCache *Cache
	json          jsoniter.API
}

func NewWxAuth(jwt interface_core.IJwtService, cache interface_core.ICache) *WxAuth {
	loadTime := time.Now()
	instance := new(WxAuth)
	enable := os.Getenv("wxauth-enable")
	if strings.EqualFold("true", enable) {
		instance.json = jsonTime.ConfigWithCustomTimeFormat
		instance.cacheInstance = cache
		instance.campaignId = os.Getenv("wxauth-campaignid")
		instance.returnUrl = os.Getenv("wxauth-return-url")
		instance.authUrl = os.Getenv("wxauth-auth-url")
		instance.authInfoUrl = os.Getenv("wxauth-auth-info-url")
		instance.jwtInstance = jwt
		instance.getInfoMode = os.Getenv("wxauth-get-info-mode")
		if strings.EqualFold("redis", instance.getInfoMode) {
			redis_db := cast.ToInt(os.Getenv("wxauth-redis-db"))
			redis_pool_size := cast.ToInt(os.Getenv("wxauth-redis-pool-size"))
			instance.authInfoCache = NewCacheByConfig(CacheConfig{
				AppName:  os.Getenv("wxauth-app-name"),
				DataBase: redis_db,
				Password: os.Getenv("wxauth-redis-password"),
				PoolSize: redis_pool_size,
				Server:   os.Getenv("wxauth-redis-server"),
			})
		}
		if lib.IS_DEV_MODE {
			fmt.Println("> Service: WxAuth loaded, ", time.Since(loadTime))
		}
	} else {
		fmt.Printf("> Service: WxAuth is disabled. if you need enable it,please set 'wxauth-enable' to 'true' in .env file.\n")
	}
	return instance
}

func (m *WxAuth) getAuthUrl() (result string) {
	result = fmt.Sprintf("%s%s", m.authUrl, m.campaignId)
	return
}

func (m *WxAuth) getAuthInfoUrl(authId string) (result string) {
	result = fmt.Sprintf("%s%s", m.authInfoUrl, authId)
	return
}

func (m *WxAuth) GetCurrentUser(ctx iris.Context) (result *types.WxAuthInfo) {
	if verifiedToken, err := m.jwtInstance.VerifyToken(ctx); err == nil {
		wxAuthInfoClaims := types.WxAuthInfoClaims{}
		if err := verifiedToken.Claims(&wxAuthInfoClaims); err == nil {
			cak := ctx.URLParam("cak")
			result = &wxAuthInfoClaims.WxAuthInfo
			result.CAK = cak
		}
	}
	return
}

func (m *WxAuth) CheckAuth(ctx iris.Context, returnUrl string) (result bool) {
	// todo:returnUrl尚未实现
	if m.GetCurrentUser(ctx) == nil {
		url := m.getAuthUrl()
		result = false
		ctx.Redirect(url)
	} else {
		result = true
	}
	return
}

func (m *WxAuth) DoAuth(ctx iris.Context) {
	wxAuthInfo := &types.WxAuthInfo{}
	authid := ctx.URLParam("authid")
	cak := ctx.URLParam("cak")
	if authid != "" {
		if strings.EqualFold("redis", m.getInfoMode) {
			// redis模式
			jsonStr, _ := m.authInfoCache.Get(authid)
			if err := m.json.UnmarshalFromString(jsonStr, &wxAuthInfo); err != nil {
				lib.LogCenter().Errorf("用户信息格式错误：%s", err.Error())
				_, _ = ctx.Text("用户信息格式错误：%s", err.Error())
			}
		} else {
			// http模式
			url := m.getAuthInfoUrl(authid)
			if resp, err := http.Get(url); err != nil {
				lib.LogCenter().Errorf("用户授权失败，请求：%s 错误：%s", url, err.Error())
				_, _ = ctx.Text("用户授权失败，错误：%s", err.Error())
			} else {
				defer resp.Body.Close()
				if body, err := ioutil.ReadAll(resp.Body); err == nil {
					if err := m.json.Unmarshal(body, &wxAuthInfo); err != nil {
						lib.LogCenter().Errorf("用户信息格式错误：%s", err.Error())
						_, _ = ctx.Text("用户信息格式错误：%s", err.Error())
					}
				}
			}
		}
		if wxAuthInfo != nil {

			claims := types.WxAuthInfoClaims{
				WxAuthInfo: *wxAuthInfo,
			}
			m.jwtInstance.CreateToken(ctx, claims)
			parmSplitor := "?"
			if strings.Index(m.returnUrl, "?") > -1 {
				parmSplitor = "&"
			}
			url := ""
			if cak != "" {
				url = fmt.Sprintf("%s%scak=%s", m.returnUrl, parmSplitor, cak)
			} else {
				url = m.returnUrl
			}
			ctx.Redirect(url)
		}
	}
}
