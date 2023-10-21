package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tarm/serial"
)

type User struct {
	Puerta   string `json:"puerta"`
	Lus      string `json:"luz"`
	Personas string `json:"persona"`
}

type Datos struct {
	Puerta   string `json:"puerta"`
	Luz      string `json:"luz"`
	Personas string `json:"personas"`
}

func GetUserJSON(w http.ResponseWriter, r *http.Request) {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		http.Error(w, "No se pudo abrir el archivo JSON", http.StatusInternalServerError)
		return
	}
	defer s.Close()

	var dato User

	err = json.NewDecoder(s).Decode(&dato)
	if err != nil {
		http.Error(w, "No se pudo decodificar el JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato)
}

func main() {
	//go Open2()
	r := mux.NewRouter()

	// Ruta para obtener el JSON del usuario
	r.HandleFunc("/prueba", GetUserJSON1).Methods("GET")

	// Ruta para servir el archivo HTML del frontend
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	// Inicia el servidor en el puerto 8080
	//port := "192.168.1.58:8000"
	port := ":8000"
	fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func Open2() {

	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	var datos Datos

	decoder := json.NewDecoder(s)

	for {
		err := decoder.Decode(&datos)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(datos)
	}
}

func GetUserJSON1(w http.ResponseWriter, r *http.Request) {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		http.Error(w, "No se pudo abrir el archivo JSON", http.StatusInternalServerError)
		return
	}
	defer s.Close()

	var dato Datos

	err = json.NewDecoder(s).Decode(&dato)
	if err != nil {
		http.Error(w, "No se pudo decodificar el JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato)
}
