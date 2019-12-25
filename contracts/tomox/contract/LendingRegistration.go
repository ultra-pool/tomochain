// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/tomochain/tomochain/accounts/abi"
	"github.com/tomochain/tomochain/accounts/abi/bind"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/core/types"
)

// AbstractRegistrationABI is the input ABI used to generate the binding from.
const AbstractRegistrationABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"RESIGN_REQUESTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRelayerByCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbstractRegistrationBin is the compiled bytecode used for deploying new contracts.
const AbstractRegistrationBin = `0x`

// DeployAbstractRegistration deploys a new Ethereum contract, binding an instance of AbstractRegistration to it.
func DeployAbstractRegistration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AbstractRegistration, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractRegistrationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbstractRegistrationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AbstractRegistration{AbstractRegistrationCaller: AbstractRegistrationCaller{contract: contract}, AbstractRegistrationTransactor: AbstractRegistrationTransactor{contract: contract}, AbstractRegistrationFilterer: AbstractRegistrationFilterer{contract: contract}}, nil
}

// AbstractRegistration is an auto generated Go binding around an Ethereum contract.
type AbstractRegistration struct {
	AbstractRegistrationCaller     // Read-only binding to the contract
	AbstractRegistrationTransactor // Write-only binding to the contract
	AbstractRegistrationFilterer   // Log filterer for contract events
}

// AbstractRegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractRegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractRegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractRegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractRegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractRegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractRegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractRegistrationSession struct {
	Contract     *AbstractRegistration // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AbstractRegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractRegistrationCallerSession struct {
	Contract *AbstractRegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AbstractRegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractRegistrationTransactorSession struct {
	Contract     *AbstractRegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AbstractRegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractRegistrationRaw struct {
	Contract *AbstractRegistration // Generic contract binding to access the raw methods on
}

// AbstractRegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractRegistrationCallerRaw struct {
	Contract *AbstractRegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractRegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractRegistrationTransactorRaw struct {
	Contract *AbstractRegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractRegistration creates a new instance of AbstractRegistration, bound to a specific deployed contract.
func NewAbstractRegistration(address common.Address, backend bind.ContractBackend) (*AbstractRegistration, error) {
	contract, err := bindAbstractRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractRegistration{AbstractRegistrationCaller: AbstractRegistrationCaller{contract: contract}, AbstractRegistrationTransactor: AbstractRegistrationTransactor{contract: contract}, AbstractRegistrationFilterer: AbstractRegistrationFilterer{contract: contract}}, nil
}

// NewAbstractRegistrationCaller creates a new read-only instance of AbstractRegistration, bound to a specific deployed contract.
func NewAbstractRegistrationCaller(address common.Address, caller bind.ContractCaller) (*AbstractRegistrationCaller, error) {
	contract, err := bindAbstractRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractRegistrationCaller{contract: contract}, nil
}

// NewAbstractRegistrationTransactor creates a new write-only instance of AbstractRegistration, bound to a specific deployed contract.
func NewAbstractRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractRegistrationTransactor, error) {
	contract, err := bindAbstractRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractRegistrationTransactor{contract: contract}, nil
}

// NewAbstractRegistrationFilterer creates a new log filterer instance of AbstractRegistration, bound to a specific deployed contract.
func NewAbstractRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractRegistrationFilterer, error) {
	contract, err := bindAbstractRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractRegistrationFilterer{contract: contract}, nil
}

// bindAbstractRegistration binds a generic wrapper to an already deployed contract.
func bindAbstractRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractRegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractRegistration *AbstractRegistrationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractRegistration.Contract.AbstractRegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractRegistration *AbstractRegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractRegistration.Contract.AbstractRegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractRegistration *AbstractRegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractRegistration.Contract.AbstractRegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractRegistration *AbstractRegistrationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AbstractRegistration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractRegistration *AbstractRegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractRegistration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractRegistration *AbstractRegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractRegistration.Contract.contract.Transact(opts, method, params...)
}

// RESIGNREQUESTS is a free data retrieval call binding the contract method 0x500f99f7.
//
// Solidity: function RESIGN_REQUESTS( address) constant returns(uint256)
func (_AbstractRegistration *AbstractRegistrationCaller) RESIGNREQUESTS(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AbstractRegistration.contract.Call(opts, out, "RESIGN_REQUESTS", arg0)
	return *ret0, err
}

