// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveMetaData contains all meta data concerning the Aave contract.
var AaveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AssetSourceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\"}],\"name\":\"BaseCurrencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"FallbackOracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CURRENCY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"getAssetsPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSourceOfAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"}],\"name\":\"setAssetSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"setFallbackOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveMetaData.ABI instead.
var AaveABI = AaveMetaData.ABI

// Aave is an auto generated Go binding around an Ethereum contract.
type Aave struct {
	AaveCaller     // Read-only binding to the contract
	AaveTransactor // Write-only binding to the contract
	AaveFilterer   // Log filterer for contract events
}

// AaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveSession struct {
	Contract     *Aave             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveCallerSession struct {
	Contract *AaveCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveTransactorSession struct {
	Contract     *AaveTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveRaw struct {
	Contract *Aave // Generic contract binding to access the raw methods on
}

// AaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveCallerRaw struct {
	Contract *AaveCaller // Generic read-only contract binding to access the raw methods on
}

// AaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveTransactorRaw struct {
	Contract *AaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAave creates a new instance of Aave, bound to a specific deployed contract.
func NewAave(address common.Address, backend bind.ContractBackend) (*Aave, error) {
	contract, err := bindAave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// NewAaveCaller creates a new read-only instance of Aave, bound to a specific deployed contract.
func NewAaveCaller(address common.Address, caller bind.ContractCaller) (*AaveCaller, error) {
	contract, err := bindAave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveCaller{contract: contract}, nil
}

// NewAaveTransactor creates a new write-only instance of Aave, bound to a specific deployed contract.
func NewAaveTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveTransactor, error) {
	contract, err := bindAave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTransactor{contract: contract}, nil
}

// NewAaveFilterer creates a new log filterer instance of Aave, bound to a specific deployed contract.
func NewAaveFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveFilterer, error) {
	contract, err := bindAave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveFilterer{contract: contract}, nil
}

