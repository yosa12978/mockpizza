package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yosa12978/mockPizza/internal/domain"
	"github.com/yosa12978/mockPizza/internal/helpers"
)

func UseJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawheader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(rawheader, "Bearer ") {
			ctx.JSON(401, gin.H{
				"message": "authorization header must begin with \"Bearer \" prefix",
			})
			ctx.Abort()
			return
		}
		token := strings.Replace(rawheader, "Bearer ", "", 1)
		claims, err := helpers.ParseJWT(token)
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		mapclaims := claims.(jwt.MapClaims)
		ctx.Set("usr", domain.Usr{
			BaseModel: domain.BaseModel{ID: mapclaims["id"].(uint)},
			Username:  mapclaims["username"].(string),
			Role:      mapclaims["role"].(string),
		})
	}
}
