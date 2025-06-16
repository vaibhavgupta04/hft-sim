package main

import (
	"fmt"
	"github.com/vaibhavgupta04/marketdatafeeder/client"
	"github.com/vaibhavgupta04/marketdatafeeder/stores"

)

func main() {
	finnhubclient :=client.NewFinnhubConnector()
	client.NewMarketStackConnector()
	redisclient := stores.NewRedisClient()
	fmt.Println("Redis client initialized:", redisclient != nil)
	fmt.Println("Finnhub client initialized:", finnhubclient != nil)
	

}