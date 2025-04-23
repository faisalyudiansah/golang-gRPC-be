package middleware

import (
	"fmt"

	"strings"

	"server/pkg/apperror"
	utils "server/pkg/utils/context"
	"server/pkg/utils/jwtutils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtUtil jwtutils.JwtUtil
}

func NewAuthMiddleware(
	jwtUtil jwtutils.JwtUtil,
) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := m.parseAccessToken(ctx)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}
		fmt.Println(accessToken)
		ctx.Next()
	}
}

func (m *AuthMiddleware) ProtectedRoles(allowedRoles ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := utils.GetValueRoleUserFromToken(ctx)

		allowed := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				allowed = true
				break
			}
		}

		if !allowed {
			ctx.Error(apperror.NewForbiddenAccessError())
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (m *AuthMiddleware) parseAccessToken(ctx *gin.Context) (string, error) {
	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" || len(accessToken) == 0 {
		return "", apperror.NewUnauthorizedError()
	}

	splitToken := strings.Split(accessToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", apperror.NewUnauthorizedError()
	}
	return splitToken[1], nil
}
