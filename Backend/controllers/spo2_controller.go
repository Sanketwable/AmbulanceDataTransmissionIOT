package controllers

import (
	"ambulancedatatransmissioniot/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Spo2andler is a Func
func Spo2Handler(w http.ResponseWriter, r *http.Request) {
	ambulanceId := r.URL.Query().Get("ambulanceID")
	aid, err := strconv.Atoi(ambulanceId)
	if err != nil {
		fmt.Println(err)
	}
	var r1 [][]int
	spo2 := config.Data.Ambulance[aid-1].Spo2
	row := []int{spo2}
	r1 = append(r1, row)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
