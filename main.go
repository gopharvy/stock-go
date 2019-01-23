package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"stocks-go/base"
	"stocks-go/model"
)


func main(){

	var (
		service  base.Service
		handler  http.Handler
		serviceName   = flag.String("service.name", "stock", "microservice name used as basepath of the service")
		httpAddr      = flag.String("httpAddr", ":8080", "port application running and serving requests")
	)
	flag.Parse()


	//read config file
	errs:= make(chan error)
	c, err:= ioutil.ReadFile("resources/config.json")
	if err!=nil{
		errs <- err
	}

	//load configs
	var config model.Config
	err = json.Unmarshal(c, &config)
	if err!=nil{
		errs <- err
	}

	service = base.NewService(config)
	handler = service.BuildHTTPHandler(*serviceName)

	go func (){
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	error := <-errs
	log.Printf("exit\t:%v", error)// probably use logger in the prod service
}

