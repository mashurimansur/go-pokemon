package entity

import "time"

type ListPokemon struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

type MyPokemon struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Image     string    `json:"image"`
}
