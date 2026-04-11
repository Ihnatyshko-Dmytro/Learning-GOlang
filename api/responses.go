package coincap

import "fmt"

type AssetsResponse struct {
	Assets    []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}
type AssetResponse struct {
	Asset     AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

type AssetData struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Names             string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCupUsd      string `json:"marketCupUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

func (d AssetData) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s",
		d.ID, d.Rank, d.Symbol, d.Names, d.PriceUsd)
}