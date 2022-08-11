package model

type UserInfo struct {
}

func (UserInfo) TableName() string {
	return "user_info"
}
