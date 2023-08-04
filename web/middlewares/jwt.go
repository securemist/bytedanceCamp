package middlewares

import (
	"bytedanceCamp/config"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Uuid     int64
	UserName string
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.GetConfig().Jwt.Secret),
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请先登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// CreateToken 生成token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	// 错误检查
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, TokenMalformed
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, TokenExpired
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, TokenNotValidYet
		} else {
			return nil, TokenInvalid
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
