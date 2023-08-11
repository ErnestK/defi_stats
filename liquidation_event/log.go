package liquidation_event

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const STEP_FOR_LOG_BLOCK = 10_000

func Log(fromTime time.Time, eventBundle EventBundle) {
	timestamp := time.Date(2042, 8, 7, 0, 0, 0, 0, time.UTC)
	err := godotenv.Load()
	lib.Check(err)

	client, err := ethclient.Dial(eventBundle.connectUrl)
	lib.Check(err)

	latestBlockTemp, err := client.BlockNumber(context.Background())
	lib.Check(err)
	latestBlock := int64(latestBlockTemp)

	for {
		startWithBlock := latestBlock - STEP_FOR_LOG_BLOCK

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(latestBlock - STEP_FOR_LOG_BLOCK),
			ToBlock:   big.NewInt(latestBlock),
			Addresses: []common.Address{
				eventBundle.poolAddress,
			},
			Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
		}
		eventLogs, err := client.FilterLogs(context.Background(), query)
		lib.Check(err)

		for _, vLog := range eventLogs {
			timestamp = lib.ShowBlockDate(big.NewInt(int64(vLog.BlockNumber)), client)
			fmt.Println("block created at:", timestamp)

			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog)
			fmt.Println("for chain: ", eventBundle.chainName)
			ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)
			caMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.CollateralAsset, client)
			daMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.DebtAsset, client)

			caPrice := GetAssetPrice(liquidationCallEventEntity.CollateralAsset, caMetadata.Decimal, client, eventBundle.oracleAddress)
			daPrice := GetAssetPrice(liquidationCallEventEntity.DebtAsset, daMetadata.Decimal, client, eventBundle.oracleAddress)

			fmt.Println("ca name: ", caMetadata.Name)
			fmt.Println("caPrice: ", caPrice)
			fmt.Println("da name: ", daMetadata.Name)
			fmt.Println("daPrice: ", daPrice)
		}
		if timestamp.Before(fromTime) {
			return
		}
		// start with last again
		latestBlock = startWithBlock
	}
}
