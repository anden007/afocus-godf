package model_draw

type Result struct {
	ActionID string
	PlayerID string
	GiftType int
	GiftID   string
	GiftName string
	Voucher  *Voucher
	Player   *Player
	Success  bool
	Message  string
}
