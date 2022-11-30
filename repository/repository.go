package repository

import (
	model "rodrigofrancisco/go-bootcamp/domain/model"
)

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
}
