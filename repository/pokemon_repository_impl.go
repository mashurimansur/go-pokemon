package repository

import (
	"go-pokemon/entity"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"gorm.io/gorm"
)

type pokemonRepositoryImpl struct {
	DB *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) PokemonRepository {
	return &pokemonRepositoryImpl{DB: db}
}

func (repository *pokemonRepositoryImpl) GetPokemon() []entity.ListPokemon {
	l, err := pokeapi.Resource("pokemon", 1, 60)
	if err != nil {
		return []entity.ListPokemon{}
	}

	lists := []entity.ListPokemon{}
	for _, v := range l.Results {
		image, _ := pokeapi.Pokemon(v.Name)

		list := entity.ListPokemon{
			Image: image.Sprites.FrontDefault,
			Name:  v.Name,
			URL:   v.URL,
		}

		lists = append(lists, list)

	}

	return lists
}

func (repository *pokemonRepositoryImpl) DetailPokemon(name string) (pokemon structs.Pokemon) {
	l, err := pokeapi.Pokemon(name)
	if err != nil {
		return
	}

	return l
}

func (repository *pokemonRepositoryImpl) CatchPokemon(myPokemon *entity.MyPokemon) (err error) {
	if err := repository.DB.Model(myPokemon).Create(&myPokemon).Error; err != nil {
		return err
	}
	return nil
}

func (repository *pokemonRepositoryImpl) MyPokemon() (poke []entity.MyPokemon, err error) {
	if err := repository.DB.Order("created_at asc").Find(&poke).Error; err != nil {
		return nil, err
	}
	return
}

func (repo *pokemonRepositoryImpl) RemovePokemon(id string) error {
	if err := repo.DB.Where("id = ?", id).Delete(&entity.MyPokemon{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pokemonRepositoryImpl) FindByID(id string) (poke entity.MyPokemon, err error) {
	if err := repo.DB.Where("id = ?", id).First(&poke).Error; err != nil {
		return poke, err
	}
	return
}

func (repo *pokemonRepositoryImpl) UpdateName(poke *entity.MyPokemon) (err error) {
	if err := repo.DB.Model(poke).Updates(poke).Error; err != nil {
		return err
	}

	return nil
}
