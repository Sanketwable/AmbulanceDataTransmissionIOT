package controllers

import (
	"ambulancedatatransmissioniot/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetambulanceLocation(w http.ResponseWriter, r *http.Request) {
	ambulanceId := r.URL.Query().Get("ambulanceID")
	aid, err := strconv.Atoi(ambulanceId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ambulanceId)
	var r1 [][]float64
	row := []float64{config.Data.Ambulance[aid-1].Location.Latitude, config.Data.Ambulance[aid-1].Location.Longitude}
	r1 = append(r1, row)
	
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
