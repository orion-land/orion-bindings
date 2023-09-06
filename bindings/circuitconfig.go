// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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

// CircuitConfigMetaData contains all meta data concerning the CircuitConfig contract.
var CircuitConfigMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50601680601d6000396000f3fe6080604052600080fdfea164736f6c634300080f000a",
}

// CircuitConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use CircuitConfigMetaData.ABI instead.
var CircuitConfigABI = CircuitConfigMetaData.ABI

// CircuitConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CircuitConfigMetaData.Bin instead.
var CircuitConfigBin = CircuitConfigMetaData.Bin

// DeployCircuitConfig deploys a new Ethereum contract, binding an instance of CircuitConfig to it.
func DeployCircuitConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CircuitConfig, error) {
	parsed, err := CircuitConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CircuitConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CircuitConfig{CircuitConfigCaller: CircuitConfigCaller{contract: contract}, CircuitConfigTransactor: CircuitConfigTransactor{contract: contract}, CircuitConfigFilterer: CircuitConfigFilterer{contract: contract}}, nil
}

// CircuitConfig is an auto generated Go binding around an Ethereum contract.
type CircuitConfig struct {
	CircuitConfigCaller     // Read-only binding to the contract
	CircuitConfigTransactor // Write-only binding to the contract
	CircuitConfigFilterer   // Log filterer for contract events
}

// CircuitConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type CircuitConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CircuitConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CircuitConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CircuitConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CircuitConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CircuitConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CircuitConfigSession struct {
	Contract     *CircuitConfig    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CircuitConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CircuitConfigCallerSession struct {
	Contract *CircuitConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CircuitConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CircuitConfigTransactorSession struct {
	Contract     *CircuitConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CircuitConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type CircuitConfigRaw struct {
	Contract *CircuitConfig // Generic contract binding to access the raw methods on
}

// CircuitConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CircuitConfigCallerRaw struct {
	Contract *CircuitConfigCaller // Generic read-only contract binding to access the raw methods on
}

// CircuitConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CircuitConfigTransactorRaw struct {
	Contract *CircuitConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCircuitConfig creates a new instance of CircuitConfig, bound to a specific deployed contract.
func NewCircuitConfig(address common.Address, backend bind.ContractBackend) (*CircuitConfig, error) {
	contract, err := bindCircuitConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CircuitConfig{CircuitConfigCaller: CircuitConfigCaller{contract: contract}, CircuitConfigTransactor: CircuitConfigTransactor{contract: contract}, CircuitConfigFilterer: CircuitConfigFilterer{contract: contract}}, nil
}

// NewCircuitConfigCaller creates a new read-only instance of CircuitConfig, bound to a specific deployed contract.
func NewCircuitConfigCaller(address common.Address, caller bind.ContractCaller) (*CircuitConfigCaller, error) {
	contract, err := bindCircuitConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CircuitConfigCaller{contract: contract}, nil
}

// NewCircuitConfigTransactor creates a new write-only instance of CircuitConfig, bound to a specific deployed contract.
func NewCircuitConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*CircuitConfigTransactor, error) {
	contract, err := bindCircuitConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CircuitConfigTransactor{contract: contract}, nil
}

// NewCircuitConfigFilterer creates a new log filterer instance of CircuitConfig, bound to a specific deployed contract.
func NewCircuitConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*CircuitConfigFilterer, error) {
	contract, err := bindCircuitConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CircuitConfigFilterer{contract: contract}, nil
}

// bindCircuitConfig binds a generic wrapper to an already deployed contract.
func bindCircuitConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CircuitConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CircuitConfig *CircuitConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CircuitConfig.Contract.CircuitConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CircuitConfig *CircuitConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CircuitConfig.Contract.CircuitConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CircuitConfig *CircuitConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CircuitConfig.Contract.CircuitConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CircuitConfig *CircuitConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CircuitConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CircuitConfig *CircuitConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CircuitConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CircuitConfig *CircuitConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CircuitConfig.Contract.contract.Transact(opts, method, params...)
}
