// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"foo1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"foo2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"foo3\",\"type\":\"string\"}],\"name\":\"FooChangeEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localFoo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"foo1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"foo2\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"foo3\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_foo1\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_foo3\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_foo2\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001607b81526020016040518060400160405280600c81526020017f48656c6c6f2c20776f726c6400000000000000000000000000000000000000008152508152505f5f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020190816100db9190610321565b509050506103f0565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061015f57607f821691505b6020821081036101725761017161011b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101d47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610199565b6101de8683610199565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61022261021d610218846101f6565b6101ff565b6101f6565b9050919050565b5f819050919050565b61023b83610208565b61024f61024782610229565b8484546101a5565b825550505050565b5f5f905090565b610266610257565b610271818484610232565b505050565b5b81811015610294576102895f8261025e565b600181019050610277565b5050565b601f8211156102d9576102aa81610178565b6102b38461018a565b810160208510156102c2578190505b6102d66102ce8561018a565b830182610276565b50505b505050565b5f82821c905092915050565b5f6102f95f19846008026102de565b1980831691505092915050565b5f61031183836102ea565b9150826002028217905092915050565b61032a826100e4565b67ffffffffffffffff811115610343576103426100ee565b5b61034d8254610148565b610358828285610298565b5f60209050601f831160018114610389575f8415610377578287015190505b6103818582610306565b8655506103e8565b601f19841661039786610178565b5f5b828110156103be57848901518255600182019150602085019450602081019050610399565b868310156103db57848901516103d7601f8916826102ea565b8355505b6001600288020188555050505b505050505050565b61097a806103fd5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80634ef65c3b1461004e57806379ecc4c51461006a5780637fcaf6661461008a578063e30081a0146100a6575b5f5ffd5b61006860048036038101906100639190610352565b6100c2565b005b610072610134565b6040516100819392919061043b565b60405180910390f35b6100a4600480360381019061009f91906104d8565b6101ef565b005b6100c060048036038101906100bb919061054d565b61026d565b005b805f600101819055507ff71f2a87bb161c10be2205ab55cbe51abcc7adc24fd0e1c0c22452d6a372bf10335f5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff165f600101545f6002016040516101299493929190610668565b60405180910390a150565b5f805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201805461016e906105a5565b80601f016020809104026020016040519081016040528092919081815260200182805461019a906105a5565b80156101e55780601f106101bc576101008083540402835291602001916101e5565b820191905f5260205f20905b8154815290600101906020018083116101c857829003601f168201915b5050505050905083565b81815f6002019182610202929190610877565b507ff71f2a87bb161c10be2205ab55cbe51abcc7adc24fd0e1c0c22452d6a372bf10335f5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff165f600101545f6002016040516102619493929190610668565b60405180910390a15050565b805f5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ff71f2a87bb161c10be2205ab55cbe51abcc7adc24fd0e1c0c22452d6a372bf10335f5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff165f600101545f60020160405161030c9493929190610668565b60405180910390a150565b5f5ffd5b5f5ffd5b5f819050919050565b6103318161031f565b811461033b575f5ffd5b50565b5f8135905061034c81610328565b92915050565b5f6020828403121561036757610366610317565b5b5f6103748482850161033e565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103a68261037d565b9050919050565b6103b68161039c565b82525050565b6103c58161031f565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61040d826103cb565b61041781856103d5565b93506104278185602086016103e5565b610430816103f3565b840191505092915050565b5f60608201905061044e5f8301866103ad565b61045b60208301856103bc565b818103604083015261046d8184610403565b9050949350505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261049857610497610477565b5b8235905067ffffffffffffffff8111156104b5576104b461047b565b5b6020830191508360018202830111156104d1576104d061047f565b5b9250929050565b5f5f602083850312156104ee576104ed610317565b5b5f83013567ffffffffffffffff81111561050b5761050a61031b565b5b61051785828601610483565b92509250509250929050565b61052c8161039c565b8114610536575f5ffd5b50565b5f8135905061054781610523565b92915050565b5f6020828403121561056257610561610317565b5b5f61056f84828501610539565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105bc57607f821691505b6020821081036105cf576105ce610578565b5b50919050565b5f819050815f5260205f209050919050565b5f81546105f3816105a5565b6105fd81866103d5565b9450600182165f8114610617576001811461062d5761065f565b60ff19831686528115156020028601935061065f565b610636856105d5565b5f5b8381101561065757815481890152600182019150602081019050610638565b808801955050505b50505092915050565b5f60808201905061067b5f8301876103ad565b61068860208301866103ad565b61069560408301856103bc565b81810360608301526106a781846105e7565b905095945050505050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026107337fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826106f8565b61073d86836106f8565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61077861077361076e8461031f565b610755565b61031f565b9050919050565b5f819050919050565b6107918361075e565b6107a561079d8261077f565b848454610704565b825550505050565b5f5f905090565b6107bc6107ad565b6107c7818484610788565b505050565b5b818110156107ea576107df5f826107b4565b6001810190506107cd565b5050565b601f82111561082f57610800816105d5565b610809846106e9565b81016020851015610818578190505b61082c610824856106e9565b8301826107cc565b50505b505050565b5f82821c905092915050565b5f61084f5f1984600802610834565b1980831691505092915050565b5f6108678383610840565b9150826002028217905092915050565b61088183836106b2565b67ffffffffffffffff81111561089a576108996106bc565b5b6108a482546105a5565b6108af8282856107ee565b5f601f8311600181146108dc575f84156108ca578287013590505b6108d4858261085c565b86555061093b565b601f1984166108ea866105d5565b5f5b82811015610911578489013582556001820191506020850194506020810190506108ec565b8683101561092e578489013561092a601f891682610840565b8355505b6001600288020188555050505b5050505050505056fea26469706673582212209b5408234fea24f90e3c49ce70a6adf1c51eb59a90b59c7296e829de195b216564736f6c634300081c0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// LocalFoo is a free data retrieval call binding the contract method 0x79ecc4c5.
//
// Solidity: function localFoo() view returns(address foo1, uint256 foo2, string foo3)
func (_Contracts *ContractsCaller) LocalFoo(opts *bind.CallOpts) (struct {
	Foo1 common.Address
	Foo2 *big.Int
	Foo3 string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "localFoo")

	outstruct := new(struct {
		Foo1 common.Address
		Foo2 *big.Int
		Foo3 string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Foo1 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Foo2 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Foo3 = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// LocalFoo is a free data retrieval call binding the contract method 0x79ecc4c5.
//
// Solidity: function localFoo() view returns(address foo1, uint256 foo2, string foo3)
func (_Contracts *ContractsSession) LocalFoo() (struct {
	Foo1 common.Address
	Foo2 *big.Int
	Foo3 string
}, error) {
	return _Contracts.Contract.LocalFoo(&_Contracts.CallOpts)
}

// LocalFoo is a free data retrieval call binding the contract method 0x79ecc4c5.
//
// Solidity: function localFoo() view returns(address foo1, uint256 foo2, string foo3)
func (_Contracts *ContractsCallerSession) LocalFoo() (struct {
	Foo1 common.Address
	Foo2 *big.Int
	Foo3 string
}, error) {
	return _Contracts.Contract.LocalFoo(&_Contracts.CallOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _foo1) returns()
func (_Contracts *ContractsTransactor) SetAddress(opts *bind.TransactOpts, _foo1 common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setAddress", _foo1)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _foo1) returns()
func (_Contracts *ContractsSession) SetAddress(_foo1 common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetAddress(&_Contracts.TransactOpts, _foo1)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _foo1) returns()
func (_Contracts *ContractsTransactorSession) SetAddress(_foo1 common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetAddress(&_Contracts.TransactOpts, _foo1)
}

// SetString is a paid mutator transaction binding the contract method 0x7fcaf666.
//
// Solidity: function setString(string _foo3) returns()
func (_Contracts *ContractsTransactor) SetString(opts *bind.TransactOpts, _foo3 string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setString", _foo3)
}

// SetString is a paid mutator transaction binding the contract method 0x7fcaf666.
//
// Solidity: function setString(string _foo3) returns()
func (_Contracts *ContractsSession) SetString(_foo3 string) (*types.Transaction, error) {
	return _Contracts.Contract.SetString(&_Contracts.TransactOpts, _foo3)
}

// SetString is a paid mutator transaction binding the contract method 0x7fcaf666.
//
// Solidity: function setString(string _foo3) returns()
func (_Contracts *ContractsTransactorSession) SetString(_foo3 string) (*types.Transaction, error) {
	return _Contracts.Contract.SetString(&_Contracts.TransactOpts, _foo3)
}

// SetUint is a paid mutator transaction binding the contract method 0x4ef65c3b.
//
// Solidity: function setUint(uint256 _foo2) returns()
func (_Contracts *ContractsTransactor) SetUint(opts *bind.TransactOpts, _foo2 *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setUint", _foo2)
}

// SetUint is a paid mutator transaction binding the contract method 0x4ef65c3b.
//
// Solidity: function setUint(uint256 _foo2) returns()
func (_Contracts *ContractsSession) SetUint(_foo2 *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetUint(&_Contracts.TransactOpts, _foo2)
}

// SetUint is a paid mutator transaction binding the contract method 0x4ef65c3b.
//
// Solidity: function setUint(uint256 _foo2) returns()
func (_Contracts *ContractsTransactorSession) SetUint(_foo2 *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetUint(&_Contracts.TransactOpts, _foo2)
}

// ContractsFooChangeEventIterator is returned from FilterFooChangeEvent and is used to iterate over the raw logs and unpacked data for FooChangeEvent events raised by the Contracts contract.
type ContractsFooChangeEventIterator struct {
	Event *ContractsFooChangeEvent // Event containing the contract specifics and raw log

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
func (it *ContractsFooChangeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFooChangeEvent)
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
		it.Event = new(ContractsFooChangeEvent)
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
func (it *ContractsFooChangeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFooChangeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFooChangeEvent represents a FooChangeEvent event raised by the Contracts contract.
type ContractsFooChangeEvent struct {
	Sender common.Address
	Foo1   common.Address
	Foo2   *big.Int
	Foo3   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFooChangeEvent is a free log retrieval operation binding the contract event 0xf71f2a87bb161c10be2205ab55cbe51abcc7adc24fd0e1c0c22452d6a372bf10.
//
// Solidity: event FooChangeEvent(address sender, address foo1, uint256 foo2, string foo3)
func (_Contracts *ContractsFilterer) FilterFooChangeEvent(opts *bind.FilterOpts) (*ContractsFooChangeEventIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "FooChangeEvent")
	if err != nil {
		return nil, err
	}
	return &ContractsFooChangeEventIterator{contract: _Contracts.contract, event: "FooChangeEvent", logs: logs, sub: sub}, nil
}

// WatchFooChangeEvent is a free log subscription operation binding the contract event 0xf71f2a87bb161c10be2205ab55cbe51abcc7adc24fd0e1c0c22452d6a372bf10.
//
// Solidity: event FooChangeEvent(address sender, address foo1, uint256 foo2, string foo3)
func (_Contracts *ContractsFilterer) WatchFooChangeEvent(opts *bind.WatchOpts, sink chan<- *ContractsFooChangeEvent) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "FooChangeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFooChangeEvent)
				if err := _Contracts.contract.UnpackLog(event, "FooChangeEvent", log); err != nil {
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

// ParseFooChangeEvent is a log parse operation binding the contract event 0xf71f2a87bb161c10be2205ab55cbe51abcc7adc24fd0e1c0c22452d6a372bf10.
//
// Solidity: event FooChangeEvent(address sender, address foo1, uint256 foo2, string foo3)
func (_Contracts *ContractsFilterer) ParseFooChangeEvent(log types.Log) (*ContractsFooChangeEvent, error) {
	event := new(ContractsFooChangeEvent)
	if err := _Contracts.contract.UnpackLog(event, "FooChangeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
