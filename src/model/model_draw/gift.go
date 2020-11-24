package model_draw

import "time"

type Gift struct {
	Id          string    `json:"id"`
	ActionID    string    `json:"actionId"`
	Enable      int       `json:"enable"`
	GiftType    int       `json:"giftType"`
	GiftName    string    `json:"giftName"`
	BeginTime   time.Time `json:"beginTime"`
	EndTime     time.Time `json:"endTime"`
	Odds        int       `json:"odds"`
	Count       int       `json:"count"`
	RemainCount int       `json:"remainCount"`
}
