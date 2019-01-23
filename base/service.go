package base

import (
	"stocks-go/api"
	"stocks-go/model"
	)

type Service struct{
	config model.Config
	api api.StockAPI
}

func NewService(conf model.Config)Service{
	srvs:= Service{config:conf}
	srvs.api = api.NewAPIClient(conf)
	return srvs
}


