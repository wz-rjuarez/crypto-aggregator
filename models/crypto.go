package models

import "time"

type Crypto struct {
	Date         time.Time `json:"date"`
	Name         string    `json:"name"`
	TickerSymbol string    `json:"ticker_symbol"`
	Price        Price     `json:"price"`
}

type Price struct {
	USD float64 `json:"usd"`
	MXN float64 `json:"mxn"`
}
