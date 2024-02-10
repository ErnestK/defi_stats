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

			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog)

			caMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.CollateralAsset, client)
			daMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.DebtAsset, client)

			caPrice := GetAssetPriceAaveV3(liquidationCallEventEntity.CollateralAsset, client, eventBundle.oracleAddress)
			daPrice := GetAssetPriceAaveV3(liquidationCallEventEntity.DebtAsset, client, eventBundle.oracleAddress)

			entity := LiquidationCallEventEntityWithMeta{
				liquidationCallEventEntity: liquidationCallEventEntity,
				caInUsd:                    caPrice,
				daInUsd:                    daPrice,
				chainName:                  eventBundle.chainName,
				caMetadata:                 caMetadata,
				daMetadata:                 daMetadata,
				timestamp:                  time.Now(),
			}
			resultToWrite = append(resultToWrite, entity)
		}
		if timestamp.Before(fromTime) {
			wgWriteToDB.Add(1)
			go writeToDB(fromTime, resultToWrite, dbConnection, &wgWriteToDB)
			break
		}
		wgWriteToDB.Add(1)
		go writeToDB(fromTime, resultToWrite, dbConnection, &wgWriteToDB)
		// start with last again
		latestBlock = startWithBlock
	}

	wgWriteToDB.Wait()
}

func writeToDB(fromTime time.Time, resultToWrite []LiquidationCallEventEntityWithMeta, dbConnection *influx.Connection, wgWriteToDB *sync.WaitGroup) {
	for _, resultEntity := range resultToWrite {
		dbConnection.WriteEventLog(resultEntity)
	}
	wgWriteToDB.Done()
}
