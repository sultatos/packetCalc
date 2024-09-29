package server

import (
	"encoding/json"
	"net/http"
	"othonas/internal/service"
	"othonas/internal/views"
	"strconv"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /packets", s.PacketHandler)
	mux.Handle("/", templ.Handler(views.Home()))
	return mux
}

type itemsOrder struct {
	Items int `json:"items"`
}

// Handler function for the API
func (s *Server) PacketHandler(w http.ResponseWriter, r *http.Request) {
	packetSize, err := strconv.Atoi(r.FormValue("items"))
	if err != nil {
		http.Error(w, "Invalid items size", http.StatusBadRequest)
		return
	}
	// Calculate the optimal pack distribution
	packDistribution := service.CalculatePacks(packetSize, s.PackSizes)
	// Convert the result to JSON and send it back to the client
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(packDistribution)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
