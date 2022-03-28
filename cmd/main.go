package main

import (
	pokemonsrv "github.com/cahllagerfeld/go-poke-cli/internal/core/services/pokemon"
	"github.com/cahllagerfeld/go-poke-cli/internal/handlers/cli"
)

func main() {

	ps := pokemonsrv.New()
	cH := cli.New(ps)
	cli.RootCmd.AddCommand(cH.GetOne())
	cli.Execute()
}
