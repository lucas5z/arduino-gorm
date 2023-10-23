

//#include <ArduinoJson.h>

int sensorPIRPin = 7;
int sensorLuz = A0;
int sensorMag = 2;

void setup() {
  Serial.begin(9600);
  pinMode(sensorPIRPin, INPUT);
  pinMode(sensorLuz, INPUT);
  pinMode(sensorMag, INPUT);
}

void loop() {
  DynamicJsonDocument jsonDoc(200); // Ajusta el tamaño según tus necesidades
  JsonObject root = jsonDoc.to<JsonObject>();

  int movimiento = digitalRead(sensorPIRPin);
  int luz = analogRead(sensorLuz);
  int mag = digitalRead(sensorMag);

  root["id"]=1;
  
  if (mag > 0) {
    root["puerta"] = "cerrado";
  } else {
    root["puerta"] = "abierto";
  }

  if (luz > 140) {
    root["luz"] = "prendido";
  } else {
    root["luz"] = "apagado";
  }

  if (movimiento > 0) {
    root["personas"] = "si";
  } else {
    root["personas"] = "no";
  }

  serializeJson(root, Serial);
  Serial.println(); // Agrega una línea en blanco al final para indicar el final del JSON

  delay(500);
}
