// Map to store AmbulanceID and AmbulanceNumber
const AmbulanceInfo = new Map();
AmbulanceInfo.set(1, "MH01 A5003")
AmbulanceInfo.set(2, "MH06 G1463")
AmbulanceInfo.set(3, "MH08 E0091")
var spo2Info = 95;
var selectedAmbulance = 1;
var originLatitude =  19.106207;
var originLongitude = 74.621587;
var destinationLatitude = 19.1178 
var destinationLongitude = 74.7300

// updateAmbulanceLocation();

function initMap() {
    console.log(originLatitude)
    console.log(originLongitude)
    const directionsRenderer = new google.maps.DirectionsRenderer();
    const directionsService = new google.maps.DirectionsService();
    const map = new google.maps.Map(document.getElementById("map"), {
      zoom: 14,
      center: { lat: originLatitude, lng: originLatitude },
    });
    directionsRenderer.setMap(map);
    calculateAndDisplayRoute(directionsService, directionsRenderer);
}
  
  function calculateAndDisplayRoute(directionsService, directionsRenderer) {
    // const selectedMode = document.getElementById("mode").value;
    directionsService
      .route({
        origin: { lat: originLatitude, lng: originLongitude },
        destination: { lat: 19.1178, lng: 74.7300 },
        travelMode: google.maps.TravelMode["DRIVING"],
      })
      .then((response) => {
        directionsRenderer.setDirections(response);
      })
      .catch((e) => window.alert("Directions request failed due to " + status));
}

function addLocationData(data3) {
    console.log(data3[0][0]);
    console.log(data3[0][1]);
    originLatitude =  data3[0][0];
    originLongitude = data3[0][1];
}

function updateAmbulanceLocation(amb) {
    $.getJSON("http://localhost:8080/location?ambulanceID="+amb, addLocationData);
}

// SelectAmbulance is selecting ambulance from drop down
function SelectAmbulance() {
    var e = document.getElementById("Ambulance");
    var selectedAmbulance = e.options[e.selectedIndex].value;
	document.getElementById("AmbulanceINFO").innerHTML = "Ambulance " + selectedAmbulance + " : " +AmbulanceInfo.get(parseInt(selectedAmbulance));
    console.log(originLatitude)
    console.log(originLongitude)
    setTimeout(() => {  updateAmbulanceLocation(selectedAmbulance); }, 1000);;
    setTimeout(() => {  initMap(); }, 2000);
}

