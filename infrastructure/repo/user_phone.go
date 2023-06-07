package repo

import (
	"schat/infrastructure/repo/data"
	"schat/types"
)

type userPhoneRepo struct {
}

func NewUserPhone() *userPhoneRepo {
	return &userPhoneRepo{}
}

func (usr *userPhoneRepo) Fetch(phone *types.Phone) data.UserPhone {

}

func (usr *userPhoneRepo) Save(data data.UserPhone) error {

}
