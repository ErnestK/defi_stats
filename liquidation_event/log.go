package liquidation_event

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func Log() {
	err := godotenv.Load()
	lib.Check(err)
	endpoint := os.Getenv("PLG_MN_ALCH_HTTPS_URL")

	// block with the event is the closest for the new year 12.2022
	polygonBlockWithEvent := int64(37446903)

	client, err := ethclient.Dial(endpoint)
	lib.Check(err)

	latestBlockTemp, err := client.BlockNumber(context.Background())
	lib.Check(err)
	latestBlock := int64(latestBlockTemp)

	fmt.Println("from block:", polygonBlockWithEvent)
	fmt.Println("to block:", latestBlock)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(polygonBlockWithEvent),
		ToBlock:   big.NewInt(polygonBlockWithEvent),
		Addresses: []common.Address{
			common.HexToAddress(os.Getenv("AAVE_PLG_MN_POOL_ADDRESS")),
		},
		Topics: [][]common.Hash{{AaveLiquidationEventSig()}},
	}
	eventLogs, err := client.FilterLogs(context.Background(), query)
	lib.Check(err)

	fmt.Println("All Logs length: ", len(eventLogs))

	for _, vLog := range eventLogs {
		liquidationCallEventEntity := LiquidationCallEventEntity{}
		DeserializeEventLog(&liquidationCallEventEntity, vLog.Data)
		ShowLiquidationEventInfo(&liquidationCallEventEntity, vLog)
		lib.ShowBlockDate(polygonBlockWithEvent, client)
	}
}
