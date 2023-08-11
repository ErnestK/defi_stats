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

const NUMBER_OF_WS_LISTEN = 4

func SubscriptionManager() {
	defer func() {
		fmt.Printf("Terminated program at: %v\n", time.Now())
	}()
	var wgGroup sync.WaitGroup

	err := godotenv.Load()
	lib.Check(err)

	done := make(chan bool)
	interrupted := make(chan EventBundle, NUMBER_OF_WS_LISTEN)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go cleanup(signalChan, done, &wgGroup)

	wgGroup.Add(1)
	go Subscribe(
		EventBundle{
			chainName:     "Polygon mainnet",
			connectUrl:    os.Getenv("PLG_MN_ALCH_WS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_PLG_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	wgGroup.Add(1)
	go Subscribe(
		EventBundle{
			chainName:     "Arbitrum mainnet",
			connectUrl:    os.Getenv("ARB_MN_ALC_WS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_ARB_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_ARB_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	wgGroup.Add(1)
	go Subscribe(
		EventBundle{
			chainName:     "Eth mainnet",
			connectUrl:    os.Getenv("ETH_MN_ALC_WS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_ETH_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_ETH_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	wgGroup.Add(1)
	go Subscribe(
		EventBundle{
			chainName:     "Optimism mainnet",
			connectUrl:    os.Getenv("OPT_MN_ALC_WS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_OPT_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_OPT_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	for interSubscription := range interrupted {
		time.Sleep(time.Second)
		interSubscription.waitGroup.Add(1)
		fmt.Printf("Subscription for chain %v is down, restaring %v\n", interSubscription.chainName, time.Now())
		fmt.Printf("----------------------------------------------------\n")
		go Subscribe(interSubscription)
	}
}
