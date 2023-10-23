package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lucas5z/arduino1/db"
	"github.com/lucas5z/arduino1/models"
	"github.com/lucas5z/arduino1/routes"
)

func main() {
	//routes.Open2()
	//db
	db.Conex()
	//db migracion
	db.DB.AutoMigrate(&models.Datos{})

	r := mux.NewRouter()

	// Ruta para obtener el JSON del usuario
	r.HandleFunc("/prueba", routes.Get_arduinio).Methods("GET")
	r.HandleFunc("/prueba", routes.Post_arduino).Methods("POST")
	go Put_arduino_time() //PUT

	// Ruta para servir el archivo HTML del frontend
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	//port := "192.168.1.58:8000"
	port := ":8000"
	fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func Put_arduino_time() {
	for {
		routes.Put_arduino(nil, nil)

		time.Sleep(500 * time.Millisecond)
	}
}
