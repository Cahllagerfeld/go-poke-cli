package pokemonsrv

import "github.com/cahllagerfeld/go-poke-cli/internal/core/domain"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (serv *Service) Get(id int) (domain.Pokemon, error) {
	return domain.Pokemon{}, nil
}
