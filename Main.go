package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/lib/pq"
)

var (
	// Array est un tableau générique
	// Voir le fichier array.go pour plus de détail sur les méthodes
	auths Array[Auth] = Array[Auth]{}
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

type Application struct {
	AuthRouter  *AuthRouter
	UserRouter  *UserRouter
	MovieRouter *MovieRouter
}

func createMovieInDb(db *sql.DB, movie Movie) {
	result, err := db.Exec(
		"INSERT INTO movies (id, title) VALUES ($1, $2)",
		movie.Id,
		movie.Title,
	)

	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(&result)
}

func findAllMovie(db *sql.DB) {

}

func main() {
	connStr := "postgresql://postgres:root@127.0.0.1/movie-app?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pq.Einfo)
	rows, err := db.Query("SELECT * FROM movies")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if err != nil {
		log.Fatalln(err) /*  */
	}
	var res string
	for rows.Next() {
		rows.Scan(&res)
		fmt.Println(res)
		// movies.ConcatValues(res)
	}
	fmt.Println(movies)

	application := &Application{
		AuthRouter:  new(AuthRouter),
		UserRouter:  new(UserRouter),
		MovieRouter: new(MovieRouter),
	}
	log.Fatal(http.ListenAndServe(":3333", application))

}

// Application Router
func (application *Application) ServeHTTP(req http.ResponseWriter, res *http.Request) {
	var head string
	head, res.URL.Path = ShiftPath(res.URL.Path)

	switch head {
	case "auth":
		application.AuthRouter.CreateRouter(req, res)
	case "user":
		application.UserRouter.CreateRouter(req, res)
	case "movies":
		application.MovieRouter.CreateRouter(req, res)
	default:
		http.Error(req, "Not Found", http.StatusNotFound)
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
