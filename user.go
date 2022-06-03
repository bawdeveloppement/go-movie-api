package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type UserRouter struct{}

func (userRouter *UserRouter) Handle(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		head, tail := ShiftPath(r.URL.Path)
		fmt.Println(head, tail)

		switch head {
		case "profile":
			userRouter.UserProfileHandler(id).ServeHTTP(w, r)
		case "movies":
			// Without declared handler
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				fmt.Println("UsersMovies")
			}).ServeHTTP(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	} else {
		userRouter.UserHomeHandler(id).ServeHTTP(w, r)
	}
}

func (h *UserRouter) UserHomeHandler(id int) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

	})
}

func (h *UserRouter) UserProfileHandler(id int) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("dazdzad")
	})
}

// ----- Types -----

type User struct {
	Email string `required max:"100"`
}

func (user *User) changeEmailAdress(email string) bool {
	user.Email = email
	return true
}

// Get the userId by an email address
func getUserIdOf(email string) (userId int) {
	userId = -1

	for idx, user := range users {
		if user.Email == email {
			userId = idx
		}
	}

	return userId
}

// -------
