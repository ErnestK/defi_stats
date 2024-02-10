package liquidation_event

import (
	"context"
	"fmt"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Subscribe(eventBundle EventBundle) {
	defer eventBundle.waitGroup.Done()

	client, err := ethclient.Dial(eventBundle.connectUrl)
	lib.Check(err)
	defer client.Close()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			eventBundle.poolAddress,
		},
		Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	lib.Check(err)

	fmt.Printf("Started to subscribe in chain %v at: %v\n", eventBundle.chainName, time.Now())
	for {
		select {
		case <-eventBundle.doneCh:
			fmt.Printf("Read done in chain %v\n", eventBundle.chainName)
			return
		case err := <-sub.Err():
			fmt.Printf("Error in chain %v at: %v, err: %v\n", eventBundle.chainName, time.Now(), err)
			eventBundle.interrupted <- eventBundle
			return
		case vLog := <-logs:
			fmt.Printf("Log at chain %v at: %v\n", eventBundle.chainName, time.Now())
			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog)
			ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)
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
			fmt.Println("full event: ", entity)

			fmt.Println("--------------------------------------------------------")
			// write in influx
		}
	}
}
