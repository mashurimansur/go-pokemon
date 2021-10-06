package service

import (
	"fmt"
	"go-pokemon/entity"
	"go-pokemon/repository"
	"math"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go/structs"
)

type pokemonServiceImpl struct {
	PokemonRepository repository.PokemonRepository
}

func NewPokemonService(pokemonRepository *repository.PokemonRepository) PokemonService {
	return &pokemonServiceImpl{PokemonRepository: *pokemonRepository}
}

func (service *pokemonServiceImpl) ListPokemon() []entity.ListPokemon {
	return service.PokemonRepository.GetPokemon()
}

func (service *pokemonServiceImpl) DetailPokemon(name string) structs.Pokemon {
	return service.PokemonRepository.DetailPokemon(name)
}

func (service *pokemonServiceImpl) CatchPokemon(myPokemon *entity.MyPokemon) error {
	return service.PokemonRepository.CatchPokemon(myPokemon)
}

func (service *pokemonServiceImpl) MyPokemon() (poke []entity.MyPokemon, err error) {
	return service.PokemonRepository.MyPokemon()
}

func (service *pokemonServiceImpl) ReleasePokemon(id string) error {
	return service.PokemonRepository.RemovePokemon(id)
}

func (service *pokemonServiceImpl) FindByID(id string) (poke entity.MyPokemon, err error) {
	return service.PokemonRepository.FindByID(id)
}

func (service *pokemonServiceImpl) UpdateName(poke *entity.MyPokemon) (err error) {
	return service.PokemonRepository.UpdateName(poke)
}

func (service *pokemonServiceImpl) ChangeNamePokemon(name string) string {
	splitName := strings.Split(name, "-")

	if len(splitName) == 1 {
		return splitName[0] + "-0"
	}

	if splitName[len(splitName)-1] == "0" {
		return splitName[0] + "-1"
	}

	num, _ := strconv.Atoi(splitName[len(splitName)-1])

	newName := fmt.Sprintf("%s-%d", splitName[0], service.NextFibonacci(num))

	return newName
}

func (service *pokemonServiceImpl) NextFibonacci(n int) int {
	a := float64(n) * (1 + math.Sqrt(5)) / 2.0
	return int(math.Round(a))
}
