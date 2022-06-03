package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthRouter struct{}

func (authRouter *AuthRouter) CreateRouter(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	if head == "" {
		switch r.Method {
		case "PUT":
			authRouter.signInHandler().ServeHTTP(w, r)
		case "POST":
			authRouter.signUpHandler().ServeHTTP(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}
}

type AuthObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (authRouter *AuthRouter) signInHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var auth AuthObject
		json.NewDecoder(req.Body).Decode(&auth)
		foundAuth := auths.FindIndex(func(value Auth) bool { return value.Email == auth.Email })
		if foundAuth != -1 {
			err := bcrypt.CompareHashAndPassword([]byte(auths[foundAuth].Password), []byte(auth.Password))
			if err != nil {
				fmt.Errorf("Error: %s", err)
				return
			} else {
				// implement jwt
			}
		} else {
			// Auth not found
		}
	})
}

func (authRouter *AuthRouter) signUpHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

	})
}

type Auth struct {
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	Password     string
}
