package controllers

import (
	"ambulancedatatransmissioniot/config"
	"ambulancedatatransmissioniot/responses"
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
	responses.JSON(w, 100, r1)
}
