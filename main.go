package main

import (
	"PayMe/pay"
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	//register the handler for the /computeIncomeDetail endpoint
	http.HandleFunc("/computeIncomeDetail", ComputeIncomeDetailHandler)
	//start the server
	http.ListenAndServe(":5000", nil)
}

func ComputeIncomeDetailHandler(w http.ResponseWriter, r *http.Request) {
	//parse the request
	workedDays, err := strconv.Atoi(r.URL.Query().Get("workedDays"))
	if err != nil {
		http.Error(w, "Invalid workedDays", http.StatusBadRequest)
		return
	}
	//parse the dailyRate request parameter
	dailyRate, err := strconv.ParseFloat(r.URL.Query().Get("dailyRate"), 64)
	if err != nil {
		http.Error(w, "Invalid dailyRate", http.StatusBadRequest)
		return

	}
	//parse float parameter called taxRate
	taxRate, err := strconv.ParseFloat(r.URL.Query().Get("taxRate"), 64)
	if err != nil {
		http.Error(w, "Invalid taxRate", http.StatusBadRequest)
		return
	}

	incomeDetail := pay.ComputeIncomeDetail(workedDays, dailyRate, taxRate)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomeDetail)
}
