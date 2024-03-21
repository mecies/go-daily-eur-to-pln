package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ResponseObject struct {
	Date string `json:"date"`
	Euro Euro   `json:"eur"`
}

type Euro struct {
	Pln float64 `json:"pln"`
}

func GetEuroToPlnExchangeRate() float64 {
	response, err := http.Get("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.min.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject ResponseObject
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Euro.Pln
}
