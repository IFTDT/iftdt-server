package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/common"
)

var JWTKET = []byte(common.ENV.JWTKey)

type MyClaims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

// 生成 token
func SetToken(id uint) (string, int) {
	expireTime := time.Now().Add(8 * time.Hour)
	claims := MyClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iftdt",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaim.SignedString(JWTKET)
	if err != nil {
		return "", 400
	}
	fmt.Println(token)
	return token, 200
}

// 验证 token
func ValidToken(token string) (*MyClaims, int) {
	t, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTKET, nil
	})
	if key, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return key, 200
	}
	return nil, 401
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizathion := c.Request.Header.Get("Authorization")
		if authorizathion == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
			})
			c.Abort()
			return
		}

		token := strings.SplitN(authorizathion, " ", 2)
		if len(token) != 2 && token[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
			})
			c.Abort()
			return
		}
		key, code := ValidToken(token[1])
		if code == 401 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
			})
			c.Abort()
			return
		}
		c.Set("id", key.Id)
		c.Next()
	}
}
