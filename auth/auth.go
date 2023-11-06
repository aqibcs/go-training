package auth

import (
	"fmt"
	"go-training/consts"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return consts.SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {
			// w.WriteHeader(http.StatusUnauthorized)
			// w.Write([]byte("not authorized"))
		}
	})
}

// func BasicAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodOptions {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			Unauthorized(w)
// 			return
// 		}

// 		auth := strings.SplitN(authHeader, " ", 2)
// 		if len(auth) != 2 || auth[0] != "Basic" {
// 			Unauthorized(w)
// 			return
// 		}

// 		payload, err := base64.StdEncoding.DecodeString(auth[1])
// 		if err != nil {
// 			Unauthorized(w)
// 			return
// 		}

// 		pair := strings.SplitN(string(payload), ":", 2)
// 		if len(pair) != 2 || pair[0] != consts.USERNAME || pair[1] != consts.PASSWORD {
// 			Unauthorized(w)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// func Unauthorized(w http.ResponseWriter) {
// 	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
// 	w.WriteHeader(http.StatusUnauthorized)
// 	w.Write([]byte("Unauthorized access"))
// }
