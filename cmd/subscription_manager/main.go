package main

import (
	"fmt"
	"os"

	"github.com/ernest_k/defi_stats/influx_repositores"
	"github.com/ernest_k/defi_stats/lib"
	"github.com/ernest_k/defi_stats/liquidation_event"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	arg1 := os.Args[1]
	err := godotenv.Load()
	lib.Check(err)

	if arg1 == "1" {
		liquidation_event.SubscriptionManager()
	}
	if arg1 == "2" {
		client, err := ethclient.Dial(os.Getenv("PLG_MN_ALCH_WS_URL"))
		lib.Check(err)
		addres := common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270")

		tokeEntity := lib.CoinMetadataForAddress(addres, client)
		price := liquidation_event.GetAssetPrice(addres, tokeEntity.Decimal, client, common.HexToAddress(os.Getenv("AAVE3_PLG_MN_ORACLE_ADDRESS")))
		fmt.Println("price: ", price)
	}
	if arg1 == "3" {
		for i := 0; i < 100; i++ {
			influx_repositores.WriteTest()
		}
	}
}
