package middleware

import (
	"net/http"

	"github.com/OctavianoRyan25/TODO-List/helper"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)
		_ = verifyToken
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"	: "Unauthenticated",
				"message" : err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}