package service

import (
	"go-arch/model/dto"
)

type authService struct {
	jwtService    JwtServiceI
	authorService AuthorServiceI
}

// Login implements AuthServiceI.
func (a *authService) Login(payload dto.AuthReqDto) (dto.AuthRespDto, error) {
	author, err := a.authorService.FindByEmail(payload.Email)

	if err != nil {
		return dto.AuthRespDto{}, err
	}

	token, err := a.jwtService.CreateToken(author)
	if err != nil {
		return dto.AuthRespDto{}, err
	}

	return token, nil
}

type AuthServiceI interface {
	Login(payload dto.AuthReqDto) (dto.AuthRespDto, error)
}

func NewAuthService(jwtService JwtServiceI, authorService AuthorServiceI) AuthServiceI {
	return &authService{
		jwtService:    jwtService,
		authorService: authorService,
	}
}
