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

func SubscriptionManager() {
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

	// polygon alchemy
	wgGroup.Add(1)
	go Subscribe(
		os.Getenv("PLG_MN_ALCH_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
	)

	// arbitrum alchemy
	wgGroup.Add(1)
	go Subscribe(
		os.Getenv("ARB_MN_ALC_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_ARB_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
	)

	// eth alchemy
	wgGroup.Add(1)
	go Subscribe(
		os.Getenv("ETH_MN_ALC_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
	)

	// optimism alchemy
	wgGroup.Add(1)
	go Subscribe(
		os.Getenv("OPT_MN_ALC_WS_URL"),
		common.HexToAddress(os.Getenv("AAVE3_OPT_MN_POOL_ADDRESS")),
		done,
		&wgGroup,
	)
	wgGroup.Wait()
}

func cleanup(chanel <-chan os.Signal, done chan<- bool, wgGroup *sync.WaitGroup) {
	<-chanel
	done <- true
	fmt.Println("cleanup")
	wgGroup.Wait()
	fmt.Println("all goroutines finished, exit")
	os.Exit(1)
}
