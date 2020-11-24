package model

var ModelDBFields map[string]map[string]string = make(map[string]map[string]string)

type DBModel struct {
}

func (DBModel) TableName() string {
	return ""
}
