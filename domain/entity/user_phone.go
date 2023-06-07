package entity

import (
	"schat/types"
)

type userPhoneEntity struct {
	phone *types.Phone
	uid   types.UserId
}

func NewUserPhone(phone *types.Phone) *userPhoneEntity {
	return &userPhoneEntity{}
}

func (ue *userPhoneEntity) fetchUserIdIfNeed() {
	if ue.uid == 0 {

	}
}

func (ue *userPhoneEntity) Exist(phone *types.Phone) (bool, error) {
	ue.fetchUserIdIfNeed()
}
