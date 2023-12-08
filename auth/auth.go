package auth

import (
	"fmt"
	"go-training/consts"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateJWt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(consts.SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenHeader := c.Request().Header.Get("Token")

		if tokenHeader == "" {
			return c.String(http.StatusUnauthorized, "not authorized")
		}

		token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("not authorized")
			}
			return consts.SECRET, nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "not authorized: "+err.Error())
		}

		if token.Valid {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "not authorized")
	}
}
