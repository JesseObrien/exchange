package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Order struct {
	Id        uuid.UUID `json:"id"`
	ClientId  uuid.UUID `json:"client_id"`
	Symbol    string    `json:"symbol"`
	Price     int64     `json:"price"`
	Shares    int64     `json:"shares"`
	Type      string    `json:"type"`
	Action    string    `json:"action"`
	Status    string    `json:"status"`
	Timestamp int64     `json:"timestamp"`
}

func NewSellOrder(symbol string) *Order {
	rand.Seed(time.Now().UnixNano())

	return &Order{
		Id:       uuid.NewV4(),
		ClientId: uuid.NewV4(),
		Symbol:   symbol,
		Price:    rand.Int63n(10000),
		Shares:   rand.Int63n(1000),
		Action:   "sell",
	}
}

func (o *Order) Validate() error {

	validTypes := make(map[string]bool)
	validTypes["market"] = true

	validActions := make(map[string]bool)
	validActions["buy"] = true
	validActions["sell"] = true

	if uuid.Nil == o.ClientId {
		return errors.New("Invalid client id")
	}

	if o.Type != "" {
		if _, ok := validTypes[o.Type]; !ok {
			return errors.New(fmt.Sprintf("Type %s is not valid.", o.Type))
		}
	}

	if o.Action != "" {
		if _, ok := validActions[o.Action]; !ok {
			return errors.New(fmt.Sprintf("Type %s is not valid.", o.Type))
		}
	}

	return nil
}

func (slice Orders) String() string {
	if slice.Len() < 1 {
		return "Bid: $0.00"
	} else {
		price := strconv.FormatInt(slice[slice.Len()-1].Price, 10)
		left := price[0 : len(price)-2]
		right := price[len(price)-2:]
		return fmt.Sprintf("Bid: $%s.%s", left, right)
	}
}

func NewOrder() *Order {
	return &Order{
		Id:        uuid.NewV4(),
		Timestamp: time.Now().Unix(),
	}
}

type Orders []*Order

func (slice Orders) Len() int {
	return len(slice)
}

func (slice Orders) Less(i, j int) bool {
	if slice[i].Price == slice[j].Price {
		return slice[i].Timestamp < slice[j].Timestamp
	}

	return slice[i].Price < slice[j].Price
}

func (slice Orders) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Orders) remove(id uuid.UUID) {
	for _, order := range slice {
		if order.Id == id {
		}
	}
}
