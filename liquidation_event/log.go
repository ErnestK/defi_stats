package liquidation_event

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const STEP_FOR_LOG_BLOCK = 1_000_000

func Log(fromTime time.Time) {
	err := godotenv.Load()
	lib.Check(err)
	endpoint := os.Getenv("PLG_MN_ALCH_HTTPS_URL")

	// block with the event is the closest for the new year 12.2022
	// polygonBlockWithEvent := int64(37446903)

	client, err := ethclient.Dial(endpoint)
	lib.Check(err)

	latestBlockTemp, err := client.BlockNumber(context.Background())
	lib.Check(err)
	latestBlock := int64(latestBlockTemp)

	for {
		startWithBlock := latestBlock - STEP_FOR_LOG_BLOCK
		fmt.Println("from block:", startWithBlock)
		fmt.Println("to block:", latestBlock)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(latestBlock - STEP_FOR_LOG_BLOCK),
			ToBlock:   big.NewInt(latestBlock),
			Addresses: []common.Address{
				common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
			},
			Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
		}
		eventLogs, err := client.FilterLogs(context.Background(), query)
		lib.Check(err)

		fmt.Println("All Logs length: ", len(eventLogs))

		for _, vLog := range eventLogs {
			// liquidationCallEventEntity := LiquidationCallEventEntity{}
			// DeserializeEventLog(&liquidationCallEventEntity, vLog.Data)
			// ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)

			// lib.ShowBlockDate(big.NewInt(int64(polygonBlockWithEvent)), client)
			timestamp := lib.ShowBlockDate(big.NewInt(int64(vLog.BlockNumber)), client)
			fmt.Println("block created at:", timestamp)
			if timestamp.Before(fromTime) {
				return
			}
		}
		// start with last again
		latestBlock = startWithBlock
	}
}
