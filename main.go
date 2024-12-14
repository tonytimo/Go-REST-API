package main

import (
	"net/http"
	"regexp"

	"github.com/tonytimo/Go-REST-API/recipes"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	mux.Handle("/recipes", &recipesHandler{})
	mux.Handle("/recipes/", &recipesHandler{})

	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

type recipesHandler struct{}

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)

func (h *recipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}
}

func (h *recipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *recipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request)  {}
func (h *recipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request)    {}
func (h *recipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {}
func (h *recipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	List() (map[string]recipes.Recipe, error)
	Remove(name string) error
}
