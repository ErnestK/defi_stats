package liquidation_event

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Subscribe(chainName string, wsUrl string, poolAddress common.Address, done <-chan bool, wgGroup *sync.WaitGroup, oracleAddress common.Address) error {
	defer wgGroup.Done()

	client, err := ethclient.Dial(wsUrl)
	lib.Check(err)
	defer client.Close()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			poolAddress,
		},
		Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	// https://github.com/ethereum/go-ethereum/issues/22266
	// logs := make(chan types.Log, 1_000)
	// sub := event.Resubscribe(1*time.Second, func(ctx context.Context) (event.Subscription, error) {
	// 	return client.SubscribeFilterLogs(ctx, query, logs)
	// })
	// defer sub.Unsubscribe()
	lib.Check(err)

	fmt.Printf("Started to subscribe in chain %v at: %v\n", chainName, time.Now())
	for {
		select {
		case <-done:
			fmt.Printf("Read done in chain %v\n", chainName)
			return nil
		case err := <-sub.Err():
			fmt.Printf("Error in chain %v at: %v\n", chainName, time.Now())
			fmt.Printf("Error: %v\n", err)
			return err
		case vLog := <-logs:
			fmt.Printf("Log at chain %v at: %v\n", chainName, time.Now())
			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog.Data)

			liquidationCallEventEntity.CollateralAsset = common.HexToAddress(vLog.Topics[1].Hex())
			liquidationCallEventEntity.DebtAsset = common.HexToAddress(vLog.Topics[2].Hex())
			liquidationCallEventEntity.User = common.HexToAddress(vLog.Topics[3].Hex())
			ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)
			caMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.CollateralAsset, client)
			daMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.DebtAsset, client)

			caPrice := GetAssetPrice(liquidationCallEventEntity.CollateralAsset, caMetadata.Decimal, client, oracleAddress)
			daPrice := GetAssetPrice(liquidationCallEventEntity.DebtAsset, daMetadata.Decimal, client, oracleAddress)

			fmt.Println("ca name: ", caMetadata.Name)
			fmt.Println("caPrice: ", caPrice)
			fmt.Println("da name: ", daMetadata.Name)
			fmt.Println("daPrice: ", daPrice)

			entity := LiquidationCallEventEntityExpanded{
				liquidationCallEventEntity: liquidationCallEventEntity,
				caInWei:                    caPrice,
				daInWei:                    daPrice,
				chainName:                  chainName,
				timestamp:                  time.Now(),
			}
			fmt.Println("full event: ", entity)

			fmt.Println("--------------------------------------------------------")
			// write in influx

			return nil
		}
	}
}
