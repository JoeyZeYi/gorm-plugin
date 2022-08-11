package dao

import "gorm-plugin/model"

type IRoleDao interface {
	IBaseStore[model.Role]
}

type roleDao struct {
	*baseStore[model.Role]
}

func NewRoleDao(db IDB) IRoleDao {
	return &roleDao{baseStore: createStore[model.Role](db)}
}
