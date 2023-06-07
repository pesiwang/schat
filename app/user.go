package app

import (
	"schat/domain/service"
	"schat/types"
)

type userApp struct {
}

func NewUser() *userApp {
	return &userApp{}
}

func (ua *userApp) CheckWhetherRegister(phone *types.Phone) (exist bool, err error) {
	userService := service.NewUser()
	exist, err = userService.UserExist(phone)

	return
}
