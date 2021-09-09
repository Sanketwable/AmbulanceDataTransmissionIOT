#include <ESP8266WiFi.h> 
#include <ESP8266HTTPClient.h> 
const char* ssid = "SSID"; 
const char* password = "Password"; 
const char* host = "http://192.168.0.103"; 
String url = "/update"; 
int adcvalue=0; 

void setup() 
{ 
Serial.begin(115200); 
delay(10); // We start by connecting to a WiFi network 
Serial.println(); 
Serial.println(); Serial.print("Connecting to "); 
Serial.println(ssid);
WiFi.mode(WIFI_STA); 
WiFi.begin(ssid, password); 
while (WiFi.status() != WL_CONNECTED) 
{ 
delay(500); 
Serial.print("."); 
} 
Serial.println(""); 
Serial.println("WiFi connected"); 
Serial.println("IP address: "); 
Serial.println(WiFi.localIP()); } 
int value = 0; 

void loop() 
{ 
delay(100); 
adcvalue= (2*adcvalue + 76) + adcvalue%6;
Serial.print("connecting to "); 
Serial.println(host);
WiFiClient client; 
Serial.print("Requesting URL: "); 
Serial.println(url);

randomSeed(adcvalue);  
String heartrateA1 = String(random(50,60));
String heartrateA2 = String(random(40,65));
String heartrateA3 = String(random(50,65));
String sysBPA1 = String(random(130,150));
String sysBPA2 = String(random(120,160));
String sysBPA3 = String(random(120,170));
String diaBPA1 = String(random(40,60));
String diaBPA2 = String(random(50,70));
String diaBPA3 = String(random(60,80));
String spo2A1 = String(random(90,100));
String spo2A2 = String(random(90,100));
String spo2A3 = String(random(90,100));


String address = host + url; 
HTTPClient http; 
String postDataA1 = "";
String postDataA2 = "";
String postDataA3 = "";
postDataA1 = "{\"id\":1,\"heart_rate\":"+ heartrateA1 + ",\"sys_bp\":" + sysBPA1 + ",\"dia_bp\":" + diaBPA1 + ",\"spo2\":" + spo2A1 + "}";
Serial.println(postDataA1);
postDataA2 = "{\"id\":2,\"heart_rate\":" + heartrateA2 + ",\"sys_bp\":" + sysBPA2 + ",\"dia_bp\":" + diaBPA2 + ",\"spo2\":" + spo2A2 + "}";
Serial.println(postDataA2);
postDataA3 = "{\"id\":3,\"heart_rate\":" + heartrateA3 + ",\"sys_bp\":" + sysBPA3 + ",\"dia_bp\":" + diaBPA3 + ",\"spo2\":" + spo2A3 + "}";
Serial.println(postDataA3);
http.begin(client, "http://192.168.0.103:8080/update"); 
http.addHeader("Content-Type", "application/json"); 
auto httpCode = http.POST(postDataA1);
httpCode = http.POST(postDataA2);
httpCode = http.POST(postDataA3); 
Serial.println(httpCode);
String payload = http.getString(); 
Serial.println(payload); 
http.end();
Serial.println("closing connection"); 
adcvalue++;
} 
