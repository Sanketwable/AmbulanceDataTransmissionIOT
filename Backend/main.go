package main

import (
	"ambulancedatatransmissioniot/config"
	"ambulancedatatransmissioniot/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Server : 8080")
	fmt.Println("Initializing Database")
	config.Init()
	r := mux.NewRouter()
	r.HandleFunc("/heartrate", controllers.HeartRateHandler).Methods(http.MethodGet)
	r.HandleFunc("/bloodpressure", controllers.BloodPressureHandler).Methods(http.MethodGet)
	r.HandleFunc("/spo2", controllers.Spo2Handler).Methods(http.MethodGet)
	r.HandleFunc("/location", controllers.GetambulanceLocation).Methods(http.MethodGet)
	r.HandleFunc("/update", config.UpdatedataHandler).Methods(http.MethodPost)
	r.HandleFunc("/status", controllers.StatusHandler).Methods(http.MethodGet)
	

	http.ListenAndServe(":8080", r)

}
