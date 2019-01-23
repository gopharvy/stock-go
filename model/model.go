package model

type Config struct {
	URL             string `json:"apiURL,omitempty"`
	Token           string `json:"apiToken,omitempty"`
	Timeout         int    `json:"timeout,omitempty"`
	DefaultExchange string `json:"defaultExchange"`
}

//StockAPIResponse is wrapper to response
type StockAPIResponse struct {
	SymbolRequested int           `json:"symbols_requested"`
	SymbolReturned  int           `json:"symbols_returned"`
	Data            []StockDetail `json:"data,omitempty"`
}

//StockDetail is the api response struct
type StockDetail struct {
	Symbol             string `json:"symbol,omitempty"`
	Name               string `json:"name,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Price              string `json:"price"`
	PriceOpen          string `json:"price_open,omitempty"`
	DayHigh            string `json:"day_high,omitempty"`
	DayLow             string `json:"day_low,omitempty"`
	Five2WeekHigh      string `json:"52_week_high,omitempty"`
	Five2WeekLow       string `json:"52_week_low,omitempty"`
	DayChange          string `json:"day_change,omitempty"`
	ChangePct          string `json:"change_pct,omitempty"`
	CloseYesterday     string `json:"close_yesterday,omitempty"`
	MarketCap          string `json:"market_cap,omitempty"`
	Volume             string `json:"volume,omitempty"`
	VolumeAvg          string `json:"volume_avg,omitempty"`
	Shares             string `json:"shares,omitempty"`
	StockExchangeLong  string `json:"stock_exchange_long,omitempty"`
	StockExchangeShort string `json:"stock_exchange_short,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	TimezoneName       string `json:"timezone_name,omitempty"`
	GmtOffset          string `json:"gmt_offset,omitempty"`
	LastTradeTime      string `json:"last_trade_time,omitempty"`
}

type StockInfo struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Price          string `json:"price"`
	CloseYesterday string `json:"close_yesterday"`
	Currency       string `json:"currency"`
	MarketCap      string `json:"market_cap"`
	Volume         string `json:"volume"`
	Timezone       string `json:"timezone"`
	TimezoneName   string `json:"timezone_name"`
	GmtOffset      string `json:"gmt_offset"`
	LastTradeTime  string `json:"last_trade_time"`
}

//Error wrapper
type Error struct {
	Code    int
	Message string
}
