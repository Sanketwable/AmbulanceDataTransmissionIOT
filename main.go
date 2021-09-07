package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Server : 8080")
	r := mux.NewRouter()
	r.HandleFunc("/heartrate", HeartRateHandler).Methods(http.MethodGet)
	r.HandleFunc("/bloodpressure", BloodPressureHandler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)

}

// HeartRateHandler is a func
func HeartRateHandler(w http.ResponseWriter, r *http.Request) {

	start := r.URL.Query().Get("xstart")
	strt, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println(err)
	}
	var r1 [][]int

	for i := strt; i < strt+10; i++ {
		x1 := rand.NewSource(time.Now().UnixNano())
		y := rand.New(x1).Intn(200) % 10 +	50
		row := []int{i, y}
		r1 = append(r1, row)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
func BloodPressureHandler(w http.ResponseWriter, r *http.Request) {

	start := r.URL.Query().Get("xstart")
	strt, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println(err)
	}
	var r1 [][]int

	for i := strt; i < strt+10; i++ {
		x1 := rand.NewSource(time.Now().UnixNano())
		y := rand.New(x1).Intn(200) % 10 +	50
		row := []int{i, y}
		r1 = append(r1, row)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
