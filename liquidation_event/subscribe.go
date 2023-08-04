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

func Subscribe(subsPackage SubscriptionPackageToStart) {
	defer subsPackage.WaitGroup.Done()

	client, err := ethclient.Dial(subsPackage.WsUrl)
	lib.Check(err)
	defer client.Close()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			subsPackage.PoolAddress,
		},
		Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	lib.Check(err)

	fmt.Printf("Started to subscribe in chain %v at: %v\n", subsPackage.ChainName, time.Now())
	for {
		select {
		case <-subsPackage.DoneCh:
			fmt.Printf("Read done in chain %v\n", subsPackage.ChainName)
			return
		case err := <-sub.Err():
			fmt.Printf("Error in chain %v at: %v\n", subsPackage.ChainName, time.Now())
			fmt.Printf("Error: %v\n", err)
			return
		case vLog := <-logs:
			fmt.Printf("Log at chain %v at: %v\n", subsPackage.ChainName, time.Now())
			liquidationCallEventEntity := LiquidationCallEventEntity{}
			DeserializeEventLog(&liquidationCallEventEntity, vLog.Data)

			liquidationCallEventEntity.CollateralAsset = common.HexToAddress(vLog.Topics[1].Hex())
			liquidationCallEventEntity.DebtAsset = common.HexToAddress(vLog.Topics[2].Hex())
			liquidationCallEventEntity.User = common.HexToAddress(vLog.Topics[3].Hex())
			ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)
			caMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.CollateralAsset, client)
			daMetadata := lib.CoinMetadataForAddress(liquidationCallEventEntity.DebtAsset, client)

			caPrice := GetAssetPrice(liquidationCallEventEntity.CollateralAsset, caMetadata.Decimal, client, subsPackage.OracleAddress)
			daPrice := GetAssetPrice(liquidationCallEventEntity.DebtAsset, daMetadata.Decimal, client, subsPackage.OracleAddress)

			fmt.Println("ca name: ", caMetadata.Name)
			fmt.Println("caPrice: ", caPrice)
			fmt.Println("da name: ", daMetadata.Name)
			fmt.Println("daPrice: ", daPrice)

			entity := LiquidationCallEventEntityExpanded{
				liquidationCallEventEntity: liquidationCallEventEntity,
				caInWei:                    caPrice,
				daInWei:                    daPrice,
				chainName:                  subsPackage.ChainName,
				timestamp:                  time.Now(),
			}
			fmt.Println("full event: ", entity)

			fmt.Println("--------------------------------------------------------")
			// write in influx
		}
	}
}
