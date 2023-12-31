package lib

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ernest_k/defi_stats/lib/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenMetadataEntity struct {
	Name    string
	Decimal uint8
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckWithMessage(e error, message string) {
	if e != nil {
		log.Println(message)
		log.Fatal(e)
	}
}

func ShowBlockData(blockNum int64, client *ethclient.Client) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("Transaction Hash:", tx.Hash().Hex())
		// fmt.Println("From:", tx.From().Hex())
		fmt.Println("To:", tx.To().Hex())
		fmt.Println("Value:", tx.Value().String())
		fmt.Println("Gas Limit:", tx.Gas())
		fmt.Println("Gas Price:", tx.GasPrice().String())
		fmt.Println("Nonce:", tx.Nonce())
		fmt.Println("Data:", tx.Data())
	}
}

func ShowBlockDate(blockNum *big.Int, client *ethclient.Client) time.Time {
	header, err := client.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		log.Fatal(err)
	}

	return time.Unix(int64(header.Time), 0)
}

func ContractABIFor(path string) abi.ABI {
	absPath, err := filepath.Abs(path)
	Check(err)
	data, err := os.ReadFile(absPath)
	Check(err)
	contractAbi, err := abi.JSON(strings.NewReader(string(data)))
	Check(err)

	return contractAbi
}

func CoinMetadataForAddress(address common.Address, client *ethclient.Client) TokenMetadataEntity {
	var name string
	instance, err := erc20.NewErc20(address, client)
	Check(err)

	opts := bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     context.Background(),
	}

	name, err = instance.Name(&opts)
	if err != nil {
		// TODO: some coin stores name not in string, need to add additional contract
		// https://ethereum.stackexchange.com/questions/130586/ethers-js-error-when-calling-symbol-on-maker-mkr-contract-but-no-problem-fo
		name = "na"
	}

	decimal, err := instance.Decimals(&opts)
	Check(err)

	return TokenMetadataEntity{name, decimal}
}
