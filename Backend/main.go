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
// Location is a struct
type Location struct {
	latitude float64
	longitude float64
}

// AmbulanceLocation Stores location of thre Ambulance
var AmbulanceLocation = [3]Location{
	{
		latitude: 19.106207,
		longitude: 74.621587,
	},
	{
		latitude: 19.126207,
		longitude: 74.821587,
	},
	{
		latitude: 19.176207,
		longitude: 74.921587,
	},
}

var location = Location{
	latitude: 19.106207,
	longitude: 74.621587,
}

func main() {
	fmt.Println("Starting Server : 8080")
	r := mux.NewRouter()
	r.HandleFunc("/heartrate", HeartRateHandler).Methods(http.MethodGet)
	r.HandleFunc("/bloodpressure", BloodPressureHandler).Methods(http.MethodGet)
	r.HandleFunc("/spo2", Spo2Handler).Methods(http.MethodGet)
	r.HandleFunc("/location", GetambulanceLocation).Methods(http.MethodGet)
	
	http.ListenAndServe(":8080", r)

}
// Spo2andler is a Func
func Spo2Handler(w http.ResponseWriter, r *http.Request) {
	ambulanceId := r.URL.Query().Get("ambulanceID")
	fmt.Println(ambulanceId)
	var r1 [][]int
	x1 := rand.NewSource(time.Now().UnixNano())
	spo2 := rand.New(x1).Intn(200) % 10 +	90
	row := []int{spo2}
	r1 = append(r1, row)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// w.Header().Add("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
func GetambulanceLocation(w http.ResponseWriter, r *http.Request) {
	ambulanceId := r.URL.Query().Get("ambulanceID")
	aid, err := strconv.Atoi(ambulanceId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ambulanceId)
	var r1 [][]float64
	row := []float64{AmbulanceLocation[aid-1].latitude, AmbulanceLocation[aid-1].longitude}
	r1 = append(r1, row)
	fmt.Println("value of air is :", aid)
	fmt.Println("data sent = \n", r1)
	
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

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
		systolic := rand.New(x1).Intn(200) % 20 +	130
		diastolic := rand.New(x1).Intn(200) % 20 +	50
		row := []int{i, systolic, diastolic}
		r1 = append(r1, row)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	json.NewEncoder(w).Encode(r1)
	w.WriteHeader(100)

}
