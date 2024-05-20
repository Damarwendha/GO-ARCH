package controller

import (
	"database/sql"
	"go-arch/middleware"
	"go-arch/model/dto"
	"go-arch/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService    service.AuthServiceI
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddlewareI
}

func (ac *AuthController) loginHandler(c *gin.Context) {
	var payload dto.AuthReqDto
	if err := c.ShouldBind(&payload); err != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := ac.authService.Login(payload)
	if err == sql.ErrNoRows {
		dto.SendErrorResponse(c, http.StatusBadRequest, "invalid email")
		return
	}
	if err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	dto.SendSingleResponse(c, token, http.StatusCreated, "Created")
}

func (ac *AuthController) Routing() {
	ac.rg.POST("/auth/login", ac.loginHandler)
}

func NewAuthController(authService service.AuthServiceI, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddlewareI) *AuthController {
	return &AuthController{
		authService:    authService,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
