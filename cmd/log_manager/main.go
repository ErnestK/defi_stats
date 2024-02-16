package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ernest_k/defi_stats/liquidation_event"
	"github.com/ernest_k/defi_stats/repositories/redis"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program_name arg1")
		fmt.Println("Please provide the required argument.")
		return
	}

	arg1 := os.Args[1]
	err := godotenv.Load()
	lib.Check(err)

	if arg1 == "2" {
		today := time.Now()
		liquidation_event.LogManager(today)
	}
	if arg1 == "3" {
		n := int64(23)
		redis.WriteLastProcessedBlock("er1", n)
	}
	if arg1 == "4" {
		n1 := redis.GetLastProcessedBlock(1, "er1")
		fmt.Println("latestBlock ", n1)
	}
}
