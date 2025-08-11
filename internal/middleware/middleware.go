package middleware

import (
	"errors"
	"net/http"
	"project2/internal/configs"
	"project2/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	 secretKey := configs.GetConf().Service.SecretJWT
	 return func(c *gin.Context){
		header:= c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		UserID, username, err :=jwt.ValidateToken(header, secretKey)
		if err!=nil{
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("UserID", UserID)
		c.Set("username", username)
		c.Next() // lanjut ke middleware selanjutnya jika ada
	 }
}