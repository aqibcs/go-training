package auth

import (
	"go-training/consts"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// ValidateJWT is a middleware that validates the presence and authenticity of a JWT in the request header.
// If the token is valid, the next handler in the chain is called. Otherwise, it returns an unauthorized response.
func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check if the "Token" header is present in the request
		if tokenStr := c.Request().Header.Get("Token"); tokenStr != "" {
			// Parse and validate the JWT from the "Token" header
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				// Check if the signing method is HMAC
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Not authorized: Invalid signing method")
				}
				return consts.Secret, nil
			})

			// Handle parsing and validation errors
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized: "+err.Error())
			}

			// Check if the token is valid
			if token.Valid {
				// Call the next handler in the chain
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized: Invalid token")
			}
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized: Token not provided")
		}
	}
}

// CreateJWT generates a new JWT with a predefined expiration time using the secret key.
// It returns the generated JWT as a string or an error if the generation fails.
func CreateJWT() (string, error) {
	// Create a new JWT token with the HS256 signing method
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token, including the expiration time
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	// Sign the token with the secret key
	tokenStr, err := token.SignedString(consts.Secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
