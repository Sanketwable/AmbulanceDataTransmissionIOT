package controllers

import (
	"ambulancedatatransmissioniot/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// HeartRateHandler is a func
func HeartRateHandler(w http.ResponseWriter, r *http.Request) {
	ambulanceId := r.URL.Query().Get("ambulanceID")
	aid, err := strconv.Atoi(ambulanceId)
	if err != nil {
		fmt.Println(err)
	}

	start := r.URL.Query().Get("xstart")
	strt, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println(err)
	}
	var r1 [][]int
	y := config.Data.Ambulance[aid-1].HeartRate
	row := []int{strt, y}
	r1 = append(r1, row)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
