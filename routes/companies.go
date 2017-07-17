package routes

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"github.com/jesseobrien/exchange/config"
	"github.com/jesseobrien/exchange/models"
)

var db *bolt.DB

func Companies(r *mux.Router, cfg *config.Config) {
	db = cfg.Database

	r.HandleFunc("/companies", createCompany).Methods("POST")
	r.HandleFunc("/companies/{symbol}", getCompanyById).Methods("GET")

}

func createCompany(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	company := models.NewCompany()

	json.NewDecoder(req.Body).Decode(company)
	company.Save(db)

	json.NewEncoder(w).Encode(company)
}

func getCompanyById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)
	symbol := vars["symbol"]

	company := &models.Company{}
	company.GetById(db, symbol)

	if company.Symbol == "" {
		http.Error(w, "", 404)
		return
	}

	json.NewEncoder(w).Encode(company)
}
