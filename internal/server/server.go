package server

import (
	"fmt"
	"net/http"
	"os"
	"othonas/internal/service"
	"sort"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port      int
	PackSizes []int `json:"pack_sizes"`
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:      port,
		PackSizes: service.LoadPackSizesFromFile("pack_sizes.json"),
	}
	sort.Slice(NewServer.PackSizes, func(i, j int) bool {
		return NewServer.PackSizes[i] > NewServer.PackSizes[j]
	})

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
