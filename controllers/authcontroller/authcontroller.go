package authcontroller

import (
	"FinalProject/app"
	"FinalProject/database"
	"FinalProject/helpers"
	"FinalProject/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// for register
func Register(x http.ResponseWriter, y *http.Request) {
	// fetch json from user
	var user models.User
	fetch := json.NewDecoder(y.Body)
	if err := fetch.Decode(&user)
	err != nil {
		responseMsg := map[string]string{"Info": err.Error()}
		helpers.ResponseJSON(x, http.StatusBadRequest, responseMsg)
		return
	}
	defer	y.Body.Close()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// hash the password using bcrypt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		responseMsg := map[string]string{"Info": err.Error()}
		helpers.ResponseJSON(x, http.StatusNotAcceptable, responseMsg)
		return
	}
	user.Password = string(passwordHash)

	// insert to database
	if err := database.DB.Create(&user).Error
	err != nil {
		responseMsg := map[string]string{"Info": err.Error()}
		helpers.ResponseJSON(x, http.StatusInternalServerError, responseMsg)
		return
	}

	// response to the user
	responseMsg := map[string]string{"Info": "Account Created"}
	helpers.ResponseJSON(x, http.StatusAccepted, responseMsg)
}

// for login
func Login(x http.ResponseWriter, y *http.Request) {
		// fetch json from user
		var user models.User
		fetch := json.NewDecoder(y.Body)
		if err := fetch.Decode(&user)
			err != nil {
				responseMsg := map[string]string{"Info": err.Error()}
				helpers.ResponseJSON(x, http.StatusBadRequest, responseMsg)
				return
			}
		defer	y.Body.Close()
			
		// check credentials
		var userDB models.User
		if err := database.DB.Where("email = ?", user.Email).First(&userDB).Error
			err != nil {
				responseMsg := map[string]string{"Info": err.Error()}
				helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
				return
			}

		if err:= bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
			err != nil {
				responseMsg := map[string]string{"Info": err.Error()}
				helpers.ResponseJSON(x, http.StatusUnauthorized, responseMsg)
				return
			}

		// create jwt
		claims := &app.JWTuser{
			Email: userDB.Email,
			StandardClaims: jwt.StandardClaims{
				Issuer: "go-jwt-mux",
				ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			},
		}

		// declare sign in algorithm
		tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// signed token

		token, err := tokenAlgorithm.SignedString(app.JWT_KEY)
		if err != nil {
			responseMsg := map[string]string{"Info": err.Error()}
			helpers.ResponseJSON(x, http.StatusInternalServerError, responseMsg)
			return
		}

		// set token to cookie
		http.SetCookie(x, &http.Cookie{
			Name: "token",
			Path: "/",
			Value: token,
			HttpOnly: true,
		})

		responseMsg := map[string]string{"Info": "Login Succeed!"}
		helpers.ResponseJSON(x, http.StatusOK, responseMsg)
}

// for logout
func Logout(x http.ResponseWriter, y *http.Request) {
	// delete existing token in cookie
	http.SetCookie(x, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	responseMsg := map[string]string{"Info": "Logout Succeed!"}
	helpers.ResponseJSON(x, http.StatusOK, responseMsg)
	
}