package repository

import (
	"go-pokemon/entity"

	"github.com/mtslzr/pokeapi-go/structs"
)

type PokemonRepository interface {
	GetPokemon() []entity.ListPokemon
	DetailPokemon(name string) structs.Pokemon
	FindByID(id string) (poke entity.MyPokemon, err error)
	UpdateName(poke *entity.MyPokemon) error
	CatchPokemon(myPokemon *entity.MyPokemon) error
	MyPokemon() (poke []entity.MyPokemon, err error)
	RemovePokemon(id string) error
}
