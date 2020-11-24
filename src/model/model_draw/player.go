package model_draw

import "time"

type Player struct {
	Id                 string
	ActionID           string
	PlayCount          int
	RemainPlayCount    int
	WinCount           int
	RemainWinCount     int
	DayPlayCount       int
	RemainDayPlayCount int
	DayWinCount        int
	RemainDayWinCount  int
	LastPlayTime       time.Time
}

func (m *Player) CanPlay() bool {
	return m.RemainDayPlayCount > 0 && m.RemainPlayCount > 0
}

func (m *Player) CanWin() bool {
	return m.RemainDayWinCount > 0 && m.RemainWinCount > 0
}
