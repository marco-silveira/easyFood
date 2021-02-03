package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"

	config "easyfood/config/http"
)

func CheckJwt(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(w, r)

			jwtCookie, err := r.Cookie("jwt")
			if err != nil {
				return
			}

			_, err = jwt.Parse(jwtCookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.JwtSecret), nil
			})
			if err != nil {
				return
			}

			//claims := token.Claims.(jwt.MapClaims)
			//id := claims["user_id"].(float64)
			//email := claims["user_email"].(string)
			//authorized := claims["authorized"].(bool)
		})
	}
}
