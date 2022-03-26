package api

import core "github.com/cahllagerfeld/go-poke-cli/internal/core/app"

type Pokemon interface {
	GetOne(i int) (*core.PokemonResponse, error)
}
