package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesseobrien/exchange/config"
	"github.com/jesseobrien/exchange/routes"
	"github.com/urfave/negroni"
)

func main() {

	// @TODO make an init function to take flags to populate the config
	cfg := config.New()

	cfg.Database = config.DatabaseInit(cfg)
	defer cfg.Database.Close()

	router := mux.NewRouter()

	routes.Companies(router, cfg)

	// router.HandleFunc("/orders", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Header().Set("Content-type", "application/json")
	//
	// 	o := NewOrder()
	//
	// 	if req.Body == nil {
	// 		http.Error(w, "Please send a request body", 400)
	// 		return
	// 	}
	//
	// 	err := json.NewDecoder(req.Body).Decode(&o)
	//
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	//
	// 	err = o.Validate()
	//
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	//
	// 	exchange.postOrder(o)
	//
	// 	json.NewEncoder(w).Encode(o)
	// }).Methods("POST")
	//
	// router.HandleFunc("/orders/{symbol}", func(w http.ResponseWriter, req *http.Request) {
	// 	vars := mux.Vars(req)
	// 	symbol := vars["symbol"]
	// 	orders, err := exchange.GetOrders(symbol)
	//
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	//
	// 	json.NewEncoder(w).Encode(orders)
	//
	// }).Methods("GET")
	//
	// router.HandleFunc("/bids/{symbol}", func(w http.ResponseWriter, req *http.Request) {
	// 	vars := mux.Vars(req)
	// 	symbol := vars["symbol"]
	//
	// 	bids, err := exchange.GetBids(symbol)
	//
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	//
	// 	json.NewEncoder(w).Encode(bids)
	//
	// }).Methods("GET")
	//
	// router.HandleFunc("/ticker/{symbol}", func(w http.ResponseWriter, req *http.Request) {
	// 	vars := mux.Vars(req)
	// 	symbol := vars["symbol"]
	//
	// 	ticker, err := exchange.ticker(symbol)
	//
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 400)
	// 		return
	// 	}
	//
	// 	json.NewEncoder(w).Encode(ticker)
	// }).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(cfg.GetHttpHost(), n)
}
