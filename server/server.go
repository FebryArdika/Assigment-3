package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Water       float64 `json:"water"`
	Wind        float64 `json:"wind"`
	WaterStatus string  `json:"water_status"`
	WindStatus  string  `json:"wind_status"`
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	data.WaterStatus = getWaterStatus(data.Water)
	data.WindStatus = getWindStatus(data.Wind)

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode response data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getWaterStatus(water float64) string {
	if water < 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func getWindStatus(wind float64) string {
	if wind < 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func main() {
	http.HandleFunc("/update", updateHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
