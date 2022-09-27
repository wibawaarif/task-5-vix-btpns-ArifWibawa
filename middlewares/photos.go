package middlewares

import (
	"FinalProject/app"
	"FinalProject/helpers"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(x http.ResponseWriter, y *http.Request) {
		cookie, err := y.Cookie("token")
		// check the token
		if err != nil {
			if err == http.ErrNoCookie {
				responseMsg := map[string]string{"info": "Unauthorized"}
				helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
				return 
			}
		}

		// take token value
		tokenStr := cookie.Value
		claims := &app.JWTuser{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return app.JWT_KEY, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			// check per case
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				responseMsg := map[string]string{"info": "Unauthorized"}
				helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
				return

			case jwt.ValidationErrorExpired:
				responseMsg := map[string]string{"info": "Unauthorized, Token Expired!"}
				helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
				return
			default:
				responseMsg := map[string]string{"info": "Unauthorized"}
				helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
				return
			}
		}

		if !token.Valid {
			responseMsg := map[string]string{"info": "Unauthorized"}
			helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
			return 
		}

		next.ServeHTTP(x, y)
		
	})
}
