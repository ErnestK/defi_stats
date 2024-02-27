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

func LogManager() {
	defer func() {
		fmt.Printf("Program finish succesfully at: %v\n", time.Now())
	}()
	var wgGroup sync.WaitGroup

	err := godotenv.Load()
	lib.Check(err)

	done := make(chan bool)
	interrupted := make(chan EventBundle, NUMBER_OF_WS_LISTEN)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go cleanup(signalChan, done, &wgGroup)

	Log(
		EventBundle{
			chainName:     "Polygon_mainnet",
			connectUrl:    os.Getenv("PLG_MN_ALCH_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_PLG_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	Log(
		EventBundle{
			chainName:     "ETH_mainnet",
			connectUrl:    os.Getenv("ETH_MN_ALC_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_ETH_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_ETH_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	Log(
		EventBundle{
			chainName:     "Arbutrum_mainnet",
			connectUrl:    os.Getenv("ARB_MN_ALC_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_ARB_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_ARB_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	Log(
		EventBundle{
			chainName:     "Optimism_mainnet",
			connectUrl:    os.Getenv("OPT_MN_ALC_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_OPT_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_OPT_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)
}
