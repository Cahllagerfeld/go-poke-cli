package pokemonsrv

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cahllagerfeld/go-poke-cli/internal/core/domain"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (serv *Service) GetOne(id int) (*domain.Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)

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
