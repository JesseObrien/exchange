package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

type Company struct {
	Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Symbol      string    `json:"symbol"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewCompany() *Company {
	return &Company{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Company) MarshalJSON() ([]byte, error) {
	type Alias Company
	return json.Marshal(&struct {
		Symbol string `json:"symbol"`
		*Alias
	}{
		Symbol: strings.ToUpper(c.Symbol),
		Alias:  (*Alias)(c),
	})
}

func (c *Company) UnmarshalJSON(data []byte) error {

	type Alias Company

	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &a); err != nil {
		log.Panic(err)
	}

	fmt.Println("%v", c)
	// Change the symbol to an uppercase representation ALWAYS
	c.Symbol = strings.ToUpper(c.Symbol)

	return nil
}

func (c *Company) Save(db *bolt.DB) {
	fmt.Println("%v", c)
	err := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("Companies"))
		fmt.Println("%s", c.Marshal())
		return b.Put([]byte(c.Symbol), c.Marshal())
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (c *Company) GetById(db *bolt.DB, symbol string) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Companies"))
		fmt.Println("%s", symbol)
		v := b.Get([]byte(symbol))

		fmt.Println("%v", v)

		if len(v) != 0 {
			c.FromBytes(v)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

//func (c *Companies) getAllCompanyTransactions(db *bolt.DB) []Company {
//}
