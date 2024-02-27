package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ernest_k/defi_stats/liquidation_event"
	"github.com/ernest_k/defi_stats/repositories/redis"
	"github.com/ethereum/go-ethereum/ethclient"
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

	if arg1 == "1" {
		blockNumber := big.NewInt(0).SetUint64(184493969)
		client, err := ethclient.Dial("https://arb-mainnet.g.alchemy.com/v2/s5YvVePEU23eV--i4wUJcxSs4prWmONW")
		fmt.Println("after creating client", err)
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		fmt.Println("after block by number", err)
		fmt.Println(block)
	}
	if arg1 == "2" {
		liquidation_event.LogManager()
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
