package dto

type SignupReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"gte=0,lte=130"`
	Gender   string `json:"gender" validate:"oneof=male female"`
}

type SignupRes struct {
}
