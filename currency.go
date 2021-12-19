package bittrex

import "github.com/shopspring/decimal"

// Currency struct
type Currency struct {
	Symbol                   string          `json:"symbol"`
	Name                     string          `json:"name"`
	CoinType                 string          `json:"coinType"`
	Status                   string          `json:"status"`
	MinConfirmation          int             `json:"minConfirmation"`
	Notice                   string          `json:"notice"`
	TxFee                    decimal.Decimal `json:"txFee"`
	LogoUrl                  string          `json:"logoUrl"`
	BaseAddress              string          `json:"baseAddress"`
	ProhibitedIn             []string        `json:"prohibitedIn"`
	AssociatedTermsOfService []string        `json:"associatedTermsOfService"`
	Tags                     []string        `json:"tags"`
}
