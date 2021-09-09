package controllers

import (
	"ambulancedatatransmissioniot/config"
	"ambulancedatatransmissioniot/responses"
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
	var r1 [][]float64
	row := []float64{config.Data.Ambulance[aid-1].Location.Latitude, config.Data.Ambulance[aid-1].Location.Longitude}
	r1 = append(r1, row)
	responses.JSON(w, 100, r1)

}
