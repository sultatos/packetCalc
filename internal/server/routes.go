package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"othonas/internal/service"
	"othonas/internal/views"
	"strconv"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /packets", s.PacketHandler)
	mux.HandleFunc("POST /packet-sizes", s.UpdateSizes)
	mux.Handle("/", templ.Handler(views.Home(&s.PackSizes)))
	return mux
}

type packetsSizes struct {
	PackSizes []int `json:"packSizes"`
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

// Handler function for the API
func (s *Server) UpdateSizes(w http.ResponseWriter, r *http.Request) {
	var sizes packetsSizes
	err := json.NewDecoder(r.Body).Decode(&sizes)
	if err != nil {
		http.Error(w, "Invalid packet sizes", http.StatusBadRequest)
		return
	}
	s.PackSizes = sizes.PackSizes
	fmt.Println("Updated pack sizes:", s.PackSizes)
}
