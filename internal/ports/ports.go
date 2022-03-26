package ports

import (
	core "github.com/cahllagerfeld/go-poke-cli/internal/core/app"
)

type APIPort interface {
	GetOne(i int) (*core.PokemonResponse, error)
}
