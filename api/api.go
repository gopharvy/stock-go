package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stocks-go/model"
	"strings"
	"time"
)

type StockAPI interface {
	GetStockPrice(string, []string) (map[string]model.StockInfo, *model.Error)
}

type APIClient struct {
	config model.Config
}

func NewAPIClient(conf model.Config) StockAPI {
	return &APIClient{config: conf}
}

func (clt APIClient) GetStockPrice(stock string, indexes []string) (map[string]model.StockInfo, *model.Error) {
	url := fmt.Sprintf(clt.config.URL, stock, clt.config.Token)
	httpResponse := model.StockAPIResponse{}

	//make http calls
	err := clt.doHTTP(url, &httpResponse)
	if err != nil {
		return map[string]model.StockInfo{}, err
	}

	//create a exchange-value map and return only asked
	outboundResponse := clt.marshalApiHTTPResponse(httpResponse)

	//return default stock in case no index passed
	if len(indexes) == 0 {
		if stock, ok := outboundResponse[clt.config.DefaultExchange]; ok {
			return map[string]model.StockInfo{clt.config.DefaultExchange: stock}, nil
		}
	}

	//final response
	return clt.buildAPIResponse(indexes, outboundResponse), nil
}

//doHTTP makes http client, make request and calls outbound
func (clt APIClient) doHTTP(url string, resp interface{}) *model.Error {
	client := http.Client{
		Timeout: time.Millisecond * time.Duration(clt.config.Timeout),
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &model.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	response, err := client.Do(request)
	if err != nil {
		return &model.Error{Code: http.StatusBadGateway, Message: err.Error()}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &model.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return &model.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

//builds a map of the stocks based on the indexes
func (clt APIClient) marshalApiHTTPResponse(httpResponse model.StockAPIResponse) map[string]model.StockInfo {
	response := make(map[string]model.StockInfo, len(httpResponse.Data))

	for _, stock := range httpResponse.Data {
		response[stock.StockExchangeShort] = model.StockInfo{
			Symbol:         stock.Symbol,
			Name:           stock.Name,
			Price:          stock.Price,
			CloseYesterday: stock.CloseYesterday,
			Currency:       stock.Currency,
			MarketCap:      stock.MarketCap,
			Volume:         stock.Volume,
			Timezone:       stock.Timezone,
			TimezoneName:   stock.TimezoneName,
			GmtOffset:      stock.GmtOffset,
			LastTradeTime:  stock.LastTradeTime,
		}
	}

	return response
}

//searches the index-stock map, and returns the final response
func (clt APIClient) buildAPIResponse(indexes []string, stockInfos map[string]model.StockInfo) map[string]model.StockInfo {
	response := make(map[string]model.StockInfo, 0)
	for _, indx := range indexes {
		if len(indx) == 0 {
			continue
		}
		if stock, ok := stockInfos[indx]; ok {
			if strings.TrimSpace(indx) == "" {
				continue
			}
			response[indx] = stock
		} else {
			response[clt.config.DefaultExchange] = stock
		}
	}

	return response
}
