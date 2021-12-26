package bittrex

import "github.com/shopspring/decimal"

// Balance struct
type Balance struct {
	CurrencySymbol string          `json:"currencySymbol"`
	Total          decimal.Decimal `json:"total"`
	Available      decimal.Decimal `json:"available"`
	UpdatedAt      *jTime          `json:"updatedAt"`
}

//BalanceUpdate struct
type BalanceUpdate struct {
	AccountID string `json:"accountId"`
	Sequence  int    `json:"sequence"`
	Delta     struct {
		CurrencySymbol string          `json:"currencySymbol"`
		Total          decimal.Decimal `json:"total"`
		Available      decimal.Decimal `json:"available"`
		UpdatedAt      *jTime          `json:"updatedAt"`
	} `json:"delta"`
}
