package main

import (
	"log"
	"net/http"

	"path"
	"strings"
)

var (
	users Array[User] = Array[User]{
		{Email: "baw.developpement@gmail.com"},
	}
	// We instantiate an Array[Movie], using the builtin syntax (e.g []Movie ) and we keep having access to Array.Map
	movies Array[Movie] = []Movie{
		{
			Id:       "pirate-caraibe",
			Title:    "The Pirate of caraibe",
			Director: &Director{Firstname: "Jason", Lastname: "Perse"},
		},
	}
)

func main() {
	application := &Application{
		UserRouter:  new(UserRouter),
		MovieRouter: new(MovieRouter),
	}
	log.Fatal(http.ListenAndServe(":3333", application))
}

type Application struct {
	UserRouter  *UserRouter
	MovieRouter *MovieRouter
}

// Application Router
func (application *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "user":
		application.UserRouter.Handle(w, r)
	case "movies":
		application.MovieRouter.Handle(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// Utils
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
