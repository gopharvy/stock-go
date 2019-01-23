package base

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	pathKey  = "symbol"
	queryKey = "stock_exchange"
)

func (s Service) BuildHTTPHandler(basePath string) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/"+basePath+"/health", s.health).Methods(http.MethodGet)
	router.HandleFunc("/"+basePath+"/", s.getStock).Methods(http.MethodGet)
	router.HandleFunc("/"+basePath+"/{symbol}", s.getStock).Methods(http.MethodGet)

	return router
}

//getStock call the api getstockinfo
func (s Service) getStock(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	stockID := vars[pathKey]

	if len(stockID) == 0 {
		s.serveResponse(resp, http.StatusBadRequest, []byte("Invalid Stock id.."))
		return
	}

	queryVal := req.URL.Query().Get(queryKey)
	indexes := strings.Split(strings.TrimSpace(queryVal), ",")

	response, apierr := s.api.GetStockPrice(stockID, indexes)
	if apierr != nil {
		log.Printf("errors\t:%v", apierr)
		s.serveResponse(resp, apierr.Code, nil)
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Printf("errors\t:%v", err)
		s.serveResponse(resp, http.StatusInternalServerError, nil)
	}

	s.serveResponse(resp, http.StatusOK, body)

	return
}

//serveResponse method resolves the http response
func (s Service) serveResponse(resp http.ResponseWriter, statusCode int, body []byte) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(statusCode)
	resp.Write(body)
}

//health returns service health
func (s Service) health(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Service Healthy and Ready to serve request..."))
}
