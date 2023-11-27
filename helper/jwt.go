package helper

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("SECREET_KEY")

func GenerateToken(id uint,email string) string {

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"id"	: id,
			"email"	: email,
			"exp"	: time.Now().Add(time.Minute * 1).Unix(),
		})

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error)  {
	errResponse := errors.New("sign into proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token ,_ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	claim, ok := token.Claims.(jwt.MapClaims);

	if  !ok && !token.Valid{
		return nil, errResponse
	}

	expirationTime := time.Unix(int64(claim["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return nil, errors.New("token expired")
	}

	return token.Claims.(jwt.MapClaims), nil
}
