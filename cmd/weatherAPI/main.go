package main

import (
	"fmt"
	"log"
	"weatherAPI/internal/weatherClient"
)

func main() {
	client := weatherClient.NewClient()

	weather, err := client.GetWeather("Kiev")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(weather.Info())
}
