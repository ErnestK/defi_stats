package liquidation_event

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func SubscriptionManagerNew() {
	defer func() {
		fmt.Printf("Terminated program at: %v\n", time.Now())
	}()
	var wgGroup sync.WaitGroup

	err := godotenv.Load()
	lib.Check(err)

	done := make(chan bool, 1)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go cleanup(signalChan, done, &wgGroup)

	wgGroup.Add(1)
	go SubscribeNew(
		"Polygon mainnet",
		os.Getenv("PLG_MN_ALCH_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
		common.HexToAddress(os.Getenv("AAVE3_PLG_MN_ORACLE_ADDRESS")),
	)

	wgGroup.Add(1)
	go SubscribeNew(
		"Arbitrum mainnet",
		os.Getenv("ARB_MN_ALC_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_ARB_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
		common.HexToAddress(os.Getenv("AAVE3_ARB_MN_ORACLE_ADDRESS")),
	)

	wgGroup.Add(1)
	go SubscribeNew(
		"Eth mainnet",
		os.Getenv("ETH_MN_ALC_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_ETH_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
		common.HexToAddress(os.Getenv("AAVE3_ETH_MN_ORACLE_ADDRESS")),
	)

	wgGroup.Add(1)
	go SubscribeNew(
		"Optimism mainnet",
		os.Getenv("OPT_MN_ALC_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_OPT_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
		common.HexToAddress(os.Getenv("AAVE3_OPT_MN_ORACLE_ADDRESS")),
	)
	wgGroup.Wait()
}
