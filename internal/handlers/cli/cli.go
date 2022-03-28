package cli

import (
	"github.com/cahllagerfeld/go-poke-cli/internal/core/domain"
	"github.com/cahllagerfeld/go-poke-cli/internal/ports"
)

type Handler struct {
	pokeService ports.PokemonService
}

func New(p *ports.PokemonService) *Handler {
	return &Handler{pokeService: *p}
}

func (h *Handler) GetOne(id int) (*domain.Pokemon, error) {
	return &domain.Pokemon{}, nil
}
