package domain

import (
	"fmt"

	"github.com/coopdloop/go-backend-service-project/internal/config"
)

const PolygonPath = "https://api.polygon.io"

var ApiKey = fmt.Sprintf("apiKey=%s", config.Envs.PolygonApiKey)

type Stock struct {
	Ticker string  `json:"ticker"`
	O      float64 `json:"o"`
	H      float64 `json:"h"`
	L      float64 `json:"l"`
	C      float64 `json:"c"`
}

type Values struct {
	Open float64 `json:"open"`
}

// NewStock creates a new Stock
func NewStock(ticker string, name string, price float64) *Stock {
	return &Stock{
		Ticker: ticker,
	}
}

// Update updates a stock
func (t *Stock) Update(stock Stock) {
	t = &stock
}
