package service

import (
	"schat/domain/entity"
	"schat/types"
)

type userService struct {
}

func NewUser() *userService {
	return &userService{}
}

func (us *userService) UserExist(phone *types.Phone) (exist bool, err error) {
	userEntity := entity.NewUser(phone)

	userEntity.Exist()
}
