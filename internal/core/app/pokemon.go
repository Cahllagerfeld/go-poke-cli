package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cahllagerfeld/go-poke-cli/internal/core/domain"
)

func New() *Pokemon {
	return &Pokemon{}
}

type Pokemon struct{}

func (p *Pokemon) GetOne(i int) (*domain.Pokemon, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", i)

	target := &domain.Pokemon{}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch data: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode response: %w", err)
	}
	return target, nil
}
