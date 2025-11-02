package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var city string
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("No .env loaded: %v", err)
		return
	}
	apiKey := os.Getenv("OWM_API_KEY")
	fmt.Println("Enter the city you want to know the weather of")
	fmt.Scan(&city)

	urlStr, err := buildWeatherURL(city, apiKey)
	if err != nil {
		fmt.Printf("No .env loaded: %v", err)
		return
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Printf("No .env loaded: %v", err)
		return
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("No .env loaded: %v", err)
		return
	}
	fmt.Println(string(b))

}
