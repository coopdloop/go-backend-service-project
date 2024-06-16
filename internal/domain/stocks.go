package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Todos is a list of Todo
type Stocks []*Stock

// NewTodos creates a new list of todos
func NewStocks() *Stocks {
	return &Stocks{}
}

// Add adds a todo to the Stocks object
func (t *Stocks) Add(ticker string) (string, error) {
	baseURI := fmt.Sprintf("/v2/aggs/ticker/%s/prev?", ticker)
	resp, err := http.Get(PolygonPath + baseURI + ApiKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	data := struct {
		Ticker  string  `json:"ticker"`
		Results []Stock `json:"results"`
	}{}

	json.Unmarshal(body, &data)
	log.Printf("Adding %s", ticker)
	log.Printf("Searching for %s", ticker)
	res := t.Search(ticker)
	if len(res) >= 0 {
		log.Printf("Found ticker object %v", res)
	}
	for _, stock := range data.Results {
		stock.Ticker = ticker
		log.Println(stock)
		*t = append(*t, &stock)
	}
	return fmt.Sprintf("Successfully added %s", ticker), nil
}

// Search a stock
func (t *Stocks) Search(ticker string) []*Stock {
	log.Printf("Searching for %s", ticker)
	list := make([]*Stock, 0)
	for _, stock := range *t {
		log.Println(stock)
		if strings.Contains(stock.Ticker, ticker) {
			list = append(list, stock)
		}
	}
	return list
}

func (t *Stocks) Delete(stock *Stock) string {
	log.Printf("Removing %s from the database", stock.Ticker)
	for index, stk := range *t {
		if strings.Contains(stk.Ticker, stock.Ticker) {
			log.Println(index)
			log.Println(stk)
			// t = append(t[:index], t[index+1]...)
		}
	}
	return "Deleted"
}
