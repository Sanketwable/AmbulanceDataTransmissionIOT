package controllers

import (
	"ambulancedatatransmissioniot/config"
	"ambulancedatatransmissioniot/responses"
	"fmt"
	"net/http"
	"strconv"
)

func BloodPressureHandler(w http.ResponseWriter, r *http.Request) {
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
	systolic := config.Data.Ambulance[aid-1].SysBP
	diastolic := config.Data.Ambulance[aid-1].DiaBP
	row := []int{strt, systolic, diastolic}
	r1 = append(r1, row)
	responses.JSON(w, 100, r1)
}
