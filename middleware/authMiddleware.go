package middleware

import (
	"go-arch/model/dto"
	"go-arch/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareI interface {
	VerifyTokenAndRole(allowedRoles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtServiceI
}

// VerifyTokenAndRole implements AuthMiddlewareI.
func (a *authMiddleware) VerifyTokenAndRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			dto.SendErrorResponse(c, http.StatusUnauthorized, "Empty token")
			return
		}

		log.Println("header", authHeader)

		token := strings.Replace(authHeader, "Bearer ", "", -1)

		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			dto.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("author_id", claims["author_id"])

		// VALIDATE ROLE
		var isValidRole bool
		for _, role := range allowedRoles {
			if role == claims["role"] {
				isValidRole = true
				break
			}
		}

		if !isValidRole {
			dto.SendErrorResponse(c, http.StatusForbidden, "Forbidden")
			return
		}
	}
}

func NewAuthMiddleware(jwtService service.JwtServiceI) AuthMiddlewareI {
	return &authMiddleware{
		jwtService: jwtService,
	}
}
