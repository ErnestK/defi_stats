package liquidation_event

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	influx "github.com/ernest_k/defi_stats/influx_repositores"
	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const STEP_FOR_LOG_BLOCK = 10_000

func Log(fromTime time.Time, eventBundle EventBundle) {
	fmt.Println("------------------------------------------------------------")
	dbConnection := influx.CreateDB()
	defer dbConnection.CloseDB()

	var wgWriteToDB sync.WaitGroup

	var resultToWrite []LiquidationCallEventEntityWithMeta
	timestamp := time.Date(2042, 8, 7, 0, 0, 0, 0, time.UTC)

	err := godotenv.Load()
	lib.Check(err)

	client, err := ethclient.Dial(eventBundle.connectUrl)
	lib.CheckWithMessage(err, "cannot connect to url: "+eventBundle.connectUrl)

	latestBlockTemp, err := client.BlockNumber(context.Background())
	lib.CheckWithMessage(err, "cannot get block number")
	latestBlock := int64(latestBlockTemp)

	for {
		resultToWrite = []LiquidationCallEventEntityWithMeta{}
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
		lib.CheckWithMessage(err, "cannot execute query on chain:"+eventBundle.chainName)

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
			resultToWrite = append(resultToWrite, entity)
		}
		if timestamp.Before(fromTime) {
			writeToDB(fromTime, resultToWrite, dbConnection, &wgWriteToDB)
			wgWriteToDB.Wait()
			return
		}
		// start with last again
		writeToDB(fromTime, resultToWrite, dbConnection, &wgWriteToDB)
		latestBlock = startWithBlock
	}
}

func writeToDB(fromTime time.Time, resultToWrite []LiquidationCallEventEntityWithMeta, dbConnection *influx.Connection, wgWriteToDB *sync.WaitGroup) {
	wgWriteToDB.Add(1)
	go writesToDBWithoutPrice(resultToWrite, dbConnection)

	twoHoursAgo := time.Now().Add(-2 * time.Hour)
	if fromTime.After(twoHoursAgo) {
		wgWriteToDB.Add(1)
		go writesToDBWithPrice(resultToWrite, dbConnection)
	}
}

func writesToDBWithoutPrice(resultToWrite []LiquidationCallEventEntityWithMeta, dbConnection *influx.Connection) {
	for _, resultEntity := range resultToWrite {
		dbConnection.WriteEventLogWoPrice(resultEntity)
	}
}

func writesToDBWithPrice(resultToWrite []LiquidationCallEventEntityWithMeta, dbConnection *influx.Connection) {
	for _, resultEntity := range resultToWrite {
		dbConnection.WriteEventLog(resultEntity)
	}
}
