package model_draw

import "time"

type Action struct {
	Id              string    `json:"id"`
	Enable          int       `json:"enable"`
	Name            string    `json:"name"`
	BeginTime       time.Time `json:"beginTime"`
	EndTime         time.Time `json:"endTime"`
	MaxPlayCount    int       `json:"maxPlayCount"`
	MaxWinCount     int       `json:"maxWinCount"`
	DayMaxPlayCount int       `json:"dayMaxPlayCount"`
	DayMaxWinCount  int       `json:"dayMaxWinCount"`
	ConfirmSeconds  int       `json:"confirmSeconds"`
}
