# AmbulanceDataTransmissionIOT
# Aim
This Project aims to transfer real time critical informations of patient in an ambulance to the doctor.

It sends `heartbeatrate`, `bloodpressure`, `oxygen level` and `location of ambulance`.

# How it works
* It uses a NodeMCU micro controller which recieves data from the machines installed in ambulance. 
* This micro controller changes the database in realtime for the corresponding ambulance.
* The server fetches the critical information from the database and sends it back to requesting client in real time.
* The client side javascript recieves the data and displays it in the form of graph(chart).

# Technology Used
* CanvasJS API is used to display information in the form of charts.
* JQuery is used for building dynamic web page.
* Google Maps API is used for displaying location of ambulance in a way that can be easily comprehended by the client.
* NodeMCU controller is used to send patient data to the server using JSON.
* Server is built using Golang programming language which has inbuilt concurrency features to handle real time critical requests.
