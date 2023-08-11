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

			caPrice := GetAssetPrice(liquidationCallEventEntity.CollateralAsset, caMetadata.Decimal, client, eventBundle.oracleAddress)
			daPrice := GetAssetPrice(liquidationCallEventEntity.DebtAsset, daMetadata.Decimal, client, eventBundle.oracleAddress)

			fmt.Println("ca name: ", caMetadata.Name)
			fmt.Println("caPrice: ", caPrice)
			fmt.Println("da name: ", daMetadata.Name)
			fmt.Println("daPrice: ", daPrice)

			entity := LiquidationCallEventEntityExpanded{
				liquidationCallEventEntity: liquidationCallEventEntity,
				caInWei:                    caPrice,
				daInWei:                    daPrice,
				chainName:                  eventBundle.chainName,
				timestamp:                  time.Now(),
			}
			fmt.Println("full event: ", entity)

			fmt.Println("--------------------------------------------------------")
			// write in influx
		}
	}
}
