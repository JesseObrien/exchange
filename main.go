package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	exchange := CreateExchange()

	exchange.seed("TSLA")

	router := mux.NewRouter()

	router.HandleFunc("/orders", func(w http.ResponseWriter, req *http.Request) {

		o := NewOrder()

		if req.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		err := json.NewDecoder(req.Body).Decode(&o)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = o.Validate()

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		exchange.postOrder(o)

		json.NewEncoder(w).Encode(o)
	}).Methods("POST")

	router.HandleFunc("/orders/{symbol}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		symbol := vars["symbol"]
		orders, err := exchange.GetOrders(symbol)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		json.NewEncoder(w).Encode(orders)

	}).Methods("GET")

	router.HandleFunc("/bids/{symbol}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		symbol := vars["symbol"]

		bids, err := exchange.GetBids(symbol)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		json.NewEncoder(w).Encode(bids)

	}).Methods("GET")

	router.HandleFunc("/ticker/{symbol}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		symbol := vars["symbol"]

		ticker, err := exchange.ticker(symbol)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		json.NewEncoder(w).Encode(ticker)
	}).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":3000", n)
}
