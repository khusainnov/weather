package handler

import (
	"github.com/gorilla/mux"
	"github.com/khusainnov/weather/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutesMux() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/weather/form", h.Form).Methods("GET")
	r.HandleFunc("/weather", h.Weather)

	return r
}
