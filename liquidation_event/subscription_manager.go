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

type SubscriptionPackageToStart struct {
	ChainName     string
	WsUrl         string
	PoolAddress   common.Address
	OracleAddress common.Address
	Interrupted   chan<- SubscriptionPackageToStart
	DoneCh        <-chan bool
	WaitGroup     *sync.WaitGroup
}

func SubscriptionManager() {
	defer func() {
		fmt.Printf("Terminated program at: %v\n", time.Now())
	}()
	var wgGroup sync.WaitGroup

	err := godotenv.Load()
	lib.Check(err)

	done := make(chan bool)
	interrupted := make(chan SubscriptionPackageToStart, NUMBER_OF_WS_LISTEN)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go cleanup(signalChan, done, &wgGroup)

	wgGroup.Add(1)
	go Subscribe(
		SubscriptionPackageToStart{
			ChainName:     "Polygon mainnet",
			WsUrl:         os.Getenv("PLG_MN_ALCH_WS_URL"),
			PoolAddress:   common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
			OracleAddress: common.HexToAddress(os.Getenv("AAVE3_PLG_MN_ORACLE_ADDRESS")),
			Interrupted:   interrupted,
			DoneCh:        done,
			WaitGroup:     &wgGroup,
		},
	)

	wgGroup.Add(1)
	go Subscribe(
		SubscriptionPackageToStart{
			ChainName:     "Arbitrum mainnet",
			WsUrl:         os.Getenv("ARB_MN_ALC_WS_URL"),
			PoolAddress:   common.HexToAddress(os.Getenv("AAVE3_ARB_MN_POOL_ADDRESS")),
			OracleAddress: common.HexToAddress(os.Getenv("AAVE3_ARB_MN_ORACLE_ADDRESS")),
			Interrupted:   interrupted,
			DoneCh:        done,
			WaitGroup:     &wgGroup,
		},
	)

	wgGroup.Add(1)
	go Subscribe(
		SubscriptionPackageToStart{
			ChainName:     "Eth mainnet",
			WsUrl:         os.Getenv("ETH_MN_ALC_WS_URL"),
			PoolAddress:   common.HexToAddress(os.Getenv("AAVE3_ETH_MN_POOL_ADDRESS")),
			OracleAddress: common.HexToAddress(os.Getenv("AAVE3_ETH_MN_ORACLE_ADDRESS")),
			Interrupted:   interrupted,
			DoneCh:        done,
			WaitGroup:     &wgGroup,
		},
	)

	wgGroup.Add(1)
	go Subscribe(
		SubscriptionPackageToStart{
			ChainName:     "Optimism mainnet",
			WsUrl:         os.Getenv("OPT_MN_ALC_WS_URL"),
			PoolAddress:   common.HexToAddress(os.Getenv("AAVE3_OPT_MN_POOL_ADDRESS")),
			OracleAddress: common.HexToAddress(os.Getenv("AAVE3_OPT_MN_ORACLE_ADDRESS")),
			Interrupted:   interrupted,
			DoneCh:        done,
			WaitGroup:     &wgGroup,
		},
	)

	for interSubscription := range interrupted {
		time.Sleep(time.Second)
		interSubscription.WaitGroup.Add(1)
		fmt.Printf("Subscription for chain %v is down, restaring %v\n", interSubscription.ChainName, time.Now())
		go Subscribe(interSubscription)
	}
}

func cleanup(chanel <-chan os.Signal, done chan<- bool, wgGroup *sync.WaitGroup) {
	<-chanel
	close(done)
	fmt.Println("cleanup")
	wgGroup.Wait()
	fmt.Println("all goroutines finished, exit")
	os.Exit(1)
}
