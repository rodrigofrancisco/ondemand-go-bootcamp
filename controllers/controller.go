package controllers

import (
	"net/http"
	model "rodrigofrancisco/go-bootcamp/domain/model"
)

type PokemonController interface {
	FetchPokemons(w http.ResponseWriter, r *http.Request) error
	ShowPokemons(w http.ResponseWriter, r *http.Request) ([]*model.Pokemon, error)
}
