package ports

import "github.com/cahllagerfeld/go-poke-cli/internal/core/domain"

type PokemonService interface {
	Get(id int) (domain.Pokemon, error)
}
