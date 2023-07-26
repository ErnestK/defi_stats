package liquidation_event

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func Subscribe() {
	defer fmt.Printf("Terminated program at: %v\n", time.Now())

	err := godotenv.Load()
	lib.Check(err)
	client, err := ethclient.Dial(os.Getenv("PLG_MN_ALCH_WS_URL"))
	lib.Check(err)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(os.Getenv("AAVE_PLG_MN_POOL_ADDRESS")),
		},
		Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	lib.Check(err)

	fmt.Printf("Started to subscribe at: %v\n", time.Now())
	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("Error at: %v\n", time.Now())
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log at at: %v\n", time.Now())
			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog.Data)

			caMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.CollateralAsset, client)
			daMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.DebtAsset, client)

			caPrice := GetAssetPrice(liquidationCallEventEntity.CollateralAsset, caMetadata.Decimal, client)
			daPrice := GetAssetPrice(liquidationCallEventEntity.DebtAsset, daMetadata.Decimal, client)

			ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)
			fmt.Println("ca name: ", caMetadata.Name)
			fmt.Println("caPrice: ", caPrice)
			fmt.Println("da name: ", daMetadata.Name)
			fmt.Println("daPrice: ", daPrice)

			// liquidationCallEventEntity.CollateralAsset
			// liquidationCallEventEntity.DebtAsset
			// liquidationCallEventEntity.User
			fmt.Println("--------------------------------------------------------")
			// LiquidationCallEventEntityExpanded()
		}
	}
}
