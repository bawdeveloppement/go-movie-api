package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MovieRouter struct{}

func (movieRouter *MovieRouter) CreateRouter(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	fmt.Println(head)

	if head == "" {
		switch r.Method {
		case "GET":
			movieRouter.getMoviesHandler().ServeHTTP(w, r)
		case "POST":
			movieRouter.createMovieHandler().ServeHTTP(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	} else {
		// We check if one movie with the target id exist
		filterResult := movies.Filter(func(value Movie) bool {
			return value.Id == head
		})
		if len(filterResult) > 0 {
			switch r.Method {
			case "GET":
				movieRouter.getSingleMovieHandler(head).ServeHTTP(w, r)
			case "PUT":
				movieRouter.updateMovieHandler(head).ServeHTTP(w, r)
			case "DELETE":
				movieRouter.deleteMovieHandler(head).ServeHTTP(w, r)
			default:
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		} else {
			http.Error(w, "Movie Not Found", http.StatusNotFound)
		}
	}
}

func (movieRouter *MovieRouter) getMoviesHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(movies)
	})
}

func (movieRouter *MovieRouter) getSingleMovieHandler(id string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		filterResult := movies.Filter(func(value Movie) bool {
			return value.Id == id
		})
		if len(filterResult) > 0 {
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(filterResult[0])
		} else {
			http.Error(res, "Movie Not Found", http.StatusNotFound)
		}
	})
}

func (movieRouter *MovieRouter) createMovieHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var movie Movie
		err := json.NewDecoder(req.Body).Decode(&movie)
		if err != nil {
			http.Error(res, "Cannot create a movie", http.StatusBadRequest)
		} else {
			if len(movies.Filter(func(value Movie) bool {
				return value.Id == movie.Id
			})) > 0 {
				http.Error(res, "Movie with same id already exist", http.StatusConflict)
			} else {
				movies.ConcatValues(movie)
				json.NewEncoder(res).Encode(movies)
			}
		}
	})
}

func (movieRouter *MovieRouter) updateMovieHandler(id string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		foundIndex := movies.FindIndex(func(value Movie) bool {
			return value.Id == id
		})
		if foundIndex != -1 {
			var movie Movie
			err := json.NewDecoder(req.Body).Decode(&movie)
			if err != nil {
				http.Error(res, "Movie with same id already exist", http.StatusNotFound)
			} else {
				movies[foundIndex] = movie
				json.NewEncoder(res).Encode(movies)
			}
		} else {
			http.Error(res, "Movie with same id already exist", http.StatusNotFound)
		}
	})
}

// Supprime un film du tableau de film movies à partir du parametre id de l'url
func (movieRouter *MovieRouter) deleteMovieHandler(id string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(movies.Filter(func(value Movie) bool {
			return value.Id == id
		})) > 0 {
			movies.FilterInternal(func(value Movie) bool {
				return value.Id != id
			})
			json.NewEncoder(res).Encode(movies)
		} else {
			http.Error(res, "Movie with same id already exist", http.StatusNotFound)
		}
	})
}

// ---- Types ----

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// ---------------