window.onload = function() {

    var HeartBeatPoints = [];
    var systolicBP = [];
    var avgsystolicBP = [];
    var averageheartrate = [];
    var diastolicBP = [];
    var avgdiastolicBP = [];

    var heartRateChart = new CanvasJS.Chart("heartRateMonitor", {

	    theme: "light2",
	    animationEnabled: true,
	    title: {
		    text: "HeartBeat Data",
	    },
        axisY: { 
		    title: "Rate",
		    suffix: " bpm",
		    crosshair: {
			    enabled: true
		    }
	    },
        axisX: { 
		    title: "Time",
		    suffix: " sec",
		    crosshair: {
			    enabled: true,
			    snapToDataPoint: true
		    }
	    },
	    toolTip:{
		    shared:true
	    },  
	    data: [{
		    type: "line",
		    showInLegend: true,
		    name: "Bpm",
		    markerType: "circle",
		    color: "#F08080",
		    dataPoints: HeartBeatPoints,
	    },{
		    type: "line",
		    showInLegend: true,
            markerSize: 0,
		    name: "average Bpm",
		    lineDashType: "dash",
		    dataPoints: averageheartrate,
	    }]
    });
    updateHeartRateData();

    var bloodpressureChart = new CanvasJS.Chart("bloodPressureMonitor", {
	    theme: "light2",
        animationEnabled: true,
	    title: {
		    text: "BloodPressure Data"
	    },
        axisY: { 
		    title: "Hg",
		    suffix: "mm",
		    crosshair: {
			    enabled: true
		    }
	    },
        axisX: { 
		    title: "Time",
		    suffix: " sec",
		    crosshair: {
			    enabled: true,
			    snapToDataPoint: true
		    }
	    },
	    toolTip:{
		    shared:true
	    }, 
	    data: [{
		    type: "line",
            showInLegend: true,
            markerType: "square",
            name: "Systolic",
            color: "#F08080",
		    dataPoints: systolicBP
	    },{
		    type: "line",
		    showInLegend: true,
		    name: "Diastolic",
		    lineDashType: "line",
		    dataPoints: diastolicBP,
	    },{
		    type: "line",
		    showInLegend: true,
            markerSize: 0,
		    name: "average Systolic",
		    lineDashType: "dash",
		    dataPoints: avgsystolicBP,
	    },{
		    type: "line",
		    showInLegend: true,
            markerSize: 0,
		    name: "average Diastolic",
		    lineDashType: "dash",
		    dataPoints: avgdiastolicBP,
	    }]
    });
    updateBloodPressureData();
    updateSpO2();

    // Initial Values
    var xValueHR = 0;
    var yValue = 10;
    var newDataCount = 6;

    var xValueBP = 0;
    var yValue1 = 10;
    var newDataCount1 = 6;

    function addHeartRateData(data) {
	    if(newDataCount != 1) {
		    $.each(data, function(key, value) {
			    HeartBeatPoints.push({x: value[0], y: parseInt(value[1])});
			    averageheartrate.push({x:value[0], y:55});
			    xValueHR++;
			    yValue = parseInt(value[1]);
		    });
	    } else {
		    //dataPoints.shift();
		    HeartBeatPoints.push({x: data[0][0], y: parseInt(data[0][1])});
		    averageheartrate.push({x:data[0][0], y:55});
		    xValueHR++;
		    yValue = parseInt(data[0][1]);
	    }
	
	    newDataCount = 1;
        var dataLength = 20;
        if (HeartBeatPoints.length > dataLength) {
		    HeartBeatPoints.shift();
		    averageheartrate.shift();
	    }
	    heartRateChart.render();
    
	    setTimeout(updateHeartRateData, 2000);
    }
    function addBloodPressureData(data1) {
	    if(newDataCount1 != 1) {
		    $.each(data1, function(key, value) {
			    systolicBP.push({x: value[0], y: parseInt(value[1])});
                avgsystolicBP.push({x:value[0], y:120})
                diastolicBP.push({x: value[0], y: parseInt(value[2])});
                avgdiastolicBP.push({x:value[0], y:80})
			    xValueBP++;
			    yValue1 = parseInt(value[1]);
		    });
	    } else {
		    systolicBP.push({x: data1[0][0], y: parseInt(data1[0][1])});
            avgsystolicBP.push({x:data1[0][0], y:120})
            diastolicBP.push({x: data1[0][0], y: parseInt(data1[0][2])});
            avgdiastolicBP.push({x:data1[0][0], y:80})
		    xValueBP++;
		    yValue1 = parseInt(data1[0][1]);
	    }
	    var dataLength = 20;
	    newDataCount1 = 1;
        if (systolicBP.length > dataLength) {
		    systolicBP.shift();
            diastolicBP.shift();
            avgdiastolicBP.shift();
            avgsystolicBP.shift();
	    }
	    bloodpressureChart.render();
	    setTimeout(updateBloodPressureData, 2000);
    }
    function addspo2Data(data2) {
		spo2Info =  data2[0][0];
	    setTimeout(updateSpO2, 10000);
        document.getElementById("spo2info").innerHTML = "SpO2 : " + spo2Info;
    }

    function updateHeartRateData() {
        $.getJSON("http://localhost:8080/heartrate?xstart="+xValueHR+"&ambulanceID="+selectedAmbulance, addHeartRateData);
    }
    function updateBloodPressureData() {
        $.getJSON("http://localhost:8080/bloodpressure?xstart="+xValueBP+"&ambulanceID="+selectedAmbulance, addBloodPressureData);
    }
    function updateSpO2() {
        $.getJSON("http://localhost:8080/spo2?ambulanceID="+selectedAmbulance, addspo2Data);
    }


}   


