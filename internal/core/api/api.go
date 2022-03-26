package api

import (
	core "github.com/cahllagerfeld/go-poke-cli/internal/core/app"
)

type Application struct {
	p Pokemon
}

func NewApplication(p Pokemon) *Application {
	return &Application{p: p}
}

func (apia Application) GetOne(i int) (*core.PokemonResponse, error) {
	resp, err := apia.GetOne(i)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
