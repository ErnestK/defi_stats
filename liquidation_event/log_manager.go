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

func LogManager(fromTime time.Time) {
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

	interval := fromTime.Add(-1 * time.Hour)
	Log(
		interval,
		EventBundle{
			chainName:     "Polygon mainnet",
			connectUrl:    os.Getenv("PLG_MN_ALCH_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_PLG_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_PLG_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	Log(
		interval,
		EventBundle{
			chainName:     "ETH mainnet",
			connectUrl:    os.Getenv("ETH_MN_ALC_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_ETH_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_ETH_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	Log(
		interval,
		EventBundle{
			chainName:     "Arbutrum mainnet",
			connectUrl:    os.Getenv("ARB_MN_ALC_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_ARB_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_ARB_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)

	Log(
		interval,
		EventBundle{
			chainName:     "Optimism mainnet",
			connectUrl:    os.Getenv("OPT_MN_ALC_HTTPS_URL"),
			poolAddress:   common.HexToAddress(os.Getenv("AAVE3_OPT_MN_POOL_ADDRESS")),
			oracleAddress: common.HexToAddress(os.Getenv("AAVE3_OPT_MN_ORACLE_ADDRESS")),
			interrupted:   interrupted,
			doneCh:        done,
			waitGroup:     &wgGroup,
		},
	)
}
