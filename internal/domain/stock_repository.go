package domain

type StockRepository interface {
	Add(ticker string) (string, error)
	Search(search string) []*Stock
}
