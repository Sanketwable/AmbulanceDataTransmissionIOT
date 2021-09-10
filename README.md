# AmbulanceDataTransmissionIOT
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
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

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://wablesanket.xyz"><img src="https://avatars.githubusercontent.com/u/43716242?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Wable Sanket</b></sub></a><br /><a href="#infra-Sanketwable" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/Sanketwable/AmbulanceDataTransmissionIOT/commits?author=Sanketwable" title="Tests">⚠️</a> <a href="https://github.com/Sanketwable/AmbulanceDataTransmissionIOT/commits?author=Sanketwable" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/bhanu-1312"><img src="https://avatars.githubusercontent.com/u/57167640?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Bhanu Pratap Singh</b></sub></a><br /><a href="#infra-bhanu-1312" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!