package ports

import "github.com/cahllagerfeld/go-poke-cli/internal/core/domain"

type PokemonService interface {
	GetOne(id int) (*domain.Pokemon, error)
}
