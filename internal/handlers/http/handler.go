package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	pokemonsrv "github.com/cahllagerfeld/go-poke-cli/internal/core/services/pokemon"
	"github.com/cahllagerfeld/go-poke-cli/internal/ports"
	"github.com/gorilla/mux"
)

type Handler struct {
	pokeService ports.PokemonService
}

func New(p *pokemonsrv.Service) *Handler {
	return &Handler{
		pokeService: p,
	}
}

func (h *Handler) GetById(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(rw, "Please provide valid ID", http.StatusBadRequest)
	}
	data, err := h.pokeService.GetOne(id)
	if err != nil {
		http.Error(rw, "Failed to fetch Pokemon", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(data)
}

func (h *Handler) Run() {
	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/pokemon/{id}", h.GetById)

	http.ListenAndServe(":9090", r)
}