// RESIGNREQUESTS is a free data retrieval call binding the contract method 0x500f99f7.
//
// Solidity: function RESIGN_REQUESTS( address) constant returns(uint256)
func (_AbstractRegistration *AbstractRegistrationSession) RESIGNREQUESTS(arg0 common.Address) (*big.Int, error) {
	return _AbstractRegistration.Contract.RESIGNREQUESTS(&_AbstractRegistration.CallOpts, arg0)
}

// RESIGNREQUESTS is a free data retrieval call binding the contract method 0x500f99f7.
//
// Solidity: function RESIGN_REQUESTS( address) constant returns(uint256)
func (_AbstractRegistration *AbstractRegistrationCallerSession) RESIGNREQUESTS(arg0 common.Address) (*big.Int, error) {
	return _AbstractRegistration.Contract.RESIGNREQUESTS(&_AbstractRegistration.CallOpts, arg0)
}

// GetRelayerByCoinbase is a free data retrieval call binding the contract method 0x540105c7.
//
// Solidity: function getRelayerByCoinbase( address) constant returns(uint256, address, uint256, uint16, address[], address[])
func (_AbstractRegistration *AbstractRegistrationCaller) GetRelayerByCoinbase(opts *bind.CallOpts, arg0 common.Address) (*big.Int, common.Address, *big.Int, uint16, []common.Address, []common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(uint16)
		ret4 = new([]common.Address)
		ret5 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _AbstractRegistration.contract.Call(opts, out, "getRelayerByCoinbase", arg0)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRelayerByCoinbase is a free data retrieval call binding the contract method 0x540105c7.
//
// Solidity: function getRelayerByCoinbase( address) constant returns(uint256, address, uint256, uint16, address[], address[])
func (_AbstractRegistration *AbstractRegistrationSession) GetRelayerByCoinbase(arg0 common.Address) (*big.Int, common.Address, *big.Int, uint16, []common.Address, []common.Address, error) {
	return _AbstractRegistration.Contract.GetRelayerByCoinbase(&_AbstractRegistration.CallOpts, arg0)
}

// GetRelayerByCoinbase is a free data retrieval call binding the contract method 0x540105c7.
//
// Solidity: function getRelayerByCoinbase( address) constant returns(uint256, address, uint256, uint16, address[], address[])
func (_AbstractRegistration *AbstractRegistrationCallerSession) GetRelayerByCoinbase(arg0 common.Address) (*big.Int, common.Address, *big.Int, uint16, []common.Address, []common.Address, error) {
	return _AbstractRegistration.Contract.GetRelayerByCoinbase(&_AbstractRegistration.CallOpts, arg0)
}

// LendingABI is the input ABI used to generate the binding from.
const LendingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"COLLATERALS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"addTerm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"LENDINGRELAYER_LIST\",\"outputs\":[{\"name\":\"_tradeFee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BASES\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"depositRate\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"COLLATERAL_LIST\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"tradeFee\",\"type\":\"uint16\"},{\"name\":\"baseTokens\",\"type\":\"address[]\"},{\"name\":\"terms\",\"type\":\"uint256[]\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"getLendingRelayerByCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// LendingBin is the compiled bytecode used for deploying new contracts.
const LendingBin = `0x608060405234801561001057600080fd5b50604051602080610a7c833981016040525160058054600160a060020a031916600160a060020a03909216919091179055610a2c806100506000396000f3006080604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630811f05a81146100a85780630c655955146100dc5780630faf292c146100f65780636d1dc42a1461012e5780636d75b9ee14610146578063822507011461016a57806383e280d91461019d5780638406c079146101be578063f80908ae146101d3578063fe82470014610277575b600080fd5b3480156100b457600080fd5b506100c0600435610340565b60408051600160a060020a039092168252519081900360200190f35b3480156100e857600080fd5b506100f4600435610368565b005b34801561010257600080fd5b50610117600160a060020a036004351661039d565b6040805161ffff9092168252519081900360200190f35b34801561013a57600080fd5b506100c06004356103b3565b34801561015257600080fd5b506100f4600160a060020a03600435166024356103c1565b34801561017657600080fd5b5061018b600160a060020a036004351661042f565b60408051918252519081900360200190f35b3480156101a957600080fd5b506100f4600160a060020a0360043516610441565b3480156101ca57600080fd5b506100c06104a0565b3480156101df57600080fd5b5060408051602060046044358181013583810280860185019096528085526100f4958335600160a060020a0316956024803561ffff1696369695606495939492019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506104af9650505050505050565b34801561028357600080fd5b50610298600160a060020a0360043516610813565b604051808461ffff1661ffff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156102ea5781810151838201526020016102d2565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015610329578181015183820152602001610311565b505050509050019550505050505060405180910390f35b600280548290811061034e57fe5b600091825260209091200154600160a060020a0316905081565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155565b60006020819052908152604090205461ffff1681565b600380548290811061034e57fe5b600160a060020a03909116600081815260016020819052604082209390935560028054938401815590527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace909101805473ffffffffffffffffffffffffffffffffffffffff19169091179055565b60016020526000908152604090205481565b600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600554600160a060020a031681565b600554604080517f540105c7000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301529151600093929092169163540105c791602480820192869290919082900301818387803b15801561051957600080fd5b505af115801561052d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c081101561055657600080fd5b81516020830151604084015160608501516080860180519496939592949193928301929164010000000081111561058c57600080fd5b8201602081018481111561059f57600080fd5b81518560208202830111640100000000821117156105bc57600080fd5b505092919060200180516401000000008111156105d857600080fd5b820160208101848111156105eb57600080fd5b815185602082028301116401000000008211171561060857600080fd5b509799505050600160a060020a0388163314965061068e9550505050505057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6f776e6572207265717569726564000000000000000000000000000000000000604482015290519081900360640190fd5b600554604080517f500f99f7000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301529151919092169163500f99f79160248083019260209291908290030181600087803b1580156106f657600080fd5b505af115801561070a573d6000803e3d6000fd5b505050506040513d602081101561072057600080fd5b50511561078e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f72656c6179657220726571756972656420746f20636c6f736500000000000000604482015290519081900360640190fd5b6040805160608101825261ffff86811682526020808301878152838501879052600160a060020a038a166000908152808352949094208351815461ffff1916931692909217825592518051929391926107ed92600185019201906108f9565b506040820151805161080991600284019160209091019061096b565b5050505050505050565b600160a060020a0381166000908152602081815260408083208054600182018054845181870281018701909552808552606095869561ffff90941694929360029093019291849183018282801561089357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610875575b50505050509150808054806020026020016040519081016040528092919081815260200182805480156108e557602002820191906000526020600020905b8154815260200190600101908083116108d1575b505050505090509250925092509193909250565b82805482825590600052602060002090810192821561095b579160200282015b8281111561095b578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190610919565b506109679291506109b2565b5090565b8280548282559060005260206000209081019282156109a6579160200282015b828111156109a657825182559160200191906001019061098b565b506109679291506109e6565b6109e391905b8082111561096757805473ffffffffffffffffffffffffffffffffffffffff191681556001016109b8565b90565b6109e391905b8082111561096757600081556001016109ec5600a165627a7a7230582073dfe07decdbff6fe062e2d4dfbfc110ae0ea689f9966937682b049c4afd2c2d0029`

// DeployLending deploys a new Ethereum contract, binding an instance of Lending to it.
func DeployLending(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Lending, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LendingBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lending{LendingCaller: LendingCaller{contract: contract}, LendingTransactor: LendingTransactor{contract: contract}, LendingFilterer: LendingFilterer{contract: contract}}, nil
}

// Lending is an auto generated Go binding around an Ethereum contract.
type Lending struct {
	LendingCaller     // Read-only binding to the contract
	LendingTransactor // Write-only binding to the contract
	LendingFilterer   // Log filterer for contract events
}

// LendingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingSession struct {
	Contract     *Lending          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingCallerSession struct {
	Contract *LendingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LendingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingTransactorSession struct {
	Contract     *LendingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LendingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingRaw struct {
	Contract *Lending // Generic contract binding to access the raw methods on
}

// LendingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingCallerRaw struct {
	Contract *LendingCaller // Generic read-only contract binding to access the raw methods on
}

// LendingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingTransactorRaw struct {
	Contract *LendingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLending creates a new instance of Lending, bound to a specific deployed contract.
func NewLending(address common.Address, backend bind.ContractBackend) (*Lending, error) {
	contract, err := bindLending(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lending{LendingCaller: LendingCaller{contract: contract}, LendingTransactor: LendingTransactor{contract: contract}, LendingFilterer: LendingFilterer{contract: contract}}, nil
}

// NewLendingCaller creates a new read-only instance of Lending, bound to a specific deployed contract.
func NewLendingCaller(address common.Address, caller bind.ContractCaller) (*LendingCaller, error) {
	contract, err := bindLending(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingCaller{contract: contract}, nil
}

// NewLendingTransactor creates a new write-only instance of Lending, bound to a specific deployed contract.
func NewLendingTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingTransactor, error) {
	contract, err := bindLending(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingTransactor{contract: contract}, nil
}

// NewLendingFilterer creates a new log filterer instance of Lending, bound to a specific deployed contract.
func NewLendingFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingFilterer, error) {
	contract, err := bindLending(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingFilterer{contract: contract}, nil
}

// bindLending binds a generic wrapper to an already deployed contract.
func bindLending(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lending *LendingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lending.Contract.LendingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lending *LendingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.Contract.LendingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lending *LendingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lending.Contract.LendingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lending *LendingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lending.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lending *LendingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lending *LendingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lending.Contract.contract.Transact(opts, method, params...)
}

// BASES is a free data retrieval call binding the contract method 0x6d1dc42a.
//
// Solidity: function BASES( uint256) constant returns(address)
func (_Lending *LendingCaller) BASES(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "BASES", arg0)
	return *ret0, err
}

// BASES is a free data retrieval call binding the contract method 0x6d1dc42a.
//
// Solidity: function BASES( uint256) constant returns(address)
func (_Lending *LendingSession) BASES(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.BASES(&_Lending.CallOpts, arg0)
}

// BASES is a free data retrieval call binding the contract method 0x6d1dc42a.
//
// Solidity: function BASES( uint256) constant returns(address)
func (_Lending *LendingCallerSession) BASES(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.BASES(&_Lending.CallOpts, arg0)
}

// COLLATERALS is a free data retrieval call binding the contract method 0x0811f05a.
//
// Solidity: function COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingCaller) COLLATERALS(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "COLLATERALS", arg0)
	return *ret0, err
}

// COLLATERALS is a free data retrieval call binding the contract method 0x0811f05a.
//
// Solidity: function COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingSession) COLLATERALS(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.COLLATERALS(&_Lending.CallOpts, arg0)
}

// COLLATERALS is a free data retrieval call binding the contract method 0x0811f05a.
//
// Solidity: function COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingCallerSession) COLLATERALS(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.COLLATERALS(&_Lending.CallOpts, arg0)
}

