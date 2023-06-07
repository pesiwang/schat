package types

type UserId uint64

type Phone struct {
	CountryCode string `json:"phone_country_code" form:"phone_country_code" binding:"min=2,max=5"`
	Number      string `json:"phone_number" form:"phone_number" binding:"min=5,max=15"`
}
