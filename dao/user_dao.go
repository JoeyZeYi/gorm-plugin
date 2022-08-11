package dao

import "gorm-plugin/model"

type IUserDao interface {
	IBaseStore[model.UserInfo]
}

type userDao struct {
	*baseStore[model.UserInfo]
}

func NewUserDao(db IDB) IUserDao {
	return &userDao{baseStore: createStore[model.UserInfo](db)}
}
