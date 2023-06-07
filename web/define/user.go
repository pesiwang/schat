package define

import (
	"schat/types"
	"time"
)

// -------- request type ---------------------
type CheckWhetherRegisterRequest struct {
	types.Phone
}

type RegisterRequest struct {
	types.Phone
	AuthCode string    `json:"auth_code" form:"auth_code" binding:"len=6"`
	Gender   uint8     `json:"gender" form:"gender" binding:"min=0,max=1"`
	Birthday time.Time `json:"birthday" form:"birthday" binding:"required" time_format:"2006-01-02"`
}

type LoginRequest struct {
	types.Phone
	AuthCode string `json:"auth_code" form:"auth_code" binding:"len=6"`
}

// -------- response type --------------------

type AuthResponse struct {
	TokenString string `json:"token"`
}

type CheckWhetherRegisterResponse struct {
	Registered bool `json:"registered"`
}
