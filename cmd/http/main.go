package main

import (
	pokemonsrv "github.com/cahllagerfeld/go-poke-cli/internal/core/services/pokemon"
	"github.com/cahllagerfeld/go-poke-cli/internal/handlers/http"
)

func main() {
	ps := pokemonsrv.New()
	hH := http.New(ps)

	hH.Run()
}
