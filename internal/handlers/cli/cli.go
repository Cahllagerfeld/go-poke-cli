package cli

import (
	"fmt"
	"strconv"

	pokemonsrv "github.com/cahllagerfeld/go-poke-cli/internal/core/services/pokemon"
	"github.com/cahllagerfeld/go-poke-cli/internal/ports"
	"github.com/spf13/cobra"
)

type Handler struct {
	pokeService ports.PokemonService
}

func New(p *pokemonsrv.Service) *Handler {
	return &Handler{pokeService: p}
}

func (h *Handler) GetOne() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Args:  cobra.ExactArgs(1),
		Short: "Get Pokemon Info by number",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				panic("Failed to parse ID")
			}
			data, err := h.pokeService.GetOne(id)
			if err != nil {
				fmt.Println("No data")
			}
			fmt.Println(data.Name)
		},
	}
}
