package controllers

import (
	"errors"
	"net/http"
	model "rodrigofrancisco/go-bootcamp/domain/model"
)

type controller struct{}

func NewPokemonController() PokemonController {
	return &controller{}
}

func (*controller) FetchPokemons(w http.ResponseWriter, r *http.Request) error {
	return errors.New("TODO: remove this impl")
}

func (*controller) ShowPokemons(w http.ResponseWriter, r *http.Request) ([]*model.Pokemon, error) {
	err := errors.New("TODO: remove this temporal err")
	return nil, err
}
