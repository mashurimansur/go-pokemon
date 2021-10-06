package service

import (
	"go-pokemon/entity"

	"github.com/mtslzr/pokeapi-go/structs"
)

type PokemonService interface {
	ListPokemon() []entity.ListPokemon
	DetailPokemon(name string) structs.Pokemon
	FindByID(id string) (poke entity.MyPokemon, err error)
	UpdateName(poke *entity.MyPokemon) error
	CatchPokemon(myPokemon *entity.MyPokemon) error
	MyPokemon() (poke []entity.MyPokemon, err error)
	ReleasePokemon(id string) error
	ChangeNamePokemon(name string) string
}
