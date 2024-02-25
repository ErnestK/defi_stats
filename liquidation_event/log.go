package liquidation_event

import (
	"context"
	"math/big"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ernest_k/defi_stats/repositories/influx"
	"github.com/ernest_k/defi_stats/repositories/redis"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

const STEP_FOR_LOG_BLOCK = 10_000

func Log(fromTime time.Time, eventBundle EventBundle) {
	dbConnection := influx.CreateDB()
	defer dbConnection.CloseDB()

	var resultToWrite []LiquidationCallEventEntityWithMeta

	err := godotenv.Load()
	lib.Check(err)

	client, err := ethclient.Dial(eventBundle.connectUrl)
	lib.CheckWithMessage(err, "cannot connect to url: "+eventBundle.connectUrl)

	latestBlockTemp, err := client.BlockNumber(context.Background())
	lib.CheckWithMessage(err, "cannot get block number")
	latestBlock := int64(latestBlockTemp)

	resultToWrite = []LiquidationCallEventEntityWithMeta{}
	startWithBlock := redis.GetLastProcessedBlock(latestBlock, eventBundle.chainName)

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
		blockNumber := big.NewInt(0).SetUint64(vLog.BlockNumber)
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		lib.Check(err)
		timestamp := time.Unix(int64(block.Time()), 0)

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
			timestamp:                  timestamp,
		}
		resultToWrite = append(resultToWrite, entity)
	}

	writeToDB(fromTime, resultToWrite, dbConnection)
	redis.WriteLastProcessedBlock(eventBundle.chainName, latestBlock)
}

func writeToDB(fromTime time.Time, resultToWrite []LiquidationCallEventEntityWithMeta, dbConnection *influx.Connection) {
	for _, resultEntity := range resultToWrite {
		dbConnection.WriteEventLog(resultEntity)
	}
}
