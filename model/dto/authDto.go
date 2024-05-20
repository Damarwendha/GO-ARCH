package dto

type AuthRespDto struct {
	Token string `json:"token"`
}

type AuthReqDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}