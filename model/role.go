package model

type Role struct {
}

func (Role) TableName() string {
	return "role"
}
