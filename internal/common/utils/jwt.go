package utils

import (
	"boilerplate-api/internal/common/context"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

type UserContext struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	jwt.RegisteredClaims
}

func generateToken(
	userId string,
	nickname string,
	avatar string,
	expirationTime int,
	secret string,
) (string, error) {
	claims := &UserContext{
		UserID:   userId,
		Nickname: nickname,
		Avatar:   avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationTime) * time.Second)),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateAccessToken(
	userId string,
	nickname string,
	avatar string,
	ctx *context.AppContext,
) (string, error) {
	return generateToken(userId, nickname, avatar, ctx.Cfg.JWT.AccessTokenExpirationTime, ctx.Cfg.JWT.Secret)
}

func GenerateRefreshToken(
	userId string,
	nickname string,
	avatar string,
	ctx *context.AppContext,
) (string, error) {
	return generateToken(userId, nickname, avatar, ctx.Cfg.JWT.RefreshTokenExpirationTime, ctx.Cfg.JWT.Secret)
}

func ValidateToken(
	tokenString string,
	ctx *context.AppContext,
) (*UserContext, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserContext{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(ctx.Cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserContext)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// no need to manually check token expiration , lib auto handle this
	//if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
	//	return nil, errors.ErrTokenExpired
	//}

	return claims, nil
}

const (
	issuer  = "boilerplate"
	UserCtx = "user_context"
)

func GetUserContext(c *gin.Context) *UserContext {
	val, exist := c.Get(UserCtx)
	if !exist {
		return nil
	}
	userCtx, ok := val.(*UserContext)
	if !ok {
		return nil
	}
	return userCtx
}
