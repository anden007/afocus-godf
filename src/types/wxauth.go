package types

import "time"

type WxAuthInfo struct {
	ID         string    `json:"id"`
	CAK        string    `json:"cak"`
	AppID      string    `json:"appId"`
	OpenID     string    `json:"openId"`
	UnionId    string    `json:"unionId"`
	NickName   string    `json:"nickName"`
	HeadImgUrl string    `json:"headImgUrl"`
	Sex        string    `json:"sex"`
	Country    string    `json:"country"`
	Province   string    `json:"province"`
	City       string    `json:"city"`
	Subscribe  int       `json:"subscribe"`
	CreateTime time.Time `json:"createTime" time_format:"2006-01-02 15:04:05"`
}
