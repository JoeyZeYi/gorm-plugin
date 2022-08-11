package main

import (
	"context"
	"fmt"
	"gorm-plugin/dao"
	"gorm-plugin/model"
)

func main() {
	myDB := &dao.MyDB{}
	ctx := context.TODO()
	roleDao := dao.NewRoleDao(myDB)

	userDao := dao.NewUserDao(myDB)

	userInfo := new(model.UserInfo)
	role := new(model.Role)

	//开启事务
	err := userDao.Transaction(ctx, func(txCtx context.Context) error {

		if err := userDao.Insert(txCtx, userInfo); err != nil {
			return err
		}
		return roleDao.Insert(txCtx, role)
	})
	if err != nil {
		fmt.Println("sql执行错误", err)
		return
	}
	fmt.Println("sql执行成功")
}
