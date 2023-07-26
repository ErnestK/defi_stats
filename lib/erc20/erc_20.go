// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"burnMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"delegateAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFeeFlat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_canReceiveMintWhiteList\",\"type\":\"address\"},{\"name\":\"_canBurnWhiteList\",\"type\":\"address\"},{\"name\":\"_blackList\",\"type\":\"address\"},{\"name\":\"_noFeesList\",\"type\":\"address\"}],\"name\":\"setLists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"delegateToNewContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_transferFeeNumerator\",\"type\":\"uint80\"},{\"name\":\"_transferFeeDenominator\",\"type\":\"uint80\"},{\"name\":\"_mintFeeNumerator\",\"type\":\"uint80\"},{\"name\":\"_mintFeeDenominator\",\"type\":\"uint80\"},{\"name\":\"_mintFeeFlat\",\"type\":\"uint256\"},{\"name\":\"_burnFeeNumerator\",\"type\":\"uint80\"},{\"name\":\"_burnFeeDenominator\",\"type\":\"uint80\"},{\"name\":\"_burnFeeFlat\",\"type\":\"uint256\"}],\"name\":\"changeStakingFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canReceiveMintWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegatedFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateApprove\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"reclaimContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"delegateBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateIncreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFeeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canBurnWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnMax\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintFeeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDelegatedFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"noFeesList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMin\",\"type\":\"uint256\"},{\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"changeBurnBounds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegateTotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"changeName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintFeeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferFeeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateDecreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newStaker\",\"type\":\"address\"}],\"name\":\"changeStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"wipeBlacklistedAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFeeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferFeeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintFeeFlat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setAllowanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newMin\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"ChangeBurnBoundsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"WipedAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"DelegatedTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowance", _owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Session) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, _owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, _owner, spender)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() view returns(address)
func (_Erc20 *Erc20Caller) Allowances(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowances")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() view returns(address)
func (_Erc20 *Erc20Session) Allowances() (common.Address, error) {
	return _Erc20.Contract.Allowances(&_Erc20.CallOpts)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() view returns(address)
func (_Erc20 *Erc20CallerSession) Allowances() (common.Address, error) {
	return _Erc20.Contract.Allowances(&_Erc20.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(address)
func (_Erc20 *Erc20Caller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balances")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(address)
func (_Erc20 *Erc20Session) Balances() (common.Address, error) {
	return _Erc20.Contract.Balances(&_Erc20.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(address)
func (_Erc20 *Erc20CallerSession) Balances() (common.Address, error) {
	return _Erc20.Contract.Balances(&_Erc20.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0xcdab73b5.
//
// Solidity: function blackList() view returns(address)
func (_Erc20 *Erc20Caller) BlackList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "blackList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlackList is a free data retrieval call binding the contract method 0xcdab73b5.
//
// Solidity: function blackList() view returns(address)
func (_Erc20 *Erc20Session) BlackList() (common.Address, error) {
	return _Erc20.Contract.BlackList(&_Erc20.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0xcdab73b5.
//
// Solidity: function blackList() view returns(address)
func (_Erc20 *Erc20CallerSession) BlackList() (common.Address, error) {
	return _Erc20.Contract.BlackList(&_Erc20.CallOpts)
}

// BurnFeeDenominator is a free data retrieval call binding the contract method 0xc18f4831.
//
// Solidity: function burnFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20Caller) BurnFeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "burnFeeDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnFeeDenominator is a free data retrieval call binding the contract method 0xc18f4831.
//
// Solidity: function burnFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20Session) BurnFeeDenominator() (*big.Int, error) {
	return _Erc20.Contract.BurnFeeDenominator(&_Erc20.CallOpts)
}

// BurnFeeDenominator is a free data retrieval call binding the contract method 0xc18f4831.
//
// Solidity: function burnFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20CallerSession) BurnFeeDenominator() (*big.Int, error) {
	return _Erc20.Contract.BurnFeeDenominator(&_Erc20.CallOpts)
}

// BurnFeeFlat is a free data retrieval call binding the contract method 0x0b8e845a.
//
// Solidity: function burnFeeFlat() view returns(uint256)
func (_Erc20 *Erc20Caller) BurnFeeFlat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "burnFeeFlat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnFeeFlat is a free data retrieval call binding the contract method 0x0b8e845a.
//
// Solidity: function burnFeeFlat() view returns(uint256)
func (_Erc20 *Erc20Session) BurnFeeFlat() (*big.Int, error) {
	return _Erc20.Contract.BurnFeeFlat(&_Erc20.CallOpts)
}

// BurnFeeFlat is a free data retrieval call binding the contract method 0x0b8e845a.
//
// Solidity: function burnFeeFlat() view returns(uint256)
func (_Erc20 *Erc20CallerSession) BurnFeeFlat() (*big.Int, error) {
	return _Erc20.Contract.BurnFeeFlat(&_Erc20.CallOpts)
}

// BurnFeeNumerator is a free data retrieval call binding the contract method 0x56e1c40d.
//
// Solidity: function burnFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20Caller) BurnFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "burnFeeNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnFeeNumerator is a free data retrieval call binding the contract method 0x56e1c40d.
//
// Solidity: function burnFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20Session) BurnFeeNumerator() (*big.Int, error) {
	return _Erc20.Contract.BurnFeeNumerator(&_Erc20.CallOpts)
}

// BurnFeeNumerator is a free data retrieval call binding the contract method 0x56e1c40d.
//
// Solidity: function burnFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20CallerSession) BurnFeeNumerator() (*big.Int, error) {
	return _Erc20.Contract.BurnFeeNumerator(&_Erc20.CallOpts)
}

// BurnMax is a free data retrieval call binding the contract method 0x5c131d70.
//
// Solidity: function burnMax() view returns(uint256)
func (_Erc20 *Erc20Caller) BurnMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "burnMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnMax is a free data retrieval call binding the contract method 0x5c131d70.
//
// Solidity: function burnMax() view returns(uint256)
func (_Erc20 *Erc20Session) BurnMax() (*big.Int, error) {
	return _Erc20.Contract.BurnMax(&_Erc20.CallOpts)
}

// BurnMax is a free data retrieval call binding the contract method 0x5c131d70.
//
// Solidity: function burnMax() view returns(uint256)
func (_Erc20 *Erc20CallerSession) BurnMax() (*big.Int, error) {
	return _Erc20.Contract.BurnMax(&_Erc20.CallOpts)
}

// BurnMin is a free data retrieval call binding the contract method 0x02d3fdc9.
//
// Solidity: function burnMin() view returns(uint256)
func (_Erc20 *Erc20Caller) BurnMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "burnMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnMin is a free data retrieval call binding the contract method 0x02d3fdc9.
//
// Solidity: function burnMin() view returns(uint256)
func (_Erc20 *Erc20Session) BurnMin() (*big.Int, error) {
	return _Erc20.Contract.BurnMin(&_Erc20.CallOpts)
}

// BurnMin is a free data retrieval call binding the contract method 0x02d3fdc9.
//
// Solidity: function burnMin() view returns(uint256)
func (_Erc20 *Erc20CallerSession) BurnMin() (*big.Int, error) {
	return _Erc20.Contract.BurnMin(&_Erc20.CallOpts)
}

// CanBurnWhiteList is a free data retrieval call binding the contract method 0x5a444139.
//
// Solidity: function canBurnWhiteList() view returns(address)
func (_Erc20 *Erc20Caller) CanBurnWhiteList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "canBurnWhiteList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanBurnWhiteList is a free data retrieval call binding the contract method 0x5a444139.
//
// Solidity: function canBurnWhiteList() view returns(address)
func (_Erc20 *Erc20Session) CanBurnWhiteList() (common.Address, error) {
	return _Erc20.Contract.CanBurnWhiteList(&_Erc20.CallOpts)
}

// CanBurnWhiteList is a free data retrieval call binding the contract method 0x5a444139.
//
// Solidity: function canBurnWhiteList() view returns(address)
func (_Erc20 *Erc20CallerSession) CanBurnWhiteList() (common.Address, error) {
	return _Erc20.Contract.CanBurnWhiteList(&_Erc20.CallOpts)
}

// CanReceiveMintWhiteList is a free data retrieval call binding the contract method 0x1f7af1df.
//
// Solidity: function canReceiveMintWhiteList() view returns(address)
func (_Erc20 *Erc20Caller) CanReceiveMintWhiteList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "canReceiveMintWhiteList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanReceiveMintWhiteList is a free data retrieval call binding the contract method 0x1f7af1df.
//
// Solidity: function canReceiveMintWhiteList() view returns(address)
func (_Erc20 *Erc20Session) CanReceiveMintWhiteList() (common.Address, error) {
	return _Erc20.Contract.CanReceiveMintWhiteList(&_Erc20.CallOpts)
}

// CanReceiveMintWhiteList is a free data retrieval call binding the contract method 0x1f7af1df.
//
// Solidity: function canReceiveMintWhiteList() view returns(address)
func (_Erc20 *Erc20CallerSession) CanReceiveMintWhiteList() (common.Address, error) {
	return _Erc20.Contract.CanReceiveMintWhiteList(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() view returns(address)
func (_Erc20 *Erc20Caller) Delegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "delegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() view returns(address)
func (_Erc20 *Erc20Session) Delegate() (common.Address, error) {
	return _Erc20.Contract.Delegate(&_Erc20.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() view returns(address)
func (_Erc20 *Erc20CallerSession) Delegate() (common.Address, error) {
	return _Erc20.Contract.Delegate(&_Erc20.CallOpts)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Caller) DelegateAllowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "delegateAllowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Session) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.DelegateAllowance(&_Erc20.CallOpts, owner, spender)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20CallerSession) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.DelegateAllowance(&_Erc20.CallOpts, owner, spender)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20Caller) DelegateBalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "delegateBalanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20Session) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20.Contract.DelegateBalanceOf(&_Erc20.CallOpts, who)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20CallerSession) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20.Contract.DelegateBalanceOf(&_Erc20.CallOpts, who)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) DelegateTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "delegateTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) DelegateTotalSupply() (*big.Int, error) {
	return _Erc20.Contract.DelegateTotalSupply(&_Erc20.CallOpts)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) DelegateTotalSupply() (*big.Int, error) {
	return _Erc20.Contract.DelegateTotalSupply(&_Erc20.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() view returns(address)
func (_Erc20 *Erc20Caller) DelegatedFrom(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "delegatedFrom")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() view returns(address)
func (_Erc20 *Erc20Session) DelegatedFrom() (common.Address, error) {
	return _Erc20.Contract.DelegatedFrom(&_Erc20.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() view returns(address)
func (_Erc20 *Erc20CallerSession) DelegatedFrom() (common.Address, error) {
	return _Erc20.Contract.DelegatedFrom(&_Erc20.CallOpts)
}

// MintFeeDenominator is a free data retrieval call binding the contract method 0x5db07aee.
//
// Solidity: function mintFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20Caller) MintFeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "mintFeeDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeDenominator is a free data retrieval call binding the contract method 0x5db07aee.
//
// Solidity: function mintFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20Session) MintFeeDenominator() (*big.Int, error) {
	return _Erc20.Contract.MintFeeDenominator(&_Erc20.CallOpts)
}

// MintFeeDenominator is a free data retrieval call binding the contract method 0x5db07aee.
//
// Solidity: function mintFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20CallerSession) MintFeeDenominator() (*big.Int, error) {
	return _Erc20.Contract.MintFeeDenominator(&_Erc20.CallOpts)
}

// MintFeeFlat is a free data retrieval call binding the contract method 0xd63a1389.
//
// Solidity: function mintFeeFlat() view returns(uint256)
func (_Erc20 *Erc20Caller) MintFeeFlat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "mintFeeFlat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeFlat is a free data retrieval call binding the contract method 0xd63a1389.
//
// Solidity: function mintFeeFlat() view returns(uint256)
func (_Erc20 *Erc20Session) MintFeeFlat() (*big.Int, error) {
	return _Erc20.Contract.MintFeeFlat(&_Erc20.CallOpts)
}

// MintFeeFlat is a free data retrieval call binding the contract method 0xd63a1389.
//
// Solidity: function mintFeeFlat() view returns(uint256)
func (_Erc20 *Erc20CallerSession) MintFeeFlat() (*big.Int, error) {
	return _Erc20.Contract.MintFeeFlat(&_Erc20.CallOpts)
}

// MintFeeNumerator is a free data retrieval call binding the contract method 0x8d93eac2.
//
// Solidity: function mintFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20Caller) MintFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "mintFeeNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeNumerator is a free data retrieval call binding the contract method 0x8d93eac2.
//
// Solidity: function mintFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20Session) MintFeeNumerator() (*big.Int, error) {
	return _Erc20.Contract.MintFeeNumerator(&_Erc20.CallOpts)
}

// MintFeeNumerator is a free data retrieval call binding the contract method 0x8d93eac2.
//
// Solidity: function mintFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20CallerSession) MintFeeNumerator() (*big.Int, error) {
	return _Erc20.Contract.MintFeeNumerator(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// NoFeesList is a free data retrieval call binding the contract method 0x6d4717fe.
//
// Solidity: function noFeesList() view returns(address)
func (_Erc20 *Erc20Caller) NoFeesList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "noFeesList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NoFeesList is a free data retrieval call binding the contract method 0x6d4717fe.
//
// Solidity: function noFeesList() view returns(address)
func (_Erc20 *Erc20Session) NoFeesList() (common.Address, error) {
	return _Erc20.Contract.NoFeesList(&_Erc20.CallOpts)
}

// NoFeesList is a free data retrieval call binding the contract method 0x6d4717fe.
//
// Solidity: function noFeesList() view returns(address)
func (_Erc20 *Erc20CallerSession) NoFeesList() (common.Address, error) {
	return _Erc20.Contract.NoFeesList(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Session) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20CallerSession) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc20 *Erc20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc20 *Erc20Session) Paused() (bool, error) {
	return _Erc20.Contract.Paused(&_Erc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc20 *Erc20CallerSession) Paused() (bool, error) {
	return _Erc20.Contract.Paused(&_Erc20.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Erc20 *Erc20Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Erc20 *Erc20Session) PendingOwner() (common.Address, error) {
	return _Erc20.Contract.PendingOwner(&_Erc20.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Erc20 *Erc20CallerSession) PendingOwner() (common.Address, error) {
	return _Erc20.Contract.PendingOwner(&_Erc20.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_Erc20 *Erc20Caller) Staker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "staker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_Erc20 *Erc20Session) Staker() (common.Address, error) {
	return _Erc20.Contract.Staker(&_Erc20.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_Erc20 *Erc20CallerSession) Staker() (common.Address, error) {
	return _Erc20.Contract.Staker(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TransferFeeDenominator is a free data retrieval call binding the contract method 0xd42cfc41.
//
// Solidity: function transferFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20Caller) TransferFeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "transferFeeDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferFeeDenominator is a free data retrieval call binding the contract method 0xd42cfc41.
//
// Solidity: function transferFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20Session) TransferFeeDenominator() (*big.Int, error) {
	return _Erc20.Contract.TransferFeeDenominator(&_Erc20.CallOpts)
}

// TransferFeeDenominator is a free data retrieval call binding the contract method 0xd42cfc41.
//
// Solidity: function transferFeeDenominator() view returns(uint80)
func (_Erc20 *Erc20CallerSession) TransferFeeDenominator() (*big.Int, error) {
	return _Erc20.Contract.TransferFeeDenominator(&_Erc20.CallOpts)
}

// TransferFeeNumerator is a free data retrieval call binding the contract method 0x8f98ce8f.
//
// Solidity: function transferFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20Caller) TransferFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "transferFeeNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferFeeNumerator is a free data retrieval call binding the contract method 0x8f98ce8f.
//
// Solidity: function transferFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20Session) TransferFeeNumerator() (*big.Int, error) {
	return _Erc20.Contract.TransferFeeNumerator(&_Erc20.CallOpts)
}

// TransferFeeNumerator is a free data retrieval call binding the contract method 0x8f98ce8f.
//
// Solidity: function transferFeeNumerator() view returns(uint80)
func (_Erc20 *Erc20CallerSession) TransferFeeNumerator() (*big.Int, error) {
	return _Erc20.Contract.TransferFeeNumerator(&_Erc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Erc20 *Erc20Transactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Erc20 *Erc20Session) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Erc20 *Erc20TransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, _value)
}

// ChangeBurnBounds is a paid mutator transaction binding the contract method 0x70df42e1.
//
// Solidity: function changeBurnBounds(uint256 newMin, uint256 newMax) returns()
func (_Erc20 *Erc20Transactor) ChangeBurnBounds(opts *bind.TransactOpts, newMin *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "changeBurnBounds", newMin, newMax)
}

// ChangeBurnBounds is a paid mutator transaction binding the contract method 0x70df42e1.
//
// Solidity: function changeBurnBounds(uint256 newMin, uint256 newMax) returns()
func (_Erc20 *Erc20Session) ChangeBurnBounds(newMin *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeBurnBounds(&_Erc20.TransactOpts, newMin, newMax)
}

// ChangeBurnBounds is a paid mutator transaction binding the contract method 0x70df42e1.
//
// Solidity: function changeBurnBounds(uint256 newMin, uint256 newMax) returns()
func (_Erc20 *Erc20TransactorSession) ChangeBurnBounds(newMin *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeBurnBounds(&_Erc20.TransactOpts, newMin, newMax)
}

// ChangeName is a paid mutator transaction binding the contract method 0x86575e40.
//
// Solidity: function changeName(string _name, string _symbol) returns()
func (_Erc20 *Erc20Transactor) ChangeName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "changeName", _name, _symbol)
}

// ChangeName is a paid mutator transaction binding the contract method 0x86575e40.
//
// Solidity: function changeName(string _name, string _symbol) returns()
func (_Erc20 *Erc20Session) ChangeName(_name string, _symbol string) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeName(&_Erc20.TransactOpts, _name, _symbol)
}

// ChangeName is a paid mutator transaction binding the contract method 0x86575e40.
//
// Solidity: function changeName(string _name, string _symbol) returns()
func (_Erc20 *Erc20TransactorSession) ChangeName(_name string, _symbol string) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeName(&_Erc20.TransactOpts, _name, _symbol)
}

// ChangeStaker is a paid mutator transaction binding the contract method 0xab55979d.
//
// Solidity: function changeStaker(address newStaker) returns()
func (_Erc20 *Erc20Transactor) ChangeStaker(opts *bind.TransactOpts, newStaker common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "changeStaker", newStaker)
}

// ChangeStaker is a paid mutator transaction binding the contract method 0xab55979d.
//
// Solidity: function changeStaker(address newStaker) returns()
func (_Erc20 *Erc20Session) ChangeStaker(newStaker common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeStaker(&_Erc20.TransactOpts, newStaker)
}

// ChangeStaker is a paid mutator transaction binding the contract method 0xab55979d.
//
// Solidity: function changeStaker(address newStaker) returns()
func (_Erc20 *Erc20TransactorSession) ChangeStaker(newStaker common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeStaker(&_Erc20.TransactOpts, newStaker)
}

// ChangeStakingFees is a paid mutator transaction binding the contract method 0x1db8cb3f.
//
// Solidity: function changeStakingFees(uint80 _transferFeeNumerator, uint80 _transferFeeDenominator, uint80 _mintFeeNumerator, uint80 _mintFeeDenominator, uint256 _mintFeeFlat, uint80 _burnFeeNumerator, uint80 _burnFeeDenominator, uint256 _burnFeeFlat) returns()
func (_Erc20 *Erc20Transactor) ChangeStakingFees(opts *bind.TransactOpts, _transferFeeNumerator *big.Int, _transferFeeDenominator *big.Int, _mintFeeNumerator *big.Int, _mintFeeDenominator *big.Int, _mintFeeFlat *big.Int, _burnFeeNumerator *big.Int, _burnFeeDenominator *big.Int, _burnFeeFlat *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "changeStakingFees", _transferFeeNumerator, _transferFeeDenominator, _mintFeeNumerator, _mintFeeDenominator, _mintFeeFlat, _burnFeeNumerator, _burnFeeDenominator, _burnFeeFlat)
}

// ChangeStakingFees is a paid mutator transaction binding the contract method 0x1db8cb3f.
//
// Solidity: function changeStakingFees(uint80 _transferFeeNumerator, uint80 _transferFeeDenominator, uint80 _mintFeeNumerator, uint80 _mintFeeDenominator, uint256 _mintFeeFlat, uint80 _burnFeeNumerator, uint80 _burnFeeDenominator, uint256 _burnFeeFlat) returns()
func (_Erc20 *Erc20Session) ChangeStakingFees(_transferFeeNumerator *big.Int, _transferFeeDenominator *big.Int, _mintFeeNumerator *big.Int, _mintFeeDenominator *big.Int, _mintFeeFlat *big.Int, _burnFeeNumerator *big.Int, _burnFeeDenominator *big.Int, _burnFeeFlat *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeStakingFees(&_Erc20.TransactOpts, _transferFeeNumerator, _transferFeeDenominator, _mintFeeNumerator, _mintFeeDenominator, _mintFeeFlat, _burnFeeNumerator, _burnFeeDenominator, _burnFeeFlat)
}

// ChangeStakingFees is a paid mutator transaction binding the contract method 0x1db8cb3f.
//
// Solidity: function changeStakingFees(uint80 _transferFeeNumerator, uint80 _transferFeeDenominator, uint80 _mintFeeNumerator, uint80 _mintFeeDenominator, uint256 _mintFeeFlat, uint80 _burnFeeNumerator, uint80 _burnFeeDenominator, uint256 _burnFeeFlat) returns()
func (_Erc20 *Erc20TransactorSession) ChangeStakingFees(_transferFeeNumerator *big.Int, _transferFeeDenominator *big.Int, _mintFeeNumerator *big.Int, _mintFeeDenominator *big.Int, _mintFeeFlat *big.Int, _burnFeeNumerator *big.Int, _burnFeeDenominator *big.Int, _burnFeeFlat *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.ChangeStakingFees(&_Erc20.TransactOpts, _transferFeeNumerator, _transferFeeDenominator, _mintFeeNumerator, _mintFeeDenominator, _mintFeeFlat, _burnFeeNumerator, _burnFeeDenominator, _burnFeeFlat)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Erc20 *Erc20Transactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Erc20 *Erc20Session) ClaimOwnership() (*types.Transaction, error) {
	return _Erc20.Contract.ClaimOwnership(&_Erc20.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Erc20 *Erc20TransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Erc20.Contract.ClaimOwnership(&_Erc20.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20 *Erc20Transactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20 *Erc20Session) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.DecreaseApproval(&_Erc20.TransactOpts, spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20 *Erc20TransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.DecreaseApproval(&_Erc20.TransactOpts, spender, subtractedValue)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(address spender, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20Transactor) DelegateApprove(opts *bind.TransactOpts, spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "delegateApprove", spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(address spender, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20Session) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateApprove(&_Erc20.TransactOpts, spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(address spender, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20TransactorSession) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateApprove(&_Erc20.TransactOpts, spender, value, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(address spender, uint256 subtractedValue, address origSender) returns(bool)
func (_Erc20 *Erc20Transactor) DelegateDecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "delegateDecreaseApproval", spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(address spender, uint256 subtractedValue, address origSender) returns(bool)
func (_Erc20 *Erc20Session) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateDecreaseApproval(&_Erc20.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(address spender, uint256 subtractedValue, address origSender) returns(bool)
func (_Erc20 *Erc20TransactorSession) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateDecreaseApproval(&_Erc20.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(address spender, uint256 addedValue, address origSender) returns(bool)
func (_Erc20 *Erc20Transactor) DelegateIncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "delegateIncreaseApproval", spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(address spender, uint256 addedValue, address origSender) returns(bool)
func (_Erc20 *Erc20Session) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateIncreaseApproval(&_Erc20.TransactOpts, spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(address spender, uint256 addedValue, address origSender) returns(bool)
func (_Erc20 *Erc20TransactorSession) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateIncreaseApproval(&_Erc20.TransactOpts, spender, addedValue, origSender)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(address newContract) returns()
func (_Erc20 *Erc20Transactor) DelegateToNewContract(opts *bind.TransactOpts, newContract common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "delegateToNewContract", newContract)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(address newContract) returns()
func (_Erc20 *Erc20Session) DelegateToNewContract(newContract common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateToNewContract(&_Erc20.TransactOpts, newContract)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(address newContract) returns()
func (_Erc20 *Erc20TransactorSession) DelegateToNewContract(newContract common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateToNewContract(&_Erc20.TransactOpts, newContract)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(address to, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20Transactor) DelegateTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "delegateTransfer", to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(address to, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20Session) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateTransfer(&_Erc20.TransactOpts, to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(address to, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20TransactorSession) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateTransfer(&_Erc20.TransactOpts, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(address from, address to, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20Transactor) DelegateTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "delegateTransferFrom", from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(address from, address to, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20Session) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateTransferFrom(&_Erc20.TransactOpts, from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(address from, address to, uint256 value, address origSender) returns(bool)
func (_Erc20 *Erc20TransactorSession) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DelegateTransferFrom(&_Erc20.TransactOpts, from, to, value, origSender)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 addedValue) returns(bool)
func (_Erc20 *Erc20Transactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 addedValue) returns(bool)
func (_Erc20 *Erc20Session) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.IncreaseApproval(&_Erc20.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 addedValue) returns(bool)
func (_Erc20 *Erc20TransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.IncreaseApproval(&_Erc20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Erc20 *Erc20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Erc20 *Erc20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Erc20 *Erc20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20Session) Pause() (*types.Transaction, error) {
	return _Erc20.Contract.Pause(&_Erc20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20TransactorSession) Pause() (*types.Transaction, error) {
	return _Erc20.Contract.Pause(&_Erc20.TransactOpts)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(address contractAddr) returns()
func (_Erc20 *Erc20Transactor) ReclaimContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "reclaimContract", contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(address contractAddr) returns()
func (_Erc20 *Erc20Session) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ReclaimContract(&_Erc20.TransactOpts, contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(address contractAddr) returns()
func (_Erc20 *Erc20TransactorSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ReclaimContract(&_Erc20.TransactOpts, contractAddr)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Erc20 *Erc20Transactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Erc20 *Erc20Session) ReclaimEther() (*types.Transaction, error) {
	return _Erc20.Contract.ReclaimEther(&_Erc20.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Erc20 *Erc20TransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _Erc20.Contract.ReclaimEther(&_Erc20.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address token) returns()
func (_Erc20 *Erc20Transactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address token) returns()
func (_Erc20 *Erc20Session) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ReclaimToken(&_Erc20.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(address token) returns()
func (_Erc20 *Erc20TransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ReclaimToken(&_Erc20.TransactOpts, token)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(address sheet) returns()
func (_Erc20 *Erc20Transactor) SetAllowanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setAllowanceSheet", sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(address sheet) returns()
func (_Erc20 *Erc20Session) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetAllowanceSheet(&_Erc20.TransactOpts, sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(address sheet) returns()
func (_Erc20 *Erc20TransactorSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetAllowanceSheet(&_Erc20.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(address sheet) returns()
func (_Erc20 *Erc20Transactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(address sheet) returns()
func (_Erc20 *Erc20Session) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetBalanceSheet(&_Erc20.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(address sheet) returns()
func (_Erc20 *Erc20TransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetBalanceSheet(&_Erc20.TransactOpts, sheet)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(address addr) returns()
func (_Erc20 *Erc20Transactor) SetDelegatedFrom(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setDelegatedFrom", addr)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(address addr) returns()
func (_Erc20 *Erc20Session) SetDelegatedFrom(addr common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetDelegatedFrom(&_Erc20.TransactOpts, addr)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(address addr) returns()
func (_Erc20 *Erc20TransactorSession) SetDelegatedFrom(addr common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetDelegatedFrom(&_Erc20.TransactOpts, addr)
}

// SetLists is a paid mutator transaction binding the contract method 0x0ce51179.
//
// Solidity: function setLists(address _canReceiveMintWhiteList, address _canBurnWhiteList, address _blackList, address _noFeesList) returns()
func (_Erc20 *Erc20Transactor) SetLists(opts *bind.TransactOpts, _canReceiveMintWhiteList common.Address, _canBurnWhiteList common.Address, _blackList common.Address, _noFeesList common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setLists", _canReceiveMintWhiteList, _canBurnWhiteList, _blackList, _noFeesList)
}

// SetLists is a paid mutator transaction binding the contract method 0x0ce51179.
//
// Solidity: function setLists(address _canReceiveMintWhiteList, address _canBurnWhiteList, address _blackList, address _noFeesList) returns()
func (_Erc20 *Erc20Session) SetLists(_canReceiveMintWhiteList common.Address, _canBurnWhiteList common.Address, _blackList common.Address, _noFeesList common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetLists(&_Erc20.TransactOpts, _canReceiveMintWhiteList, _canBurnWhiteList, _blackList, _noFeesList)
}

// SetLists is a paid mutator transaction binding the contract method 0x0ce51179.
//
// Solidity: function setLists(address _canReceiveMintWhiteList, address _canBurnWhiteList, address _blackList, address _noFeesList) returns()
func (_Erc20 *Erc20TransactorSession) SetLists(_canReceiveMintWhiteList common.Address, _canBurnWhiteList common.Address, _blackList common.Address, _noFeesList common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.SetLists(&_Erc20.TransactOpts, _canReceiveMintWhiteList, _canBurnWhiteList, _blackList, _noFeesList)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address from_, uint256 value_, bytes data_) returns()
func (_Erc20 *Erc20Transactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address from_, uint256 value_, bytes data_) returns()
func (_Erc20 *Erc20Session) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Erc20.Contract.TokenFallback(&_Erc20.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address from_, uint256 value_, bytes data_) returns()
func (_Erc20 *Erc20TransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Erc20.Contract.TokenFallback(&_Erc20.TransactOpts, from_, value_, data_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Erc20 *Erc20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Erc20 *Erc20Session) Unpause() (*types.Transaction, error) {
	return _Erc20.Contract.Unpause(&_Erc20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Erc20 *Erc20TransactorSession) Unpause() (*types.Transaction, error) {
	return _Erc20.Contract.Unpause(&_Erc20.TransactOpts)
}

// WipeBlacklistedAccount is a paid mutator transaction binding the contract method 0xbd7243f6.
//
// Solidity: function wipeBlacklistedAccount(address account) returns()
func (_Erc20 *Erc20Transactor) WipeBlacklistedAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "wipeBlacklistedAccount", account)
}

// WipeBlacklistedAccount is a paid mutator transaction binding the contract method 0xbd7243f6.
//
// Solidity: function wipeBlacklistedAccount(address account) returns()
func (_Erc20 *Erc20Session) WipeBlacklistedAccount(account common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.WipeBlacklistedAccount(&_Erc20.TransactOpts, account)
}

// WipeBlacklistedAccount is a paid mutator transaction binding the contract method 0xbd7243f6.
//
// Solidity: function wipeBlacklistedAccount(address account) returns()
func (_Erc20 *Erc20TransactorSession) WipeBlacklistedAccount(account common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.WipeBlacklistedAccount(&_Erc20.TransactOpts, account)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Erc20 *Erc20Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Erc20.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Erc20 *Erc20Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Erc20.Contract.Fallback(&_Erc20.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Erc20 *Erc20TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Erc20.Contract.Fallback(&_Erc20.TransactOpts, calldata)
}

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct {
	Event *Erc20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20ApprovalIterator{contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Erc20 contract.
type Erc20BurnIterator struct {
	Event *Erc20Burn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Burn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20Burn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Burn represents a Burn event raised by the Erc20 contract.
type Erc20Burn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Erc20 *Erc20Filterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*Erc20BurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20BurnIterator{contract: _Erc20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Erc20 *Erc20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *Erc20Burn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Burn)
				if err := _Erc20.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Erc20 *Erc20Filterer) ParseBurn(log types.Log) (*Erc20Burn, error) {
	event := new(Erc20Burn)
	if err := _Erc20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20ChangeBurnBoundsEventIterator is returned from FilterChangeBurnBoundsEvent and is used to iterate over the raw logs and unpacked data for ChangeBurnBoundsEvent events raised by the Erc20 contract.
type Erc20ChangeBurnBoundsEventIterator struct {
	Event *Erc20ChangeBurnBoundsEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20ChangeBurnBoundsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20ChangeBurnBoundsEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20ChangeBurnBoundsEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20ChangeBurnBoundsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ChangeBurnBoundsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20ChangeBurnBoundsEvent represents a ChangeBurnBoundsEvent event raised by the Erc20 contract.
type Erc20ChangeBurnBoundsEvent struct {
	NewMin *big.Int
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeBurnBoundsEvent is a free log retrieval operation binding the contract event 0xf8f7312d8aa9257dcfe43287f24cacc0f267875658809b6c7953b27756562522.
//
// Solidity: event ChangeBurnBoundsEvent(uint256 newMin, uint256 newMax)
func (_Erc20 *Erc20Filterer) FilterChangeBurnBoundsEvent(opts *bind.FilterOpts) (*Erc20ChangeBurnBoundsEventIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "ChangeBurnBoundsEvent")
	if err != nil {
		return nil, err
	}
	return &Erc20ChangeBurnBoundsEventIterator{contract: _Erc20.contract, event: "ChangeBurnBoundsEvent", logs: logs, sub: sub}, nil
}

// WatchChangeBurnBoundsEvent is a free log subscription operation binding the contract event 0xf8f7312d8aa9257dcfe43287f24cacc0f267875658809b6c7953b27756562522.
//
// Solidity: event ChangeBurnBoundsEvent(uint256 newMin, uint256 newMax)
func (_Erc20 *Erc20Filterer) WatchChangeBurnBoundsEvent(opts *bind.WatchOpts, sink chan<- *Erc20ChangeBurnBoundsEvent) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "ChangeBurnBoundsEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20ChangeBurnBoundsEvent)
				if err := _Erc20.contract.UnpackLog(event, "ChangeBurnBoundsEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangeBurnBoundsEvent is a log parse operation binding the contract event 0xf8f7312d8aa9257dcfe43287f24cacc0f267875658809b6c7953b27756562522.
//
// Solidity: event ChangeBurnBoundsEvent(uint256 newMin, uint256 newMax)
func (_Erc20 *Erc20Filterer) ParseChangeBurnBoundsEvent(log types.Log) (*Erc20ChangeBurnBoundsEvent, error) {
	event := new(Erc20ChangeBurnBoundsEvent)
	if err := _Erc20.contract.UnpackLog(event, "ChangeBurnBoundsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20DelegatedToIterator is returned from FilterDelegatedTo and is used to iterate over the raw logs and unpacked data for DelegatedTo events raised by the Erc20 contract.
type Erc20DelegatedToIterator struct {
	Event *Erc20DelegatedTo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20DelegatedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20DelegatedTo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20DelegatedTo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20DelegatedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20DelegatedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20DelegatedTo represents a DelegatedTo event raised by the Erc20 contract.
type Erc20DelegatedTo struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatedTo is a free log retrieval operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: event DelegatedTo(address indexed newContract)
func (_Erc20 *Erc20Filterer) FilterDelegatedTo(opts *bind.FilterOpts, newContract []common.Address) (*Erc20DelegatedToIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "DelegatedTo", newContractRule)
	if err != nil {
		return nil, err
	}
	return &Erc20DelegatedToIterator{contract: _Erc20.contract, event: "DelegatedTo", logs: logs, sub: sub}, nil
}

// WatchDelegatedTo is a free log subscription operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: event DelegatedTo(address indexed newContract)
func (_Erc20 *Erc20Filterer) WatchDelegatedTo(opts *bind.WatchOpts, sink chan<- *Erc20DelegatedTo, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "DelegatedTo", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20DelegatedTo)
				if err := _Erc20.contract.UnpackLog(event, "DelegatedTo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatedTo is a log parse operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: event DelegatedTo(address indexed newContract)
func (_Erc20 *Erc20Filterer) ParseDelegatedTo(log types.Log) (*Erc20DelegatedTo, error) {
	event := new(Erc20DelegatedTo)
	if err := _Erc20.contract.UnpackLog(event, "DelegatedTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Erc20 contract.
type Erc20MintIterator struct {
	Event *Erc20Mint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Mint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20Mint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Mint represents a Mint event raised by the Erc20 contract.
type Erc20Mint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_Erc20 *Erc20Filterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*Erc20MintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20MintIterator{contract: _Erc20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_Erc20 *Erc20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Erc20Mint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Mint)
				if err := _Erc20.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_Erc20 *Erc20Filterer) ParseMint(log types.Log) (*Erc20Mint, error) {
	event := new(Erc20Mint)
	if err := _Erc20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20 contract.
type Erc20OwnershipTransferredIterator struct {
	Event *Erc20OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20OwnershipTransferred represents a OwnershipTransferred event raised by the Erc20 contract.
type Erc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20OwnershipTransferredIterator{contract: _Erc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20OwnershipTransferred)
				if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20 *Erc20Filterer) ParseOwnershipTransferred(log types.Log) (*Erc20OwnershipTransferred, error) {
	event := new(Erc20OwnershipTransferred)
	if err := _Erc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20PauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Erc20 contract.
type Erc20PauseIterator struct {
	Event *Erc20Pause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20PauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Pause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20Pause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20PauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20PauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Pause represents a Pause event raised by the Erc20 contract.
type Erc20Pause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Erc20 *Erc20Filterer) FilterPause(opts *bind.FilterOpts) (*Erc20PauseIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &Erc20PauseIterator{contract: _Erc20.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Erc20 *Erc20Filterer) WatchPause(opts *bind.WatchOpts, sink chan<- *Erc20Pause) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Pause)
				if err := _Erc20.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Erc20 *Erc20Filterer) ParsePause(log types.Log) (*Erc20Pause, error) {
	event := new(Erc20Pause)
	if err := _Erc20.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20UnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Erc20 contract.
type Erc20UnpauseIterator struct {
	Event *Erc20Unpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20UnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Unpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20Unpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20UnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20UnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Unpause represents a Unpause event raised by the Erc20 contract.
type Erc20Unpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Erc20 *Erc20Filterer) FilterUnpause(opts *bind.FilterOpts) (*Erc20UnpauseIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &Erc20UnpauseIterator{contract: _Erc20.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Erc20 *Erc20Filterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *Erc20Unpause) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Unpause)
				if err := _Erc20.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Erc20 *Erc20Filterer) ParseUnpause(log types.Log) (*Erc20Unpause, error) {
	event := new(Erc20Unpause)
	if err := _Erc20.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20WipedAccountIterator is returned from FilterWipedAccount and is used to iterate over the raw logs and unpacked data for WipedAccount events raised by the Erc20 contract.
type Erc20WipedAccountIterator struct {
	Event *Erc20WipedAccount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc20WipedAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20WipedAccount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc20WipedAccount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc20WipedAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20WipedAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20WipedAccount represents a WipedAccount event raised by the Erc20 contract.
type Erc20WipedAccount struct {
	Account common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWipedAccount is a free log retrieval operation binding the contract event 0xdf58d2368c06216a398f05a7a88c8edc64a25c33f33fd2bd8b56fbc8822c02d8.
//
// Solidity: event WipedAccount(address indexed account, uint256 balance)
func (_Erc20 *Erc20Filterer) FilterWipedAccount(opts *bind.FilterOpts, account []common.Address) (*Erc20WipedAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "WipedAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &Erc20WipedAccountIterator{contract: _Erc20.contract, event: "WipedAccount", logs: logs, sub: sub}, nil
}

// WatchWipedAccount is a free log subscription operation binding the contract event 0xdf58d2368c06216a398f05a7a88c8edc64a25c33f33fd2bd8b56fbc8822c02d8.
//
// Solidity: event WipedAccount(address indexed account, uint256 balance)
func (_Erc20 *Erc20Filterer) WatchWipedAccount(opts *bind.WatchOpts, sink chan<- *Erc20WipedAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "WipedAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20WipedAccount)
				if err := _Erc20.contract.UnpackLog(event, "WipedAccount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWipedAccount is a log parse operation binding the contract event 0xdf58d2368c06216a398f05a7a88c8edc64a25c33f33fd2bd8b56fbc8822c02d8.
//
// Solidity: event WipedAccount(address indexed account, uint256 balance)
func (_Erc20 *Erc20Filterer) ParseWipedAccount(log types.Log) (*Erc20WipedAccount, error) {
	event := new(Erc20WipedAccount)
	if err := _Erc20.contract.UnpackLog(event, "WipedAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
