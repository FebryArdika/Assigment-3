package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const serverURL = "http://localhost:8080/update"

type Data struct {
	Water float64 `json:"water"`
	Wind  float64 `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		sendUpdate()
		time.Sleep(15 * time.Second)
	}
}

func sendUpdate() {
	water := rand.Float64() * 10 
	wind := rand.Float64() * 20  

	data := Data{Water: water, Wind: wind}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Failed to encode data:", err)
		return
	}

	response, err := http.Post(serverURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Received response:", string(body))
}
