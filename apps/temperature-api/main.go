package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	mux.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		location := r.FormValue("location")
		sensorID := getSensorIDByLocation(location)

		val := float64(rand.Int32N(100)) + rand.Float64()
		data := TemperatureResponse{
			Value:       val,
			Unit:        "C",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "OK",
			SensorID:    sensorID,
			SensorType:  "warm",
			Description: "well describe sensor data",
		}

		resp, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	})

	mux.HandleFunc("/temperature/{sensorID}", func(w http.ResponseWriter, r *http.Request) {
		sensorID := r.PathValue("sensorID")
		location := getLocationBySensorID(sensorID)

		val := float64(rand.Int32N(100)) + rand.Float64()
		data := TemperatureResponse{
			Value:       val,
			Unit:        "C",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "OK",
			SensorID:    sensorID,
			SensorType:  "warm",
			Description: "well describe sensor data",
		}

		resp, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	})

	log.Fatal(http.ListenAndServe(":8081", mux))
}

// TemperatureResponse represents the response from the temperature API
type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func getLocationBySensorID(id string) string {
	switch id {
	case "1":
		return "Living Room"
	case "2":
		return "Bedroom"
	case "3":
		return "Kitchen"
	default:
		return "Unknown"
	}
}

func getSensorIDByLocation(location string) string {
	switch location {
	case "Living Room":
		return "1"
	case "Bedroom":
		return "2"
	case "Kitchen":
		return "3"
	default:
		return "0"
	}
}
