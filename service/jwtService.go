package service

import (
	"go-arch/config"
	"go-arch/model"
	"go-arch/model/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServiceI interface {
	CreateToken(author model.Author) (dto.AuthRespDto, error)
	ValidateToken(token string) (jwt.MapClaims, error)
}

type jwtService struct {
	co config.TokenConfig
}

type CustomClaims struct {
	jwt.RegisteredClaims
	AuthorId string `json:"author_id"`
	Role     string `json:"role"`
}

// CreateToken implements JwtServiceI.
func (j *jwtService) CreateToken(author model.Author) (dto.AuthRespDto, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.co.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.co.ExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AuthorId: author.Id,
		Role:     author.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.co.SigningKey)
	if err != nil {
		return dto.AuthRespDto{}, err
	}

	return dto.AuthRespDto{Token: signedToken}, nil
}

// ValidateToken implements JwtServiceI.
func (j *jwtService) ValidateToken(token string) (jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.co.SigningKey), nil
	})

	if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func NewJwtService(co config.TokenConfig) JwtServiceI {
	return &jwtService{
		co: co,
	}
}
