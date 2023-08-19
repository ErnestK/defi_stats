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
	fmt.Println("------------------------------------------------------------")
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
			FromBlock: big.NewInt(startWithBlock),
			ToBlock:   big.NewInt(latestBlock),
			Addresses: []common.Address{
				eventBundle.poolAddress,
			},
			Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
		}
		eventLogs, err := client.FilterLogs(context.Background(), query)
		lib.Check(err)

		for _, vLog := range eventLogs {
			blockNumber := big.NewInt(int64(vLog.BlockNumber))
			timestamp = lib.ShowBlockDate(blockNumber, client)
			fmt.Println("block created at:", timestamp)
			fmt.Println("block num:", blockNumber)

			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog)
			fmt.Println("for chain: ", eventBundle.chainName)
			ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)

			caMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.CollateralAsset, client)
			daMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.DebtAsset, client)

			caPrice := GetAssetPriceAaveV3(liquidationCallEventEntity.CollateralAsset, client, eventBundle.oracleAddress)
			daPrice := GetAssetPriceAaveV3(liquidationCallEventEntity.DebtAsset, client, eventBundle.oracleAddress)

			fmt.Println("ca name: ", caMetadata.Name)
			fmt.Println("caPrice: ", caPrice)
			fmt.Println("da name: ", daMetadata.Name)
			fmt.Println("daPrice: ", daPrice)
			entity := LiquidationCallEventEntityWithMeta{
				liquidationCallEventEntity: liquidationCallEventEntity,
				caInUsd:                    caPrice,
				daInUsd:                    daPrice,
				chainName:                  eventBundle.chainName,
				caName:                     caMetadata.Name,
				daName:                     daMetadata.Name,
				timestamp:                  time.Now(),
			}
			fmt.Println("full event: ", entity)
		}
		if timestamp.Before(fromTime) {
			return
		}
		// start with last again
		latestBlock = startWithBlock
	}
}
