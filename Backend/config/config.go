package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var Data Ambulances
var Status int = -1


type Ambulance struct {
	ID        int      `json:"id"`
	Location  Location `json:"location"`
	HeartRate int      `json:"heart_rate"`
	SysBP     int      `json:"sys_bp"`
	DiaBP     int      `json:"dia_bp"`
	Spo2      int      `json:"spo2"`
}
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Ambulances struct {
	Ambulance []Ambulance
}

func Init() {
	if Status == -1 {
		go OnlinefortwoSec()
	}
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Data)
}

func UpdatedataHandler(w http.ResponseWriter, r *http.Request) {
	if Status == -1 {
		go OnlinefortwoSec()
	}
	ambulance := Ambulance{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &ambulance)
	if err != nil {
		fmt.Println(err)
	}
	AmbulanceID := ambulance.ID

	Data.Ambulance[AmbulanceID-1].DiaBP = ambulance.DiaBP
	Data.Ambulance[AmbulanceID-1].SysBP = ambulance.SysBP
	Data.Ambulance[AmbulanceID-1].HeartRate = ambulance.HeartRate
	Data.Ambulance[AmbulanceID-1].Spo2 = ambulance.Spo2

	WriteToConfig()

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(http.StatusAccepted)
}

func WriteToConfig() {
	file, _ := json.MarshalIndent(Data, "", " ")
	_ = ioutil.WriteFile("config.json", file, 0644)
	Init()
}

func OnlinefortwoSec() {
	Status = 1;
	time.Sleep(11*time.Second)
	Status = -1
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if Status == 1 {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