// COLLATERALLIST is a free data retrieval call binding the contract method 0x82250701.
//
// Solidity: function COLLATERAL_LIST( address) constant returns(uint256)
func (_Lending *LendingCaller) COLLATERALLIST(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "COLLATERAL_LIST", arg0)
	return *ret0, err
}

// COLLATERALLIST is a free data retrieval call binding the contract method 0x82250701.
//
// Solidity: function COLLATERAL_LIST( address) constant returns(uint256)
func (_Lending *LendingSession) COLLATERALLIST(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.COLLATERALLIST(&_Lending.CallOpts, arg0)
}

// COLLATERALLIST is a free data retrieval call binding the contract method 0x82250701.
//
// Solidity: function COLLATERAL_LIST( address) constant returns(uint256)
func (_Lending *LendingCallerSession) COLLATERALLIST(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.COLLATERALLIST(&_Lending.CallOpts, arg0)
}

// LENDINGRELAYERLIST is a free data retrieval call binding the contract method 0x0faf292c.
//
// Solidity: function LENDINGRELAYER_LIST( address) constant returns(_tradeFee uint16)
func (_Lending *LendingCaller) LENDINGRELAYERLIST(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "LENDINGRELAYER_LIST", arg0)
	return *ret0, err
}

// LENDINGRELAYERLIST is a free data retrieval call binding the contract method 0x0faf292c.
//
// Solidity: function LENDINGRELAYER_LIST( address) constant returns(_tradeFee uint16)
func (_Lending *LendingSession) LENDINGRELAYERLIST(arg0 common.Address) (uint16, error) {
	return _Lending.Contract.LENDINGRELAYERLIST(&_Lending.CallOpts, arg0)
}

