// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20InterfaceABI is the input ABI used to generate the binding from.
const ERC20InterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20InterfaceBin is the compiled bytecode used for deploying new contracts.
const ERC20InterfaceBin = `0x`

// DeployERC20Interface deploys a new Ethereum contract, binding an instance of ERC20Interface to it.
func DeployERC20Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// ERC20Interface is an auto generated Go binding around an Ethereum contract.
type ERC20Interface struct {
	ERC20InterfaceCaller     // Read-only binding to the contract
	ERC20InterfaceTransactor // Write-only binding to the contract
	ERC20InterfaceFilterer   // Log filterer for contract events
}

// ERC20InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InterfaceSession struct {
	Contract     *ERC20Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InterfaceCallerSession struct {
	Contract *ERC20InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InterfaceTransactorSession struct {
	Contract     *ERC20InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InterfaceRaw struct {
	Contract *ERC20Interface // Generic contract binding to access the raw methods on
}

// ERC20InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InterfaceCallerRaw struct {
	Contract *ERC20InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactorRaw struct {
	Contract *ERC20InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Interface creates a new instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20Interface(address common.Address, backend bind.ContractBackend) (*ERC20Interface, error) {
	contract, err := bindERC20Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// NewERC20InterfaceCaller creates a new read-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC20InterfaceCaller, error) {
	contract, err := bindERC20Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceCaller{contract: contract}, nil
}

// NewERC20InterfaceTransactor creates a new write-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InterfaceTransactor, error) {
	contract, err := bindERC20Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransactor{contract: contract}, nil
}

// NewERC20InterfaceFilterer creates a new log filterer instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InterfaceFilterer, error) {
	contract, err := bindERC20Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceFilterer{contract: contract}, nil
}

// bindERC20Interface binds a generic wrapper to an already deployed contract.
func bindERC20Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.ERC20InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.Allowance(&_ERC20Interface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.Allowance(&_ERC20Interface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Interface *ERC20InterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Interface *ERC20InterfaceSession) Decimals() (uint8, error) {
	return _ERC20Interface.Contract.Decimals(&_ERC20Interface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Interface *ERC20InterfaceCallerSession) Decimals() (uint8, error) {
	return _ERC20Interface.Contract.Decimals(&_ERC20Interface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceSession) TotalSupply() (*big.Int, error) {
	return _ERC20Interface.Contract.TotalSupply(&_ERC20Interface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Interface.Contract.TotalSupply(&_ERC20Interface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Approve(&_ERC20Interface.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Approve(&_ERC20Interface.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, from, to, value)
}

// ERC20InterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Interface contract.
type ERC20InterfaceApprovalIterator struct {
	Event *ERC20InterfaceApproval // Event containing the contract specifics and raw log

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
func (it *ERC20InterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterfaceApproval)
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
		it.Event = new(ERC20InterfaceApproval)
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
func (it *ERC20InterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterfaceApproval represents a Approval event raised by the ERC20Interface contract.
type ERC20InterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20InterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Interface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceApprovalIterator{contract: _ERC20Interface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20InterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Interface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterfaceApproval)
				if err := _ERC20Interface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20InterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Interface contract.
type ERC20InterfaceTransferIterator struct {
	Event *ERC20InterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20InterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterfaceTransfer)
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
		it.Event = new(ERC20InterfaceTransfer)
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
func (it *ERC20InterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterfaceTransfer represents a Transfer event raised by the ERC20Interface contract.
type ERC20InterfaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20InterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Interface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransferIterator{contract: _ERC20Interface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20InterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Interface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterfaceTransfer)
				if err := _ERC20Interface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058201d54ee966bad20992905061356126a84fb69b8b4e8feeb4fb880ec1f503df9310029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StaticLiquidityABI is the input ABI used to generate the binding from.
const StaticLiquidityABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"withdrawnTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokensToSell\",\"type\":\"uint256\"}],\"name\":\"sellTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20I\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableLiquidity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableLiquidity\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiPerToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_weiPerToken\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// StaticLiquidityBin is the compiled bytecode used for deploying new contracts.
const StaticLiquidityBin = `0x608060405234801561001057600080fd5b50604051602080610a788339810180604052602081101561003057600080fd5b505160008054600160a060020a0319908116331790915560028054600160a060020a0390931692909116919091179055610a098061006f6000396000f3fe608060405260043610610092577c010000000000000000000000000000000000000000000000000000000060003504630905aa5c81146100945780636c11bcd3146100bd5780636c4cc568146100e75780639e06f57214610118578063a68fee871461012d578063b556f69d14610142578063dab8263a14610157578063db068e0e1461017e578063f851a440146101a8575b005b3480156100a057600080fd5b506100a96101bd565b604080519115158252519081900360200190f35b3480156100c957600080fd5b506100a9600480360360208110156100e057600080fd5b5035610477565b3480156100f357600080fd5b506100fc610736565b60408051600160a060020a039092168252519081900360200190f35b34801561012457600080fd5b506100a9610745565b34801561013957600080fd5b506100a96107c5565b34801561014e57600080fd5b506100a96107d5565b34801561016357600080fd5b5061016c61085c565b60408051918252519081900360200190f35b34801561018a57600080fd5b506100a9600480360360208110156101a157600080fd5b5035610862565b3480156101b457600080fd5b506100fc61091e565b60008054600160a060020a0316331461020e576040805160e560020a62461bcd02815260206004820152601460248201526000805160206109be833981519152604482015290519081900360640190fd5b60005460a060020a900460ff1615610270576040805160e560020a62461bcd02815260206004820152601b60248201527f636f6e7472616374206d757374206e6f74206265206c69717569640000000000604482015290519081900360640190fd5b600254604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600092600160a060020a0316916370a08231916024808301926020929190829003018186803b1580156102d457600080fd5b505afa1580156102e8573d6000803e3d6000fd5b505050506040513d60208110156102fe57600080fd5b5051905060008111610380576040805160e560020a62461bcd02815260206004820152603160248201527f636f6e7472616374206d7573742068617665206120746f6b656e2062616c616e60448201527f63652067726561746572207468616e2030000000000000000000000000000000606482015290519081900360840190fd5b600254604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018490529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b1580156103ed57600080fd5b505af1158015610401573d6000803e3d6000fd5b505050506040513d602081101561041757600080fd5b5051151561046f576040805160e560020a62461bcd02815260206004820152601560248201527f746f6b656e207472616e73666572206661696c65640000000000000000000000604482015290519081900360640190fd5b600191505090565b6000805460a060020a900460ff1615156104db576040805160e560020a62461bcd02815260206004820152601760248201527f636f6e7472616374206d757374206265206c6971756964000000000000000000604482015290519081900360640190fd5b600030311161055a576040805160e560020a62461bcd02815260206004820152603060248201527f636f6e7472616374206d757374206861766520616e206574682062616c616e6360448201527f652067726561746572207468616e203000000000000000000000000000000000606482015290519081900360840190fd5b600254604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b1580156105cd57600080fd5b505af11580156105e1573d6000803e3d6000fd5b505050506040513d60208110156105f757600080fd5b50511515610675576040805160e560020a62461bcd02815260206004820152603b60248201527f6661696c656420746f2065786563757465207472616e7366657246726f6d2c2060448201527f6c696b656c79206e6565647320617070726f76616c2066697273740000000000606482015290519081900360840190fd5b60006106a4670de0b6b3a76400006106986001548661092d90919063ffffffff16565b9063ffffffff6109a616565b905030318111156106ff576040805160e560020a62461bcd02815260206004820152601b60248201527f6e6f7420656e6f75676820657468657265756d20746f2073656e640000000000604482015290519081900360640190fd5b604051339082156108fc029083906000818181858888f1935050505015801561072c573d6000803e3d6000fd5b5060019392505050565b600254600160a060020a031681565b60008054600160a060020a03163314610796576040805160e560020a62461bcd02815260206004820152601460248201526000805160206109be833981519152604482015290519081900360640190fd5b506000805474ff000000000000000000000000000000000000000019169081905560a060020a900460ff161590565b60005460a060020a900460ff1681565b60008054600160a060020a03163314610826576040805160e560020a62461bcd02815260206004820152601460248201526000805160206109be833981519152604482015290519081900360640190fd5b506000805474ff0000000000000000000000000000000000000000191660a060020a90811791829055900460ff16151560011490565b60015481565b60008054600160a060020a031633146108b3576040805160e560020a62461bcd02815260206004820152601460248201526000805160206109be833981519152604482015290519081900360640190fd5b60005460a060020a900460ff1615610915576040805160e560020a62461bcd02815260206004820152601b60248201527f636f6e7472616374206d757374206e6f74206265206c69717569640000000000604482015290519081900360640190fd5b50600190815590565b600054600160a060020a031681565b6000828202831580610949575082848281151561094657fe5b04145b151561099f576040805160e560020a62461bcd02815260206004820152601660248201527f696e76616c6964206d756c7469706c69636174696f6e00000000000000000000604482015290519081900360640190fd5b9392505050565b60008082848115156109b457fe5b0494935050505056fe73656e646572206d7573742062652061646d696e000000000000000000000000a165627a7a723058205135ff9d500f4f33d64fb5ccc144d137c198d09ef85a03bd7e81dda5606f943b0029`

// DeployStaticLiquidity deploys a new Ethereum contract, binding an instance of StaticLiquidity to it.
func DeployStaticLiquidity(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContractAddress common.Address) (common.Address, *types.Transaction, *StaticLiquidity, error) {
	parsed, err := abi.JSON(strings.NewReader(StaticLiquidityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StaticLiquidityBin), backend, _tokenContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StaticLiquidity{StaticLiquidityCaller: StaticLiquidityCaller{contract: contract}, StaticLiquidityTransactor: StaticLiquidityTransactor{contract: contract}, StaticLiquidityFilterer: StaticLiquidityFilterer{contract: contract}}, nil
}

// StaticLiquidity is an auto generated Go binding around an Ethereum contract.
type StaticLiquidity struct {
	StaticLiquidityCaller     // Read-only binding to the contract
	StaticLiquidityTransactor // Write-only binding to the contract
	StaticLiquidityFilterer   // Log filterer for contract events
}

// StaticLiquidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaticLiquidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticLiquidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaticLiquidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticLiquidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaticLiquidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticLiquiditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaticLiquiditySession struct {
	Contract     *StaticLiquidity  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaticLiquidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaticLiquidityCallerSession struct {
	Contract *StaticLiquidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StaticLiquidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaticLiquidityTransactorSession struct {
	Contract     *StaticLiquidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StaticLiquidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaticLiquidityRaw struct {
	Contract *StaticLiquidity // Generic contract binding to access the raw methods on
}

// StaticLiquidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaticLiquidityCallerRaw struct {
	Contract *StaticLiquidityCaller // Generic read-only contract binding to access the raw methods on
}

// StaticLiquidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaticLiquidityTransactorRaw struct {
	Contract *StaticLiquidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaticLiquidity creates a new instance of StaticLiquidity, bound to a specific deployed contract.
func NewStaticLiquidity(address common.Address, backend bind.ContractBackend) (*StaticLiquidity, error) {
	contract, err := bindStaticLiquidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaticLiquidity{StaticLiquidityCaller: StaticLiquidityCaller{contract: contract}, StaticLiquidityTransactor: StaticLiquidityTransactor{contract: contract}, StaticLiquidityFilterer: StaticLiquidityFilterer{contract: contract}}, nil
}

// NewStaticLiquidityCaller creates a new read-only instance of StaticLiquidity, bound to a specific deployed contract.
func NewStaticLiquidityCaller(address common.Address, caller bind.ContractCaller) (*StaticLiquidityCaller, error) {
	contract, err := bindStaticLiquidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaticLiquidityCaller{contract: contract}, nil
}

// NewStaticLiquidityTransactor creates a new write-only instance of StaticLiquidity, bound to a specific deployed contract.
func NewStaticLiquidityTransactor(address common.Address, transactor bind.ContractTransactor) (*StaticLiquidityTransactor, error) {
	contract, err := bindStaticLiquidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaticLiquidityTransactor{contract: contract}, nil
}

// NewStaticLiquidityFilterer creates a new log filterer instance of StaticLiquidity, bound to a specific deployed contract.
func NewStaticLiquidityFilterer(address common.Address, filterer bind.ContractFilterer) (*StaticLiquidityFilterer, error) {
	contract, err := bindStaticLiquidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaticLiquidityFilterer{contract: contract}, nil
}

// bindStaticLiquidity binds a generic wrapper to an already deployed contract.
func bindStaticLiquidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StaticLiquidityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticLiquidity *StaticLiquidityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StaticLiquidity.Contract.StaticLiquidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticLiquidity *StaticLiquidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.StaticLiquidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticLiquidity *StaticLiquidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.StaticLiquidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticLiquidity *StaticLiquidityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StaticLiquidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticLiquidity *StaticLiquidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticLiquidity *StaticLiquidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_StaticLiquidity *StaticLiquidityCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StaticLiquidity.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_StaticLiquidity *StaticLiquiditySession) Admin() (common.Address, error) {
	return _StaticLiquidity.Contract.Admin(&_StaticLiquidity.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_StaticLiquidity *StaticLiquidityCallerSession) Admin() (common.Address, error) {
	return _StaticLiquidity.Contract.Admin(&_StaticLiquidity.CallOpts)
}

// Erc20I is a free data retrieval call binding the contract method 0x6c4cc568.
//
// Solidity: function erc20I() constant returns(address)
func (_StaticLiquidity *StaticLiquidityCaller) Erc20I(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StaticLiquidity.contract.Call(opts, out, "erc20I")
	return *ret0, err
}

// Erc20I is a free data retrieval call binding the contract method 0x6c4cc568.
//
// Solidity: function erc20I() constant returns(address)
func (_StaticLiquidity *StaticLiquiditySession) Erc20I() (common.Address, error) {
	return _StaticLiquidity.Contract.Erc20I(&_StaticLiquidity.CallOpts)
}

// Erc20I is a free data retrieval call binding the contract method 0x6c4cc568.
//
// Solidity: function erc20I() constant returns(address)
func (_StaticLiquidity *StaticLiquidityCallerSession) Erc20I() (common.Address, error) {
	return _StaticLiquidity.Contract.Erc20I(&_StaticLiquidity.CallOpts)
}

// Liquid is a free data retrieval call binding the contract method 0xa68fee87.
//
// Solidity: function liquid() constant returns(bool)
func (_StaticLiquidity *StaticLiquidityCaller) Liquid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StaticLiquidity.contract.Call(opts, out, "liquid")
	return *ret0, err
}

// Liquid is a free data retrieval call binding the contract method 0xa68fee87.
//
// Solidity: function liquid() constant returns(bool)
func (_StaticLiquidity *StaticLiquiditySession) Liquid() (bool, error) {
	return _StaticLiquidity.Contract.Liquid(&_StaticLiquidity.CallOpts)
}

// Liquid is a free data retrieval call binding the contract method 0xa68fee87.
//
// Solidity: function liquid() constant returns(bool)
func (_StaticLiquidity *StaticLiquidityCallerSession) Liquid() (bool, error) {
	return _StaticLiquidity.Contract.Liquid(&_StaticLiquidity.CallOpts)
}

// WeiPerToken is a free data retrieval call binding the contract method 0xdab8263a.
//
// Solidity: function weiPerToken() constant returns(uint256)
func (_StaticLiquidity *StaticLiquidityCaller) WeiPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StaticLiquidity.contract.Call(opts, out, "weiPerToken")
	return *ret0, err
}

// WeiPerToken is a free data retrieval call binding the contract method 0xdab8263a.
//
// Solidity: function weiPerToken() constant returns(uint256)
func (_StaticLiquidity *StaticLiquiditySession) WeiPerToken() (*big.Int, error) {
	return _StaticLiquidity.Contract.WeiPerToken(&_StaticLiquidity.CallOpts)
}

// WeiPerToken is a free data retrieval call binding the contract method 0xdab8263a.
//
// Solidity: function weiPerToken() constant returns(uint256)
func (_StaticLiquidity *StaticLiquidityCallerSession) WeiPerToken() (*big.Int, error) {
	return _StaticLiquidity.Contract.WeiPerToken(&_StaticLiquidity.CallOpts)
}

// DisableLiquidity is a paid mutator transaction binding the contract method 0x9e06f572.
//
// Solidity: function disableLiquidity() returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactor) DisableLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticLiquidity.contract.Transact(opts, "disableLiquidity")
}

// DisableLiquidity is a paid mutator transaction binding the contract method 0x9e06f572.
//
// Solidity: function disableLiquidity() returns(bool)
func (_StaticLiquidity *StaticLiquiditySession) DisableLiquidity() (*types.Transaction, error) {
	return _StaticLiquidity.Contract.DisableLiquidity(&_StaticLiquidity.TransactOpts)
}

// DisableLiquidity is a paid mutator transaction binding the contract method 0x9e06f572.
//
// Solidity: function disableLiquidity() returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactorSession) DisableLiquidity() (*types.Transaction, error) {
	return _StaticLiquidity.Contract.DisableLiquidity(&_StaticLiquidity.TransactOpts)
}

// EnableLiquidity is a paid mutator transaction binding the contract method 0xb556f69d.
//
// Solidity: function enableLiquidity() returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactor) EnableLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticLiquidity.contract.Transact(opts, "enableLiquidity")
}

// EnableLiquidity is a paid mutator transaction binding the contract method 0xb556f69d.
//
// Solidity: function enableLiquidity() returns(bool)
func (_StaticLiquidity *StaticLiquiditySession) EnableLiquidity() (*types.Transaction, error) {
	return _StaticLiquidity.Contract.EnableLiquidity(&_StaticLiquidity.TransactOpts)
}

// EnableLiquidity is a paid mutator transaction binding the contract method 0xb556f69d.
//
// Solidity: function enableLiquidity() returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactorSession) EnableLiquidity() (*types.Transaction, error) {
	return _StaticLiquidity.Contract.EnableLiquidity(&_StaticLiquidity.TransactOpts)
}

// SellTokens is a paid mutator transaction binding the contract method 0x6c11bcd3.
//
// Solidity: function sellTokens(_tokensToSell uint256) returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactor) SellTokens(opts *bind.TransactOpts, _tokensToSell *big.Int) (*types.Transaction, error) {
	return _StaticLiquidity.contract.Transact(opts, "sellTokens", _tokensToSell)
}

// SellTokens is a paid mutator transaction binding the contract method 0x6c11bcd3.
//
// Solidity: function sellTokens(_tokensToSell uint256) returns(bool)
func (_StaticLiquidity *StaticLiquiditySession) SellTokens(_tokensToSell *big.Int) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.SellTokens(&_StaticLiquidity.TransactOpts, _tokensToSell)
}

// SellTokens is a paid mutator transaction binding the contract method 0x6c11bcd3.
//
// Solidity: function sellTokens(_tokensToSell uint256) returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactorSession) SellTokens(_tokensToSell *big.Int) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.SellTokens(&_StaticLiquidity.TransactOpts, _tokensToSell)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(_weiPerToken uint256) returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactor) SetExchangeRate(opts *bind.TransactOpts, _weiPerToken *big.Int) (*types.Transaction, error) {
	return _StaticLiquidity.contract.Transact(opts, "setExchangeRate", _weiPerToken)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(_weiPerToken uint256) returns(bool)
func (_StaticLiquidity *StaticLiquiditySession) SetExchangeRate(_weiPerToken *big.Int) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.SetExchangeRate(&_StaticLiquidity.TransactOpts, _weiPerToken)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(_weiPerToken uint256) returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactorSession) SetExchangeRate(_weiPerToken *big.Int) (*types.Transaction, error) {
	return _StaticLiquidity.Contract.SetExchangeRate(&_StaticLiquidity.TransactOpts, _weiPerToken)
}

// WithdrawnTokens is a paid mutator transaction binding the contract method 0x0905aa5c.
//
// Solidity: function withdrawnTokens() returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactor) WithdrawnTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticLiquidity.contract.Transact(opts, "withdrawnTokens")
}

// WithdrawnTokens is a paid mutator transaction binding the contract method 0x0905aa5c.
//
// Solidity: function withdrawnTokens() returns(bool)
func (_StaticLiquidity *StaticLiquiditySession) WithdrawnTokens() (*types.Transaction, error) {
	return _StaticLiquidity.Contract.WithdrawnTokens(&_StaticLiquidity.TransactOpts)
}

// WithdrawnTokens is a paid mutator transaction binding the contract method 0x0905aa5c.
//
// Solidity: function withdrawnTokens() returns(bool)
func (_StaticLiquidity *StaticLiquidityTransactorSession) WithdrawnTokens() (*types.Transaction, error) {
	return _StaticLiquidity.Contract.WithdrawnTokens(&_StaticLiquidity.TransactOpts)
}
