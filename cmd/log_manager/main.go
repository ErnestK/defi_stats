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

	if arg1 == "2" {
		today := time.Now()
		liquidation_event.LogManager(today)
	}
}
