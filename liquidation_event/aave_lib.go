package liquidation_event

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ernest_k/defi_stats/lib/aave"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LiquidationCallEventEntity struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
}

type EventBundle struct {
	chainName     string
	connectUrl    string
	poolAddress   common.Address
	oracleAddress common.Address
	interrupted   chan<- EventBundle
	doneCh        <-chan bool
	waitGroup     *sync.WaitGroup
}

type LiquidationCallEventEntityExpanded struct {
	liquidationCallEventEntity LiquidationCallEventEntity
	caInWei                    float64
	daInWei                    float64
	chainName                  string
	timestamp                  time.Time
}

func DeserializeEventLog(liquidationCallEventEntity *LiquidationCallEventEntity, vLog types.Log) {
	liqudationEventName := "LiquidationCall"
	contractAbi := lib.ContractABIFor("abi/aave_pool.json")

	err := contractAbi.UnpackIntoInterface(liquidationCallEventEntity, liqudationEventName, vLog.Data)
	lib.Check(err)

	liquidationCallEventEntity.CollateralAsset = common.HexToAddress(vLog.Topics[1].Hex())
	liquidationCallEventEntity.DebtAsset = common.HexToAddress(vLog.Topics[2].Hex())
	liquidationCallEventEntity.User = common.HexToAddress(vLog.Topics[3].Hex())
}

func ShowLiquidationEventInfo(liquidationCallEventEntity *LiquidationCallEventEntity, vLog types.Log) {
	fmt.Println("liquidationCallEventEntity.CollateralAsset", liquidationCallEventEntity.CollateralAsset)
	fmt.Println("liquidationCallEventEntity.DebtAsset", liquidationCallEventEntity.DebtAsset)
	fmt.Println("liquidationCallEventEntity.User", liquidationCallEventEntity.User)
	fmt.Println("liquidationCallEventEntity.DebtToCover", liquidationCallEventEntity.DebtToCover)
	fmt.Println("liquidationCallEventEntity.LiquidatedCollateralAmount", liquidationCallEventEntity.LiquidatedCollateralAmount)
	fmt.Println("liquidationCallEventEntity.Liquidator", liquidationCallEventEntity.Liquidator)
	fmt.Println("liquidationCallEventEntity.ReceiveAToken", liquidationCallEventEntity.ReceiveAToken)
	fmt.Println("------------------------------------------------------")
}

func GetAssetPrice(tokenAddress common.Address, decimal uint8, client *ethclient.Client, oracleAddress common.Address) float64 {
	opts := bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     context.Background(),
	}

	instance, err := aave.NewAave(oracleAddress, client)
	lib.Check(err)

	tokenValue, err := instance.GetAssetPrice(&opts, tokenAddress)
	lib.Check(err)

	stringValue := tokenValue.String()
	floatTokenValue, err := strconv.ParseFloat(stringValue, 64)
	lib.Check(err)

	result := floatTokenValue / math.Pow(10, float64(decimal))

	return result
}

func AaveLiquidationEventSig() common.Hash {
	liquidationEventSig := []byte("LiquidationCall(address,address,address,uint256,uint256,address,bool)")
	return crypto.Keccak256Hash(liquidationEventSig)
}

func cleanup(chanel <-chan os.Signal, done chan<- bool, wgGroup *sync.WaitGroup) {
	<-chanel
	close(done)
	fmt.Println("cleanup")
	wgGroup.Wait()
	fmt.Println("all goroutines finished, exit")
	os.Exit(1)
}
