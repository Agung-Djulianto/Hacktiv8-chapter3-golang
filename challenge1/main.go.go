package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)

	for t := range ticker.C {
		fmt.Println(t)
		fmt.Println(strings.Repeat("#", 25))
		postJsonPlaceHolder()
		randomNumber()
	}
}

func postJsonPlaceHolder() {
	randUserInt := rand.Intn(10)
	data := map[string]interface{}{
		"title":  "Airell",
		"body":   "Jordan",
		"userId": randUserInt,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(requestJson))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func randomNumber() {
	randIntWater, randIntWind := rand.Intn(100), rand.Intn(100)
	statusWater, statusWind := "aman", "aman"

	if randIntWater >= 6 && randIntWater <= 8 {
		statusWater = "siaga"
	} else if randIntWater > 8 {
		statusWater = "bahaya"
	}

	if randIntWind >= 7 && randIntWind <= 15 {
		statusWind = "siaga"
	} else if randIntWind > 8 {
		statusWind = "bahaya"
	}

	fmt.Printf(`{"water": %d, "wind":%d}`, randIntWater, randIntWind)
	fmt.Printf("\n")
	fmt.Println("status water :", statusWater)
	fmt.Println("status wind :", statusWind)
}