// bindAave binds a generic wrapper to an already deployed contract.
func bindAave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.AaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Aave *AaveCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Aave *AaveSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Aave.Contract.ADDRESSESPROVIDER(&_Aave.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Aave *AaveCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Aave.Contract.ADDRESSESPROVIDER(&_Aave.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_Aave *AaveCaller) BASECURRENCY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "BASE_CURRENCY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_Aave *AaveSession) BASECURRENCY() (common.Address, error) {
	return _Aave.Contract.BASECURRENCY(&_Aave.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_Aave *AaveCallerSession) BASECURRENCY() (common.Address, error) {
	return _Aave.Contract.BASECURRENCY(&_Aave.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_Aave *AaveCaller) BASECURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "BASE_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_Aave *AaveSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _Aave.Contract.BASECURRENCYUNIT(&_Aave.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_Aave *AaveCallerSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _Aave.Contract.BASECURRENCYUNIT(&_Aave.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Aave *AaveCaller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getAssetPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Aave *AaveSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _Aave.Contract.GetAssetPrice(&_Aave.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Aave *AaveCallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _Aave.Contract.GetAssetPrice(&_Aave.CallOpts, asset)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_Aave *AaveCaller) GetAssetsPrices(opts *bind.CallOpts, assets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getAssetsPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_Aave *AaveSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _Aave.Contract.GetAssetsPrices(&_Aave.CallOpts, assets)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_Aave *AaveCallerSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _Aave.Contract.GetAssetsPrices(&_Aave.CallOpts, assets)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_Aave *AaveCaller) GetFallbackOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getFallbackOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_Aave *AaveSession) GetFallbackOracle() (common.Address, error) {
	return _Aave.Contract.GetFallbackOracle(&_Aave.CallOpts)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_Aave *AaveCallerSession) GetFallbackOracle() (common.Address, error) {
	return _Aave.Contract.GetFallbackOracle(&_Aave.CallOpts)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_Aave *AaveCaller) GetSourceOfAsset(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getSourceOfAsset", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_Aave *AaveSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _Aave.Contract.GetSourceOfAsset(&_Aave.CallOpts, asset)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_Aave *AaveCallerSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _Aave.Contract.GetSourceOfAsset(&_Aave.CallOpts, asset)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_Aave *AaveTransactor) SetAssetSources(opts *bind.TransactOpts, assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setAssetSources", assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_Aave *AaveSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetAssetSources(&_Aave.TransactOpts, assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_Aave *AaveTransactorSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetAssetSources(&_Aave.TransactOpts, assets, sources)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_Aave *AaveTransactor) SetFallbackOracle(opts *bind.TransactOpts, fallbackOracle common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setFallbackOracle", fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_Aave *AaveSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetFallbackOracle(&_Aave.TransactOpts, fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_Aave *AaveTransactorSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetFallbackOracle(&_Aave.TransactOpts, fallbackOracle)
}

// AaveAssetSourceUpdatedIterator is returned from FilterAssetSourceUpdated and is used to iterate over the raw logs and unpacked data for AssetSourceUpdated events raised by the Aave contract.
type AaveAssetSourceUpdatedIterator struct {
	Event *AaveAssetSourceUpdated // Event containing the contract specifics and raw log

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
func (it *AaveAssetSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAssetSourceUpdated)
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
		it.Event = new(AaveAssetSourceUpdated)
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
func (it *AaveAssetSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAssetSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAssetSourceUpdated represents a AssetSourceUpdated event raised by the Aave contract.
type AaveAssetSourceUpdated struct {
	Asset  common.Address
	Source common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetSourceUpdated is a free log retrieval operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_Aave *AaveFilterer) FilterAssetSourceUpdated(opts *bind.FilterOpts, asset []common.Address, source []common.Address) (*AaveAssetSourceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &AaveAssetSourceUpdatedIterator{contract: _Aave.contract, event: "AssetSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetSourceUpdated is a free log subscription operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_Aave *AaveFilterer) WatchAssetSourceUpdated(opts *bind.WatchOpts, sink chan<- *AaveAssetSourceUpdated, asset []common.Address, source []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAssetSourceUpdated)
				if err := _Aave.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
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

// ParseAssetSourceUpdated is a log parse operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_Aave *AaveFilterer) ParseAssetSourceUpdated(log types.Log) (*AaveAssetSourceUpdated, error) {
	event := new(AaveAssetSourceUpdated)
	if err := _Aave.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveBaseCurrencySetIterator is returned from FilterBaseCurrencySet and is used to iterate over the raw logs and unpacked data for BaseCurrencySet events raised by the Aave contract.
type AaveBaseCurrencySetIterator struct {
	Event *AaveBaseCurrencySet // Event containing the contract specifics and raw log

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
func (it *AaveBaseCurrencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveBaseCurrencySet)
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
		it.Event = new(AaveBaseCurrencySet)
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
func (it *AaveBaseCurrencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveBaseCurrencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveBaseCurrencySet represents a BaseCurrencySet event raised by the Aave contract.
type AaveBaseCurrencySet struct {
	BaseCurrency     common.Address
	BaseCurrencyUnit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBaseCurrencySet is a free log retrieval operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_Aave *AaveFilterer) FilterBaseCurrencySet(opts *bind.FilterOpts, baseCurrency []common.Address) (*AaveBaseCurrencySetIterator, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &AaveBaseCurrencySetIterator{contract: _Aave.contract, event: "BaseCurrencySet", logs: logs, sub: sub}, nil
}

// WatchBaseCurrencySet is a free log subscription operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_Aave *AaveFilterer) WatchBaseCurrencySet(opts *bind.WatchOpts, sink chan<- *AaveBaseCurrencySet, baseCurrency []common.Address) (event.Subscription, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveBaseCurrencySet)
				if err := _Aave.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
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

// ParseBaseCurrencySet is a log parse operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_Aave *AaveFilterer) ParseBaseCurrencySet(log types.Log) (*AaveBaseCurrencySet, error) {
	event := new(AaveBaseCurrencySet)
	if err := _Aave.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveFallbackOracleUpdatedIterator is returned from FilterFallbackOracleUpdated and is used to iterate over the raw logs and unpacked data for FallbackOracleUpdated events raised by the Aave contract.
type AaveFallbackOracleUpdatedIterator struct {
	Event *AaveFallbackOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AaveFallbackOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveFallbackOracleUpdated)
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
		it.Event = new(AaveFallbackOracleUpdated)
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
func (it *AaveFallbackOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveFallbackOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveFallbackOracleUpdated represents a FallbackOracleUpdated event raised by the Aave contract.
type AaveFallbackOracleUpdated struct {
	FallbackOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFallbackOracleUpdated is a free log retrieval operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_Aave *AaveFilterer) FilterFallbackOracleUpdated(opts *bind.FilterOpts, fallbackOracle []common.Address) (*AaveFallbackOracleUpdatedIterator, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return &AaveFallbackOracleUpdatedIterator{contract: _Aave.contract, event: "FallbackOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchFallbackOracleUpdated is a free log subscription operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_Aave *AaveFilterer) WatchFallbackOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveFallbackOracleUpdated, fallbackOracle []common.Address) (event.Subscription, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveFallbackOracleUpdated)
				if err := _Aave.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
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

// ParseFallbackOracleUpdated is a log parse operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_Aave *AaveFilterer) ParseFallbackOracleUpdated(log types.Log) (*AaveFallbackOracleUpdated, error) {
	event := new(AaveFallbackOracleUpdated)
	if err := _Aave.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
