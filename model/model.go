package model

type Config struct{
	URL string `json:"apiURL, omitempty"`
	Token string`json:"apiToken, omitempty"`
	Timeout int `json:"timeout, omitempty"`
	DefaultExchange   string `json:"defaultExchange"`
}
//StockAPIResponse is wrapper to response
type StockAPIResponse struct{
	SymbolRequested int `json:"symbols_requested"`
	SymbolReturned int `json:"symbols_returned"`
	Data []StockDetail `json:"data"`
}

//StockDetail is the api response struct
type StockDetail  struct{
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	Currency           string `json:"currency"`
	Price              string `json:"price"`
	PriceOpen          string `json:"price_open"`
	DayHigh            string `json:"day_high"`
	DayLow             string `json:"day_low"`
	Five2WeekHigh      string `json:"52_week_high"`
	Five2WeekLow       string `json:"52_week_low"`
	DayChange          string `json:"day_change"`
	ChangePct          string `json:"change_pct"`
	CloseYesterday     string `json:"close_yesterday"`
	MarketCap          string `json:"market_cap"`
	Volume             string `json:"volume"`
	VolumeAvg          string `json:"volume_avg"`
	Shares             string `json:"shares"`
	StockExchangeLong  string `json:"stock_exchange_long"`
	StockExchangeShort string `json:"stock_exchange_short"`
	Timezone           string `json:"timezone"`
	TimezoneName       string `json:"timezone_name"`
	GmtOffset          string `json:"gmt_offset"`
	LastTradeTime      string `json:"last_trade_time"`
}


type StockInfo  struct {
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
type Error struct{
	Code int
	Message string
}
