package main

import (
	"os"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ernest_k/defi_stats/liquidation_event"
	"github.com/joho/godotenv"
)

func main() {
	arg1 := os.Args[1]
	err := godotenv.Load()
	lib.Check(err)

	// use arg1 to decide which packages to call
	if arg1 == "1" {
		theTime := time.Date(2023, 8, 3, 0, 0, 0, 0, time.UTC)
		liquidation_event.Log(theTime)
	}
	if arg1 == "2" {
		liquidation_event.LogManager()
	}
}
