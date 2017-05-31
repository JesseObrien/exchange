package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Exchange struct {
	orders map[string]Orders
}

func CreateExchange() *Exchange {
	return &Exchange{
		orders: make(map[string]Orders),
	}
}

func (e *Exchange) seed(symbol string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rand.Intn(100); i++ {
		e.postOrder(NewSellOrder(symbol))
	}
}

func (e *Exchange) postOrder(o *Order) {
	e.orders[o.Symbol] = append(e.orders[o.Symbol], o)
	sort.Sort(e.orders[o.Symbol])

	if o.Action == "buy" {
		for _, order := range e.orders[o.Symbol] {
			if o.Type == "market" {
				if o.Shares <= order.Shares {

				}
			}
		}
	}
}

func (e *Exchange) executeOrder(o *Order) {
	e.orders[o.Symbol] = append(e.orders[o.Symbol], o)

	sort.Sort(e.orders[o.Symbol])
}

func (e *Exchange) ticker(symbol string) (string, error) {

	bids, err := e.GetBids(symbol)

	return fmt.Sprintf("%s > %s", symbol, bids.String()), err
}

func (e *Exchange) GetAllOrders() map[string]Orders {
	return e.orders
}

func (e *Exchange) GetOrders(symbol string) (Orders, error) {
	orders, ok := e.orders[symbol]

	if !ok {
		return make(Orders, 0), errors.New(fmt.Sprintf("Ticker symbol [%s] not found in orders", symbol))
	}

	return orders, nil
}

func (e *Exchange) GetBids(symbol string) (Orders, error) {
	orders, ok := e.orders[symbol]

	if !ok {
		return make(Orders, 0), errors.New(fmt.Sprintf("Ticker symbol [%s] not found", symbol))
	}

	return orders, nil
}
