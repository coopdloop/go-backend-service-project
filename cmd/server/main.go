package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/coopdloop/go-backend-service-project/internal/config"
	"github.com/coopdloop/go-backend-service-project/internal/domain"
	"github.com/coopdloop/go-backend-service-project/internal/features/stock"
)

func main() {

	router := chi.NewRouter()
	stocks := domain.NewStocks()

	stock.Mount(router, stock.NewHandler(stock.NewService(stocks)))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Envs.Port),
		Handler: http.TimeoutHandler(router, 30*time.Second, "request timed out"),
	}

	// Display the localhost address and port
	fmt.Printf("Listening on http://localhost%s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