// LENDINGRELAYERLIST is a free data retrieval call binding the contract method 0x0faf292c.
//
// Solidity: function LENDINGRELAYER_LIST( address) constant returns(_tradeFee uint16)
func (_Lending *LendingCallerSession) LENDINGRELAYERLIST(arg0 common.Address) (uint16, error) {
	return _Lending.Contract.LENDINGRELAYERLIST(&_Lending.CallOpts, arg0)
}

// GetLendingRelayerByCoinbase is a free data retrieval call binding the contract method 0xfe824700.
//
// Solidity: function getLendingRelayerByCoinbase(coinbase address) constant returns(uint16, address[], uint256[])
func (_Lending *LendingCaller) GetLendingRelayerByCoinbase(opts *bind.CallOpts, coinbase common.Address) (uint16, []common.Address, []*big.Int, error) {
	var (
		ret0 = new(uint16)
		ret1 = new([]common.Address)
		ret2 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Lending.contract.Call(opts, out, "getLendingRelayerByCoinbase", coinbase)
	return *ret0, *ret1, *ret2, err
}

// GetLendingRelayerByCoinbase is a free data retrieval call binding the contract method 0xfe824700.
//
// Solidity: function getLendingRelayerByCoinbase(coinbase address) constant returns(uint16, address[], uint256[])
func (_Lending *LendingSession) GetLendingRelayerByCoinbase(coinbase common.Address) (uint16, []common.Address, []*big.Int, error) {
	return _Lending.Contract.GetLendingRelayerByCoinbase(&_Lending.CallOpts, coinbase)
}

// GetLendingRelayerByCoinbase is a free data retrieval call binding the contract method 0xfe824700.
//
// Solidity: function getLendingRelayerByCoinbase(coinbase address) constant returns(uint16, address[], uint256[])
func (_Lending *LendingCallerSession) GetLendingRelayerByCoinbase(coinbase common.Address) (uint16, []common.Address, []*big.Int, error) {
	return _Lending.Contract.GetLendingRelayerByCoinbase(&_Lending.CallOpts, coinbase)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_Lending *LendingCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "relayer")
	return *ret0, err
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_Lending *LendingSession) Relayer() (common.Address, error) {
	return _Lending.Contract.Relayer(&_Lending.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_Lending *LendingCallerSession) Relayer() (common.Address, error) {
	return _Lending.Contract.Relayer(&_Lending.CallOpts)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(token address) returns()
func (_Lending *LendingTransactor) AddBaseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addBaseToken", token)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(token address) returns()
func (_Lending *LendingSession) AddBaseToken(token common.Address) (*types.Transaction, error) {
	return _Lending.Contract.AddBaseToken(&_Lending.TransactOpts, token)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(token address) returns()
func (_Lending *LendingTransactorSession) AddBaseToken(token common.Address) (*types.Transaction, error) {
	return _Lending.Contract.AddBaseToken(&_Lending.TransactOpts, token)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6d75b9ee.
//
// Solidity: function addCollateral(token address, depositRate uint256) returns()
func (_Lending *LendingTransactor) AddCollateral(opts *bind.TransactOpts, token common.Address, depositRate *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addCollateral", token, depositRate)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6d75b9ee.
//
// Solidity: function addCollateral(token address, depositRate uint256) returns()
func (_Lending *LendingSession) AddCollateral(token common.Address, depositRate *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddCollateral(&_Lending.TransactOpts, token, depositRate)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6d75b9ee.
//
// Solidity: function addCollateral(token address, depositRate uint256) returns()
func (_Lending *LendingTransactorSession) AddCollateral(token common.Address, depositRate *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddCollateral(&_Lending.TransactOpts, token, depositRate)
}

// AddTerm is a paid mutator transaction binding the contract method 0x0c655955.
//
// Solidity: function addTerm(term uint256) returns()
func (_Lending *LendingTransactor) AddTerm(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addTerm", term)
}

// AddTerm is a paid mutator transaction binding the contract method 0x0c655955.
//
// Solidity: function addTerm(term uint256) returns()
func (_Lending *LendingSession) AddTerm(term *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddTerm(&_Lending.TransactOpts, term)
}

// AddTerm is a paid mutator transaction binding the contract method 0x0c655955.
//
// Solidity: function addTerm(term uint256) returns()
func (_Lending *LendingTransactorSession) AddTerm(term *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddTerm(&_Lending.TransactOpts, term)
}

// Register is a paid mutator transaction binding the contract method 0xf80908ae.
//
// Solidity: function register(coinbase address, tradeFee uint16, baseTokens address[], terms uint256[]) returns()
func (_Lending *LendingTransactor) Register(opts *bind.TransactOpts, coinbase common.Address, tradeFee uint16, baseTokens []common.Address, terms []*big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "register", coinbase, tradeFee, baseTokens, terms)
}

// Register is a paid mutator transaction binding the contract method 0xf80908ae.
//
// Solidity: function register(coinbase address, tradeFee uint16, baseTokens address[], terms uint256[]) returns()
func (_Lending *LendingSession) Register(coinbase common.Address, tradeFee uint16, baseTokens []common.Address, terms []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Register(&_Lending.TransactOpts, coinbase, tradeFee, baseTokens, terms)
}

// Register is a paid mutator transaction binding the contract method 0xf80908ae.
//
// Solidity: function register(coinbase address, tradeFee uint16, baseTokens address[], terms uint256[]) returns()
func (_Lending *LendingTransactorSession) Register(coinbase common.Address, tradeFee uint16, baseTokens []common.Address, terms []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Register(&_Lending.TransactOpts, coinbase, tradeFee, baseTokens, terms)
}
