package stock

import (
	"context"

	"github.com/coopdloop/go-backend-service-project/internal/domain"
)

type (
	Service interface {
		// Add adds a ticker to the list
		Add(ctx context.Context, ticker string) (string, error)
		// Search returns a list of tickers that match the search string
		Search(ctx context.Context, search string) ([]*domain.Stock, error)
	}

	service struct {
		stocks domain.StockRepository
	}
)

func NewService(stocks domain.StockRepository) Service {
	return &service{
		stocks: stocks,
	}
}

func (s service) Add(_ context.Context, ticker string) (string, error) {
	msg, err := s.stocks.Add(ticker)
	return msg, err
}

func (s service) Search(_ context.Context, search string) ([]*domain.Stock, error) {
	todos := s.stocks.Search(search)

	return todos, nil
}
