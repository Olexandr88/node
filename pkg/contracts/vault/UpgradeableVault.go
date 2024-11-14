// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Asset         common.Address
	Amount        uint64
	RevertMessage []byte
}

// UpgradeableVaultMetaData contains all meta data concerning the UpgradeableVault contract.
var UpgradeableVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStrategyAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStrategyChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreasuryAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZRC20Address\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedeemExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawExceedsLimit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimed\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structRevertContext\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ContextDataRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PerformanceFeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"PerformanceFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulated\",\"type\":\"uint256\"}],\"name\":\"RewardsPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"start\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"end\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"RewardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newStrategyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newStrategyChainId\",\"type\":\"uint32\"}],\"name\":\"StrategyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidRewardPerToken\",\"type\":\"uint256\"}],\"name\":\"UserRewardsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perfFee\",\"type\":\"uint256\"}],\"name\":\"VaultInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedRewards\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accumulated\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"checkpoint\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPerfFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"perfFee_\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structMessageContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"revertMessage\",\"type\":\"bytes\"}],\"internalType\":\"structRevertContext\",\"name\":\"revertContext\",\"type\":\"tuple\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"start\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"end\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"rate\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPerToken\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accumulated\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newFeeRate\",\"type\":\"uint16\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"setRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"name\":\"setRewardsInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_strategyChainId\",\"type\":\"uint32\"}],\"name\":\"setStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newStrategyChainId\",\"type\":\"uint32\"}],\"name\":\"switchStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"updateTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5061005261005760201b60201c565b6101c1565b600061006761015b60201b60201c565b90508060000160089054906101000a900460ff16156100b2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80168160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101585767ffffffffffffffff8160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014f91906101a6565b60405180910390a15b50565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b600067ffffffffffffffff82169050919050565b6101a081610183565b82525050565b60006020820190506101bb6000830184610197565b92915050565b6080516181e16101ea60003960008181612bcf01528181612c240152612ddf01526181e16000f3fe6080604052600436106102ae5760003560e01c8063841e456111610175578063ba087652116100dc578063ef5cfb8c11610095578063f7c618c11161006f578063f7c618c114610b74578063faec96e114610b9f578063fc3c1dc314610bca578063fe68c81f14610bf7576102ae565b8063ef5cfb8c14610ae5578063ef8b30f714610b0e578063f2fde38b14610b4b576102ae565b8063ba08765214610977578063c63d75b6146109b4578063c6e6f592146109f1578063ce96cb7714610a2e578063d905777e14610a6b578063dd62ed3e14610aa8576102ae565b8063a9059cbb1161012e578063a9059cbb14610843578063aa290e6d14610880578063ad3cb1cc146108a9578063b3d7f6b9146108d4578063b460af9414610911578063b5be6f221461094e576102ae565b8063841e4561146107355780638aee81271461075e5780638da5cb5b1461078757806394bf804d146107b257806395d89b41146107ef578063a2cf74f71461081a576102ae565b80634cdad506116102195780636e553f65116101d25780636e553f65146106115780636ff1c9bc1461064e57806370641a361461067757806370a08231146106a3578063715018a6146106e057806373f273fc146106f7576102ae565b80634cdad506146105125780634f1ef2861461054f57806352d1902d1461056b5780635a5c1823146105965780635bcfd616146105bf578063660b9de0146105e8576102ae565b806318160ddd1161026b57806318160ddd146103ec57806323b872dd14610417578063313ce5671461045457806338d52e0f1461047f5780633b19e84a146104aa578063402d267d146104d5576102ae565b806301e1d114146102b357806306fdde03146102de57806307a2d13a1461030957806307da060314610346578063095ea7b3146103725780630a28a477146103af575b600080fd5b3480156102bf57600080fd5b506102c8610c20565b6040516102d59190616166565b60405180910390f35b3480156102ea57600080fd5b506102f3610d95565b6040516103009190616211565b60405180910390f35b34801561031557600080fd5b50610330600480360381019061032b9190616273565b610e36565b60405161033d9190616166565b60405180910390f35b34801561035257600080fd5b5061035b610e4a565b604051610369929190616300565b60405180910390f35b34801561037e57600080fd5b5061039960048036038101906103949190616355565b610e9c565b6040516103a691906163b0565b60405180910390f35b3480156103bb57600080fd5b506103d660048036038101906103d19190616273565b610ebf565b6040516103e39190616166565b60405180910390f35b3480156103f857600080fd5b50610401610ed3565b60405161040e9190616166565b60405180910390f35b34801561042357600080fd5b5061043e600480360381019061043991906163cb565b610eeb565b60405161044b91906163b0565b60405180910390f35b34801561046057600080fd5b50610469610f1a565b604051610476919061643a565b60405180910390f35b34801561048b57600080fd5b50610494610f51565b6040516104a19190616455565b60405180910390f35b3480156104b657600080fd5b506104bf610f89565b6040516104cc9190616455565b60405180910390f35b3480156104e157600080fd5b506104fc60048036038101906104f79190616470565b610fc1565b6040516105099190616166565b60405180910390f35b34801561051e57600080fd5b5061053960048036038101906105349190616273565b610feb565b6040516105469190616166565b60405180910390f35b610569600480360381019061056491906165d2565b610fff565b005b34801561057757600080fd5b5061058061101e565b60405161058d9190616647565b60405180910390f35b3480156105a257600080fd5b506105bd60048036038101906105b8919061677b565b611051565b005b3480156105cb57600080fd5b506105e660048036038101906105e191906168b2565b611314565b005b3480156105f457600080fd5b5061060f600480360381019061060a9190616975565b6113df565b005b34801561061d57600080fd5b50610638600480360381019061063391906169be565b611419565b6040516106459190616166565b60405180910390f35b34801561065a57600080fd5b5061067560048036038101906106709190616470565b61149b565b005b34801561068357600080fd5b5061068c611571565b60405161069a929190616a29565b60405180910390f35b3480156106af57600080fd5b506106ca60048036038101906106c59190616470565b6115af565b6040516106d79190616166565b60405180910390f35b3480156106ec57600080fd5b506106f5611606565b005b34801561070357600080fd5b5061071e60048036038101906107199190616470565b61161a565b60405161072c929190616a52565b60405180910390f35b34801561074157600080fd5b5061075c60048036038101906107579190616470565b611676565b005b34801561076a57600080fd5b5061078560048036038101906107809190616470565b611737565b005b34801561079357600080fd5b5061079c611782565b6040516107a99190616455565b60405180910390f35b3480156107be57600080fd5b506107d960048036038101906107d491906169be565b6117ba565b6040516107e69190616166565b60405180910390f35b3480156107fb57600080fd5b5061080461183c565b6040516108119190616211565b60405180910390f35b34801561082657600080fd5b50610841600480360381019061083c9190616a7b565b6118dd565b005b34801561084f57600080fd5b5061086a60048036038101906108659190616355565b611af0565b60405161087791906163b0565b60405180910390f35b34801561088c57600080fd5b506108a760048036038101906108a29190616ace565b611b13565b005b3480156108b557600080fd5b506108be611bc1565b6040516108cb9190616211565b60405180910390f35b3480156108e057600080fd5b506108fb60048036038101906108f69190616273565b611bfa565b6040516109089190616166565b60405180910390f35b34801561091d57600080fd5b5061093860048036038101906109339190616afb565b611c0e565b6040516109459190616166565b60405180910390f35b34801561095a57600080fd5b5061097560048036038101906109709190616b7a565b611c92565b005b34801561098357600080fd5b5061099e60048036038101906109999190616afb565b611e05565b6040516109ab9190616166565b60405180910390f35b3480156109c057600080fd5b506109db60048036038101906109d69190616470565b611e89565b6040516109e89190616166565b60405180910390f35b3480156109fd57600080fd5b50610a186004803603810190610a139190616273565b611eb3565b604051610a259190616166565b60405180910390f35b348015610a3a57600080fd5b50610a556004803603810190610a509190616470565b611ec7565b604051610a629190616166565b60405180910390f35b348015610a7757600080fd5b50610a926004803603810190610a8d9190616470565b611ee3565b604051610a9f9190616166565b60405180910390f35b348015610ab457600080fd5b50610acf6004803603810190610aca9190616bba565b611ef5565b604051610adc9190616166565b60405180910390f35b348015610af157600080fd5b50610b0c6004803603810190610b079190616470565b611f8a565b005b348015610b1a57600080fd5b50610b356004803603810190610b309190616273565b61214f565b604051610b429190616166565b60405180910390f35b348015610b5757600080fd5b50610b726004803603810190610b6d9190616470565b612163565b005b348015610b8057600080fd5b50610b896121e9565b604051610b969190616c59565b60405180910390f35b348015610bab57600080fd5b50610bb461220d565b604051610bc19190616c83565b60405180910390f35b348015610bd657600080fd5b50610bdf612233565b604051610bee93929190616cc5565b60405180910390f35b348015610c0357600080fd5b50610c1e6004803603810190610c199190616b7a565b612283565b005b600080610c2b6127d0565b90506000610c37610f51565b73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610c6f9190616455565b602060405180830381865afa158015610c8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb09190616d11565b905060008260000160149054906101000a900463ffffffff1663ffffffff164603610d80578260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634fb9bbba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d689190616d11565b90508082610d769190616d6d565b9350505050610d92565b8082610d8c9190616d6d565b93505050505b90565b60606000610da16127f8565b9050806003018054610db290616dd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610dde90616dd0565b8015610e2b5780601f10610e0057610100808354040283529160200191610e2b565b820191906000526020600020905b815481529060010190602001808311610e0e57829003601f168201915b505050505091505090565b6000610e43826000612820565b9050919050565b6000806000610e576127d0565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160149054906101000a900463ffffffff1692509250509091565b600080610ea761290f565b9050610eb4818585612917565b600191505092915050565b6000610ecc826001612929565b9050919050565b600080610ede6127f8565b9050806002015491505090565b600080610ef661290f565b9050610f03858285612a18565b610f0e858585612aac565b60019150509392505050565b600080610f25612ba0565b9050610f2f612bc8565b8160000160149054906101000a900460ff16610f4b9190616e01565b91505090565b600080610f5c612ba0565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b600080610f946127d0565b90508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9050919050565b6000610ff8826000612820565b9050919050565b611007612bcd565b61101082612cb3565b61101a8282612cbe565b5050565b6000611028612ddd565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b600061105b612e64565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff161480156110a95750825b9050600060018367ffffffffffffffff161480156110de575060003073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156110ec575080155b15611123576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156111735760018560000160086101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16036111d9576040517fcfe2ea6300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111e38a8a612e8c565b6111ec88612ea2565b6111f533612eb6565b6111fd612eca565b60006112076127d0565b9050878160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550868160010160146101000a81548161ffff021916908361ffff1602179055507f7203e25960677692487f12bc9315b5b0cc817a4cbc645714efcb9bcef4b74dca611294610f1a565b886040516112a3929190616e67565b60405180910390a15083156113085760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516112ff9190616edf565b60405180910390a15b50505050505050505050565b60008060008085858101906113299190616f38565b935093509350935060008703611348576113438484612ed4565b6113d4565b600183036113615761135c8488848461334d565b6113d3565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036113c7576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113d284888a613669565b5b5b505050505050505050565b7f35a9324413457251c1059312318f6f1cec6bd0da4105d01315f3151b1e3a2c768160405161140e9190617131565b60405180910390a150565b60008061142583610fc1565b905080841115611470578284826040517f79012fb200000000000000000000000000000000000000000000000000000000815260040161146793929190617153565b60405180910390fd5b600061147b8561214f565b905061149061148861290f565b858784613982565b809250505092915050565b6114a3613c30565b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016114de9190616455565b602060405180830381865afa1580156114fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061151f9190616d11565b90506000810361155b576040517fd0d04f6000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61156d82611567611782565b83613cb7565b5050565b60028060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a900463ffffffff16905082565b6000806115ba6127f8565b90508060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054915050919050565b61160e613c30565b6116186000613d36565b565b60036020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b61167e613c30565b60006116886127d0565b9050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036116f0576040517fcfe2ea6300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61173f613c30565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008061178d613e0d565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6000806117c683611e89565b905080841115611811578284826040517f284ff66700000000000000000000000000000000000000000000000000000000815260040161180893929190617153565b60405180910390fd5b600061181c85611bfa565b905061183161182961290f565b858388613982565b809250505092915050565b606060006118486127f8565b905080600401805461185990616dd0565b80601f016020809104026020016040519081016040528092919081815260200182805461188590616dd0565b80156118d25780601f106118a7576101008083540402835291602001916118d2565b820191906000526020600020905b8154815290600101906020018083116118b557829003601f168201915b505050505091505090565b6118e5613c30565b818310611927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161191e906171d6565b60405180910390fd5b600160000160009054906101000a900463ffffffff1663ffffffff1642108061196a5750600160000160049054906101000a900463ffffffff1663ffffffff1642115b6119a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119a090617242565b60405180910390fd5b6119b1613e35565b600083836119bf9190617262565b826119ca91906172c5565b905060405180606001604052808563ffffffff1681526020018463ffffffff168152602001826bffffffffffffffffffffffff16815250600160008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548163ffffffff021916908363ffffffff16021790555060408201518160000160086101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555090505083600260000160106101000a81548163ffffffff021916908363ffffffff1602179055507f95efd8a2a0aa591f48fd9673cf5d13c2150ca7a1fe1cbe438dd3f0ae47064663848483604051611ae2939291906172f6565b60405180910390a150505050565b600080611afb61290f565b9050611b08818585612aac565b600191505092915050565b611b1b613c30565b6000611b256127d0565b90506107d08261ffff161115611b67576040517f25ef66ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818160010160146101000a81548161ffff021916908361ffff1602179055507f9b49d0cd76012d9c67241c2f68f836efbaf50ea29901a250040671402d5263f582604051611bb5919061732d565b60405180910390a15050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6000611c07826001612820565b9050919050565b600080611c1a83611ec7565b905080851115611c65578285826040517ffe9cceec000000000000000000000000000000000000000000000000000000008152600401611c5c93929190617153565b60405180910390fd5b6000611c7086610ebf565b9050611c86611c7d61290f565b86868985614048565b80925050509392505050565b611c9a613c30565b6000611ca46127d0565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611d0c576040517f3408148f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008263ffffffff1603611d4c576040517f9060358a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818160000160146101000a81548163ffffffff021916908363ffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f3f1cc356d56ba3d612ed5b8fa9d919d35bfc89e02529efd6723156bffbd2404283604051611df89190617348565b60405180910390a2505050565b600080611e1183611ee3565b905080851115611e5c578285826040517fb94abeec000000000000000000000000000000000000000000000000000000008152600401611e5393929190617153565b60405180910390fd5b6000611e6786610feb565b9050611e7d611e7461290f565b8686848a614048565b80925050509392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9050919050565b6000611ec0826000612929565b9050919050565b6000611edc611ed5836115af565b6000612820565b9050919050565b6000611eee826115af565b9050919050565b600080611f006127f8565b90508060010160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205491505092915050565b611f93336145c9565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1690506000811161204b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612042906173af565b60405180910390fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550612110828260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16613cb79092919063ffffffff16565b7ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683338383604051612143939291906173cf565b60405180910390a15050565b600061215c826000612929565b9050919050565b61216b613c30565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036121dd5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016121d49190616455565b60405180910390fd5b6121e681613d36565b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806122186127d0565b90508060010160149054906101000a900461ffff1691505090565b60018060000160009054906101000a900463ffffffff16908060000160049054906101000a900463ffffffff16908060000160089054906101000a90046bffffffffffffffffffffffff16905083565b61228b613c30565b60006122956127d0565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036122fd576040517f3408148f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603612386576040517f3408148f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008263ffffffff16036123c6576040517f9060358a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008260000160149054906101000a900463ffffffff169050848360000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838360000160146101000a81548163ffffffff021916908363ffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff167f3f1cc356d56ba3d612ed5b8fa9d919d35bfc89e02529efd6723156bffbd24042856040516124b49190617348565b60405180910390a260008273ffffffffffffffffffffffffffffffffffffffff16634fb9bbba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612509573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061252d9190616d11565b905060008111156125c4578273ffffffffffffffffffffffffffffffffffffffff1663441a3e70826b033b2e3c9fd0803ce80000006040518363ffffffff1660e01b815260040161257f929190617441565b6020604051808303816000875af115801561259e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c29190616d11565b505b60006125ce610f51565b73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016126069190616455565b602060405180830381865afa158015612623573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126479190616d11565b905060008111156127c757600061265c610f51565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b38760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b81526004016126ba92919061746a565b6020604051808303816000875af11580156126d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126fd91906174bf565b905080612736576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632afcf480836040518263ffffffff1660e01b81526004016127939190616166565b600060405180830381600087803b1580156127ad57600080fd5b505af11580156127c1573d6000803e3d6000fd5b50505050505b50505050505050565b60007f1a0ee6983e121525fbe4b5f5f8fd996faa9a018f8e366b3f036f295ddafb46df905090565b60007f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00905090565b60008061282b610ed3565b0361283857829050612909565b60006128426127d0565b9050600061284e612bc8565b600a61285a919061761f565b612862610ed3565b61286c9190616d6d565b90506000600161287a610c20565b6128849190616d6d565b90508260020154612893610c20565b11156128ec576127108360010160149054906101000a900461ffff1661ffff1684600201546128c0610c20565b6128ca9190617262565b6128d4919061766a565b6128de91906172c5565b816128e99190617262565b90505b6129038183878961490a909392919063ffffffff16565b93505050505b92915050565b600033905090565b6129248383836001614959565b505050565b600080612934610ed3565b0361294157829050612a12565b600061294b6127d0565b90506000612957612bc8565b600a612963919061761f565b61296b610ed3565b6129759190616d6d565b905060006001612983610c20565b61298d9190616d6d565b9050826002015461299c610c20565b11156129f5576127108360010160149054906101000a900461ffff1661ffff1684600201546129c9610c20565b6129d39190617262565b6129dd919061766a565b6129e791906172c5565b816129f29190617262565b90505b612a0c8282878961490a909392919063ffffffff16565b93505050505b92915050565b6000612a248484611ef5565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114612aa65781811015612a96578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401612a8d93929190617153565b60405180910390fd5b612aa584848484036000614959565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603612b1e5760006040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401612b159190616455565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612b905760006040517fec442f05000000000000000000000000000000000000000000000000000000008152600401612b879190616455565b60405180910390fd5b612b9b838383614b3f565b505050565b60007f0773e532dfede91f04b12a73d3d2acd361424f41f76b4fb79f090161e36b4e00905090565b600090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480612c7a57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16612c61614bcb565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612cb1576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612cbb613c30565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612d2657506040513d601f19601f82011682018060405250810190612d2391906176d8565b60015b612d6757816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401612d5e9190616455565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b8114612dce57806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401612dc59190616647565b60405180910390fd5b612dd88383614c22565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614612e62576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b612e94614c95565b612e9e8282614cd5565b5050565b612eaa614c95565b612eb381614d12565b50565b612ebe614c95565b612ec781614da9565b50565b612ed2614c95565b565b6000612edf83611ec7565b905080821115612f2a578282826040517ffe9cceec000000000000000000000000000000000000000000000000000000008152600401612f2193929190617153565b60405180910390fd5b6000612f3583610ebf565b90506000612f416127d0565b90506000612f4f8686614e2f565b90508160000160149054906101000a900463ffffffff1663ffffffff1646036130775760006001612f7e610c20565b6b033b2e3c9fd0803ce80000008489612f979190616d6d565b612fa1919061766a565b612fab91906172c5565b612fb59190616d6d565b90508260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663441a3e7083886130039190616d6d565b836040518363ffffffff1660e01b8152600401613021929190617705565b6020604051808303816000875af1158015613040573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130649190616d11565b506130718787848761334d565b50613345565b6000732ca7d64a7efe2d62a725e2b35cf7230d6677ffee90508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3736c533f7fe93fae114d0954697069df33c9b74fd77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b81526004016130ff92919061746a565b6020604051808303816000875af115801561311e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061314291906174bf565b5060006301c9c380905060008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516020016131839190617776565b604051602081830303815290604052905060006040518060600160405280602981526020016181836029913980519060200120905060008a8a878a6040516020016131d19493929190617791565b6040516020818303038152906040529050600082826040516020016131f792919061786a565b604051602081830303815290604052905060006040518060a0016040528073c3e53f4d16ae77db1c982e75a937b9f60fe6369073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020013073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600e81526020017f726576657274206d65737361676500000000000000000000000000000000000081525081526020016301c9c3808152509050736c533f7fe93fae114d0954697069df33c9b74fd773ffffffffffffffffffffffffffffffffffffffff16631cb5ea75866132e8610f51565b858a866040518663ffffffff1660e01b815260040161330b9594939291906179a9565b600060405180830381600087803b15801561332557600080fd5b505af1158015613339573d6000803e3d6000fd5b50505050505050505050505b505050505050565b60006133576127d0565b905060008311156133e7578473ffffffffffffffffffffffffffffffffffffffff167fafdbf88d8a6e3bd4dbd770e775c7e168f61db74aea5e21852ac1061bf1cb625d846040516133a89190616166565b60405180910390a26133e66133bb610f51565b8260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685613cb7565b5b6133f1858361507a565b6133f9610f51565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b3736c533f7fe93fae114d0954697069df33c9b74fd7866040518363ffffffff1660e01b815260040161344792919061746a565b6020604051808303816000875af1158015613466573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061348a91906174bf565b5060008560405160200161349e9190617776565b604051602081830303815290604052905060006040518060a0016040528073c3e53f4d16ae77db1c982e75a937b9f60fe6369073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020013073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600e81526020017f726576657274206d65737361676500000000000000000000000000000000000081525081526020016301c9c3808152509050736c533f7fe93fae114d0954697069df33c9b74fd773ffffffffffffffffffffffffffffffffffffffff16637c0dcb5f8388613590610f51565b856040518563ffffffff1660e01b81526004016135b09493929190617a11565b600060405180830381600087803b1580156135ca57600080fd5b505af11580156135de573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db8988604051613658929190617705565b60405180910390a450505050505050565b600061367484610fc1565b9050808311156136bf578383826040517f79012fb20000000000000000000000000000000000000000000000000000000081526004016136b693929190617153565b60405180910390fd5b60006136ca8461214f565b905060006136d66127d0565b9050848160030160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546137299190616d6d565b92505081905550848160020160008282546137449190616d6d565b9250508190555061375586836150fc565b8060000160149054906101000a900463ffffffff1663ffffffff1646036138ef576000613780610f51565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886040518363ffffffff1660e01b81526004016137de92919061746a565b6020604051808303816000875af11580156137fd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061382191906174bf565b90508061385a576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632afcf480876040518263ffffffff1660e01b81526004016138b79190616166565b600060405180830381600087803b1580156138d157600080fd5b505af11580156138e5573d6000803e3d6000fd5b5050505050613912565b600061390585876138fe610f51565b600061517e565b905061391081615531565b505b8573ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d78785604051613972929190617705565b60405180910390a3505050505050565b600061398c6127d0565b9050828160030160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546139df9190616d6d565b92505081905550828160020160008282546139fa9190616d6d565b92505081905550613a14613a0c610f51565b8630866158fa565b613a1e84836150fc565b8060000160149054906101000a900463ffffffff1663ffffffff164603613bb8576000613a49610f51565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866040518363ffffffff1660e01b8152600401613aa792919061746a565b6020604051808303816000875af1158015613ac6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613aea91906174bf565b905080613b23576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632afcf480856040518263ffffffff1660e01b8152600401613b809190616166565b600060405180830381600087803b158015613b9a57600080fd5b505af1158015613bae573d6000803e3d6000fd5b5050505050613bc2565b613bc183615531565b5b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d78585604051613c21929190617705565b60405180910390a35050505050565b613c3861290f565b73ffffffffffffffffffffffffffffffffffffffff16613c56611782565b73ffffffffffffffffffffffffffffffffffffffff1614613cb557613c7961290f565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401613cac9190616455565b60405180910390fd5b565b613d31838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401613cea92919061746a565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061597c565b505050565b6000613d40613e0d565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b600160000160009054906101000a900463ffffffff1663ffffffff164210614046576000600160000160049054906101000a900463ffffffff1663ffffffff164210613ebc57600260000160109054906101000a900463ffffffff16600160000160049054906101000a900463ffffffff16613eb19190617a64565b63ffffffff16613ee4565b600260000160109054906101000a900463ffffffff1663ffffffff1642613ee39190617262565b5b90506000811480613efc57506000613efa610ed3565b145b15613f075750614046565b613f0f610ed3565b670de0b6b3a7640000600160000160089054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff1683613f4f919061766a565b613f59919061766a565b613f6391906172c5565b600260000160008282829054906101000a90046fffffffffffffffffffffffffffffffff16613f929190617a9c565b92506101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555042600260000160106101000a81548163ffffffff021916908363ffffffff1602179055507fe972555b20cae8150e291bb40efce3ef4e3ed6b6b2c39c029346600e95469d48600260000160009054906101000a90046fffffffffffffffffffffffffffffffff1660405161403c9190617b11565b60405180910390a1505b565b60006140526127d0565b90508373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161461409357614092848784612a18565b5b600061409f8585614e2f565b90508160000160149054906101000a900463ffffffff1663ffffffff1646036142f257600060016140ce610c20565b6b033b2e3c9fd0803ce800000084886140e79190616d6d565b6140f1919061766a565b6140fb91906172c5565b6141059190616d6d565b905060008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663441a3e7084886141559190616d6d565b846040518363ffffffff1660e01b8152600401614173929190617705565b6020604051808303816000875af1158015614192573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141b69190616d11565b90506000831115614246578673ffffffffffffffffffffffffffffffffffffffff167fafdbf88d8a6e3bd4dbd770e775c7e168f61db74aea5e21852ac1061bf1cb625d846040516142079190616166565b60405180910390a261424561421a610f51565b8560010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685613cb7565b5b614250878661507a565b61426261425b610f51565b8988613cb7565b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db86856142d49190617262565b896040516142e3929190617705565b60405180910390a450506145c0565b6000732ca7d64a7efe2d62a725e2b35cf7230d6677ffee90508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3736c533f7fe93fae114d0954697069df33c9b74fd77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b815260040161437a92919061746a565b6020604051808303816000875af1158015614399573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906143bd91906174bf565b5060006301c9c380905060008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516020016143fe9190617776565b604051602081830303815290604052905060006040518060600160405280602981526020016181836029913980519060200120905060008989878a60405160200161444c9493929190617791565b60405160208183030381529060405290506000828260405160200161447292919061786a565b604051602081830303815290604052905060006040518060a0016040528073c3e53f4d16ae77db1c982e75a937b9f60fe6369073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020013073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600e81526020017f726576657274206d65737361676500000000000000000000000000000000000081525081526020016301c9c3808152509050736c533f7fe93fae114d0954697069df33c9b74fd773ffffffffffffffffffffffffffffffffffffffff16631cb5ea7586614563610f51565b858a866040518663ffffffff1660e01b81526004016145869594939291906179a9565b600060405180830381600087803b1580156145a057600080fd5b505af11580156145b4573d6000803e3d6000fd5b50505050505050505050505b50505050505050565b6145d1613e35565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020016000820160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815250509050600260000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681602001516fffffffffffffffffffffffffffffffff16036147095750614907565b670de0b6b3a76400008160200151600260000160009054906101000a90046fffffffffffffffffffffffffffffffff166147439190617b2c565b6fffffffffffffffffffffffffffffffff1661475e846115af565b614768919061766a565b61477291906172c5565b816000018181516147839190617a9c565b9150906fffffffffffffffffffffffffffffffff1690816fffffffffffffffffffffffffffffffff1681525050600260000160009054906101000a90046fffffffffffffffffffffffffffffffff1681602001906fffffffffffffffffffffffffffffffff1690816fffffffffffffffffffffffffffffffff168152505080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff1602179055509050507f5b9aaf4cc5141c090a75f8b8a627863eba92df81f0c83c096350da9b79aedd0482826000015183602001516040516148fd93929190617b70565b60405180910390a1505b50565b600061493a61491883615a1e565b80156149355750600084806149305761492f617296565b5b868809115b615a4c565b614945868686615a58565b61494f9190616d6d565b9050949350505050565b60006149636127f8565b9050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16036149d75760006040517fe602df050000000000000000000000000000000000000000000000000000000081526004016149ce9190616455565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603614a495760006040517f94280d62000000000000000000000000000000000000000000000000000000008152600401614a409190616455565b60405180910390fd5b828160010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508115614b38578373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051614b2f9190616166565b60405180910390a35b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614614b7d57614b7c836145c9565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614614bbb57614bba826145c9565b5b614bc6838383615b46565b505050565b6000614bf97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b615d85565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b614c2b82615d8f565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a2600081511115614c8857614c828282615e5c565b50614c91565b614c90615ee0565b5b5050565b614c9d615f1d565b614cd3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b614cdd614c95565b6000614ce76127f8565b905082816003019081614cfa9190617d49565b5081816004019081614d0c9190617d49565b50505050565b614d1a614c95565b6000614d24612ba0565b9050600080614d3284615f3d565b9150915081614d42576012614d44565b805b8360000160146101000a81548160ff021916908360ff160217905550838360000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b614db1614c95565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603614e235760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401614e1a9190616455565b60405180910390fd5b614e2c81613d36565b50565b600080614e3a6127d0565b905060008160030160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000614e95614e90876115af565b610e36565b9050600080600084841115614ff3578484614eb09190617262565b91508560010160149054906101000a900461ffff16612710614ed29190617e1b565b61ffff168660010160149054906101000a900461ffff1661ffff1683614ef8919061766a565b614f0291906172c5565b9050838589614f11919061766a565b614f1b91906172c5565b925060008389614f2b9190617262565b90508660010160149054906101000a900461ffff16612710614f4d9190617e1b565b61ffff1683614f5c919061766a565b818860010160149054906101000a900461ffff1661ffff1685614f7f919061766a565b614f89919061766a565b614f9391906172c5565b9750838760030160008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254614fe69190617262565b9250508190555050615053565b87925060009650828660030160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461504b9190617262565b925050819055505b828660020160008282546150679190617262565b9250508190555050505050505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036150ec5760006040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016150e39190616455565b60405180910390fd5b6150f882600083614b3f565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361516e5760006040517fec442f050000000000000000000000000000000000000000000000000000000081526004016151659190616455565b60405180910390fd5b61517a60008383614b3f565b5050565b60006060600267ffffffffffffffff81111561519d5761519c6164a7565b5b6040519080825280602002602001820160405280156151cb5781602001602082028036833780820191505090505b50905085816000815181106151e3576151e2617e51565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050838160018151811061523257615231617e51565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600367ffffffffffffffff811115615287576152866164a7565b5b6040519080825280602002602001820160405280156152b55781602001602082028036833780820191505090505b50905085816000815181106152cd576152cc617e51565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050735f0b1a82749cb4e2278ec87f8bf6b618dc71a8bf816001815181106153305761532f617e51565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050838160028151811061537f5761537e617e51565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508573ffffffffffffffffffffffffffffffffffffffff1663095ea7b3732ca7d64a7efe2d62a725e2b35cf7230d6677ffee876040518363ffffffff1660e01b815260040161540892919061746a565b6020604051808303816000875af1158015615427573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061544b91906174bf565b506000732ca7d64a7efe2d62a725e2b35cf7230d6677ffee73ffffffffffffffffffffffffffffffffffffffff166338ed17398786853060c861ffff16426154939190616d6d565b6040518663ffffffff1660e01b81526004016154b3959493929190617f2f565b6000604051808303816000875af11580156154d2573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906154fb919061804c565b9050806001835161550c9190617262565b8151811061551d5761551c617e51565b5b602002602001015192505050949350505050565b600061553b6127d0565b9050600073236b0de675cc8f46ae186897fccefe3370c9eded90508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b3736c533f7fe93fae114d0954697069df33c9b74fd77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b81526004016155c592919061746a565b6020604051808303816000875af11580156155e4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061560891906174bf565b506000626acfc09050615619610f51565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146156e657615653610f51565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b3736c533f7fe93fae114d0954697069df33c9b74fd7866040518363ffffffff1660e01b81526004016156a192919061746a565b6020604051808303816000875af11580156156c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906156e491906174bf565b505b60008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405160200161571d9190617776565b604051602081830303815290604052905060006040518060400160405280600f81526020017f696e766573742875696e743235362900000000000000000000000000000000008152508051906020012090506000866040516020016157829190616166565b6040516020818303038152906040529050600082826040516020016157a892919061786a565b604051602081830303815290604052905060006040518060a0016040528073c3e53f4d16ae77db1c982e75a937b9f60fe6369073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020013073ffffffffffffffffffffffffffffffffffffffff1681526020016040518060400160405280600e81526020017f726576657274206d6573736167650000000000000000000000000000000000008152508152602001626acfc08152509050736c533f7fe93fae114d0954697069df33c9b74fd773ffffffffffffffffffffffffffffffffffffffff1663048ae42c868b615899610f51565b868b876040518763ffffffff1660e01b81526004016158bd96959493929190618095565b600060405180830381600087803b1580156158d757600080fd5b505af11580156158eb573d6000803e3d6000fd5b50505050505050505050505050565b615976848573ffffffffffffffffffffffffffffffffffffffff166323b872dd86868660405160240161592f939291906173cf565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061597c565b50505050565b600080602060008451602086016000885af18061599f576040513d6000823e3d81fd5b3d9250600051915050600082146159ba5760018114156159d6565b60008473ffffffffffffffffffffffffffffffffffffffff163b145b15615a1857836040517f5274afe7000000000000000000000000000000000000000000000000000000008152600401615a0f9190616455565b60405180910390fd5b50505050565b600060016002836003811115615a3757615a3661810b565b5b615a41919061813a565b60ff16149050919050565b60008115159050919050565b6000808385029050600080198587098281108382030391505060008103615a9357838281615a8957615a88617296565b5b0492505050615b3f565b808411615ab357615ab2615aad600086146012601161604d565b616067565b5b600084868809905082811182039150808303925060008560000386169050808604955080840493506001818260000304019050808302841793506000600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808502955050505050505b9392505050565b6000615b506127f8565b9050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603615ba65781816002016000828254615b9a9190616d6d565b92505081905550615c7f565b60008160000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015615c35578481846040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401615c2c93929190617153565b60405180910390fd5b8281038260000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603615cca57818160020160008282540392505081905550615d1a565b818160000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051615d779190616166565b60405180910390a350505050565b6000819050919050565b60008173ffffffffffffffffffffffffffffffffffffffff163b03615deb57806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401615de29190616455565b60405180910390fd5b80615e187f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b615d85565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051615e86919061816b565b600060405180830381855af49150503d8060008114615ec1576040519150601f19603f3d011682016040523d82523d6000602084013e615ec6565b606091505b5091509150615ed6858383616079565b9250505092915050565b6000341115615f1b576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000615f27612e64565b60000160089054906101000a900460ff16905090565b6000806000808473ffffffffffffffffffffffffffffffffffffffff1660405160240160405160208183030381529060405263313ce56760e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051615fb2919061816b565b600060405180830381855afa9150503d8060008114615fed576040519150601f19603f3d011682016040523d82523d6000602084013e615ff2565b606091505b509150915081801561600657506020815110155b1561603e576000818060200190518101906160219190616d11565b905060ff8016811161603c5760018194509450505050616048565b505b6000809350935050505b915091565b600061605884615a4c565b82841802821890509392505050565b634e487b71600052806020526024601cfd5b60608261608e5761608982616108565b616100565b600082511480156160b6575060008473ffffffffffffffffffffffffffffffffffffffff163b145b156160f857836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016160ef9190616455565b60405180910390fd5b819050616101565b5b9392505050565b60008151111561611b5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000819050919050565b6161608161614d565b82525050565b600060208201905061617b6000830184616157565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156161bb5780820151818401526020810190506161a0565b60008484015250505050565b6000601f19601f8301169050919050565b60006161e382616181565b6161ed818561618c565b93506161fd81856020860161619d565b616206816161c7565b840191505092915050565b6000602082019050818103600083015261622b81846161d8565b905092915050565b6000604051905090565b600080fd5b600080fd5b6162508161614d565b811461625b57600080fd5b50565b60008135905061626d81616247565b92915050565b6000602082840312156162895761628861623d565b5b60006162978482850161625e565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006162cb826162a0565b9050919050565b6162db816162c0565b82525050565b600063ffffffff82169050919050565b6162fa816162e1565b82525050565b600060408201905061631560008301856162d2565b61632260208301846162f1565b9392505050565b616332816162c0565b811461633d57600080fd5b50565b60008135905061634f81616329565b92915050565b6000806040838503121561636c5761636b61623d565b5b600061637a85828601616340565b925050602061638b8582860161625e565b9150509250929050565b60008115159050919050565b6163aa81616395565b82525050565b60006020820190506163c560008301846163a1565b92915050565b6000806000606084860312156163e4576163e361623d565b5b60006163f286828701616340565b935050602061640386828701616340565b92505060406164148682870161625e565b9150509250925092565b600060ff82169050919050565b6164348161641e565b82525050565b600060208201905061644f600083018461642b565b92915050565b600060208201905061646a60008301846162d2565b92915050565b6000602082840312156164865761648561623d565b5b600061649484828501616340565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6164df826161c7565b810181811067ffffffffffffffff821117156164fe576164fd6164a7565b5b80604052505050565b6000616511616233565b905061651d82826164d6565b919050565b600067ffffffffffffffff82111561653d5761653c6164a7565b5b616546826161c7565b9050602081019050919050565b82818337600083830152505050565b600061657561657084616522565b616507565b905082815260208101848484011115616591576165906164a2565b5b61659c848285616553565b509392505050565b600082601f8301126165b9576165b861649d565b5b81356165c9848260208601616562565b91505092915050565b600080604083850312156165e9576165e861623d565b5b60006165f785828601616340565b925050602083013567ffffffffffffffff81111561661857616617616242565b5b616624858286016165a4565b9150509250929050565b6000819050919050565b6166418161662e565b82525050565b600060208201905061665c6000830184616638565b92915050565b600067ffffffffffffffff82111561667d5761667c6164a7565b5b616686826161c7565b9050602081019050919050565b60006166a66166a184616662565b616507565b9050828152602081018484840111156166c2576166c16164a2565b5b6166cd848285616553565b509392505050565b600082601f8301126166ea576166e961649d565b5b81356166fa848260208601616693565b91505092915050565b600061670e826162c0565b9050919050565b61671e81616703565b811461672957600080fd5b50565b60008135905061673b81616715565b92915050565b600061ffff82169050919050565b61675881616741565b811461676357600080fd5b50565b6000813590506167758161674f565b92915050565b600080600080600060a086880312156167975761679661623d565b5b600086013567ffffffffffffffff8111156167b5576167b4616242565b5b6167c1888289016166d5565b955050602086013567ffffffffffffffff8111156167e2576167e1616242565b5b6167ee888289016166d5565b94505060406167ff8882890161672c565b935050606061681088828901616340565b925050608061682188828901616766565b9150509295509295909350565b600080fd5b6000606082840312156168495761684861682e565b5b81905092915050565b600080fd5b600080fd5b60008083601f8401126168725761687161649d565b5b8235905067ffffffffffffffff81111561688f5761688e616852565b5b6020830191508360018202830111156168ab576168aa616857565b5b9250929050565b6000806000806000608086880312156168ce576168cd61623d565b5b600086013567ffffffffffffffff8111156168ec576168eb616242565b5b6168f888828901616833565b955050602061690988828901616340565b945050604061691a8882890161625e565b935050606086013567ffffffffffffffff81111561693b5761693a616242565b5b6169478882890161685c565b92509250509295509295909350565b60006060828403121561696c5761696b61682e565b5b81905092915050565b60006020828403121561698b5761698a61623d565b5b600082013567ffffffffffffffff8111156169a9576169a8616242565b5b6169b584828501616956565b91505092915050565b600080604083850312156169d5576169d461623d565b5b60006169e38582860161625e565b92505060206169f485828601616340565b9150509250929050565b60006fffffffffffffffffffffffffffffffff82169050919050565b616a23816169fe565b82525050565b6000604082019050616a3e6000830185616a1a565b616a4b60208301846162f1565b9392505050565b6000604082019050616a676000830185616a1a565b616a746020830184616a1a565b9392505050565b600080600060608486031215616a9457616a9361623d565b5b6000616aa28682870161625e565b9350506020616ab38682870161625e565b9250506040616ac48682870161625e565b9150509250925092565b600060208284031215616ae457616ae361623d565b5b6000616af284828501616766565b91505092915050565b600080600060608486031215616b1457616b1361623d565b5b6000616b228682870161625e565b9350506020616b3386828701616340565b9250506040616b4486828701616340565b9150509250925092565b616b57816162e1565b8114616b6257600080fd5b50565b600081359050616b7481616b4e565b92915050565b60008060408385031215616b9157616b9061623d565b5b6000616b9f85828601616340565b9250506020616bb085828601616b65565b9150509250929050565b60008060408385031215616bd157616bd061623d565b5b6000616bdf85828601616340565b9250506020616bf085828601616340565b9150509250929050565b6000819050919050565b6000616c1f616c1a616c15846162a0565b616bfa565b6162a0565b9050919050565b6000616c3182616c04565b9050919050565b6000616c4382616c26565b9050919050565b616c5381616c38565b82525050565b6000602082019050616c6e6000830184616c4a565b92915050565b616c7d81616741565b82525050565b6000602082019050616c986000830184616c74565b92915050565b60006bffffffffffffffffffffffff82169050919050565b616cbf81616c9e565b82525050565b6000606082019050616cda60008301866162f1565b616ce760208301856162f1565b616cf46040830184616cb6565b949350505050565b600081519050616d0b81616247565b92915050565b600060208284031215616d2757616d2661623d565b5b6000616d3584828501616cfc565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000616d788261614d565b9150616d838361614d565b9250828201905080821115616d9b57616d9a616d3e565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680616de857607f821691505b602082108103616dfb57616dfa616da1565b5b50919050565b6000616e0c8261641e565b9150616e178361641e565b9250828201905060ff811115616e3057616e2f616d3e565b5b92915050565b6000616e51616e4c616e4784616741565b616bfa565b61614d565b9050919050565b616e6181616e36565b82525050565b6000604082019050616e7c600083018561642b565b616e896020830184616e58565b9392505050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000616ec9616ec4616ebf84616e90565b616bfa565b616e9a565b9050919050565b616ed981616eae565b82525050565b6000602082019050616ef46000830184616ed0565b92915050565b6000616f05826162a0565b9050919050565b616f1581616efa565b8114616f2057600080fd5b50565b600081359050616f3281616f0c565b92915050565b60008060008060808587031215616f5257616f5161623d565b5b6000616f6087828801616f23565b9450506020616f718782880161625e565b9350506040616f828782880161625e565b9250506060616f938782880161625e565b91505092959194509250565b6000616fae6020840184616340565b905092915050565b616fbf816162c0565b82525050565b616fce81616e9a565b8114616fd957600080fd5b50565b600081359050616feb81616fc5565b92915050565b60006170006020840184616fdc565b905092915050565b61701181616e9a565b82525050565b600080fd5b600080fd5b600080fd5b6000808335600160200384360303811261704357617042617021565b5b83810192508235915060208301925067ffffffffffffffff82111561706b5761706a617017565b5b6001820236038313156170815761708061701c565b5b509250929050565b600082825260208201905092915050565b60006170a68385617089565b93506170b3838584616553565b6170bc836161c7565b840190509392505050565b6000606083016170da6000840184616f9f565b6170e76000860182616fb6565b506170f56020840184616ff1565b6171026020860182617008565b506171106040840184617026565b858303604087015261712383828461709a565b925050508091505092915050565b6000602082019050818103600083015261714b81846170c7565b905092915050565b600060608201905061716860008301866162d2565b6171756020830185616157565b6171826040830184616157565b949350505050565b7f496e76616c696420696e74657276616c00000000000000000000000000000000600082015250565b60006171c060108361618c565b91506171cb8261718a565b602082019050919050565b600060208201905081810360008301526171ef816171b3565b9050919050565b7f52657761726473206f6e676f696e670000000000000000000000000000000000600082015250565b600061722c600f8361618c565b9150617237826171f6565b602082019050919050565b6000602082019050818103600083015261725b8161721f565b9050919050565b600061726d8261614d565b91506172788361614d565b92508282039050818111156172905761728f616d3e565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006172d08261614d565b91506172db8361614d565b9250826172eb576172ea617296565b5b828204905092915050565b600060608201905061730b60008301866162f1565b61731860208301856162f1565b6173256040830184616157565b949350505050565b60006020820190506173426000830184616e58565b92915050565b600060208201905061735d60008301846162f1565b92915050565b7f4e6f207265776172647320746f20636c61696d00000000000000000000000000600082015250565b600061739960138361618c565b91506173a482617363565b602082019050919050565b600060208201905081810360008301526173c88161738c565b9050919050565b60006060820190506173e460008301866162d2565b6173f160208301856162d2565b6173fe6040830184616157565b949350505050565b6000819050919050565b600061742b61742661742184617406565b616bfa565b61614d565b9050919050565b61743b81617410565b82525050565b60006040820190506174566000830185616157565b6174636020830184617432565b9392505050565b600060408201905061747f60008301856162d2565b61748c6020830184616157565b9392505050565b61749c81616395565b81146174a757600080fd5b50565b6000815190506174b981617493565b92915050565b6000602082840312156174d5576174d461623d565b5b60006174e3848285016174aa565b91505092915050565b60008160011c9050919050565b6000808291508390505b60018511156175435780860481111561751f5761751e616d3e565b5b600185161561752e5780820291505b808102905061753c856174ec565b9450617503565b94509492505050565b60008261755c5760019050617618565b8161756a5760009050617618565b8160018114617580576002811461758a576175b9565b6001915050617618565b60ff84111561759c5761759b616d3e565b5b8360020a9150848211156175b3576175b2616d3e565b5b50617618565b5060208310610133831016604e8410600b84101617156175ee5782820a9050838111156175e9576175e8616d3e565b5b617618565b6175fb84848460016174f9565b9250905081840481111561761257617611616d3e565b5b81810290505b9392505050565b600061762a8261614d565b91506176358361641e565b92506176627fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461754c565b905092915050565b60006176758261614d565b91506176808361614d565b925082820261768e8161614d565b915082820484148315176176a5576176a4616d3e565b5b5092915050565b6176b58161662e565b81146176c057600080fd5b50565b6000815190506176d2816176ac565b92915050565b6000602082840312156176ee576176ed61623d565b5b60006176fc848285016176c3565b91505092915050565b600060408201905061771a6000830185616157565b6177276020830184616157565b9392505050565b60008160601b9050919050565b60006177468261772e565b9050919050565b60006177588261773b565b9050919050565b61777061776b826162c0565b61774d565b82525050565b6000617782828461775f565b60148201915081905092915050565b60006080820190506177a660008301876162d2565b6177b36020830186616157565b6177c06040830185616157565b6177cd6060830184616157565b95945050505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b61781d617818826177d6565b617802565b82525050565b600081519050919050565b600081905092915050565b600061784482617823565b61784e818561782e565b935061785e81856020860161619d565b80840191505092915050565b6000617876828561780c565b6004820191506178868284617839565b91508190509392505050565b600082825260208201905092915050565b60006178ae82617823565b6178b88185617892565b93506178c881856020860161619d565b6178d1816161c7565b840191505092915050565b6178e581616395565b82525050565b60006178f682617823565b6179008185617089565b935061791081856020860161619d565b617919816161c7565b840191505092915050565b61792d8161614d565b82525050565b600060a08301600083015161794b6000860182616fb6565b50602083015161795e60208601826178dc565b5060408301516179716040860182616fb6565b506060830151848203606086015261798982826178eb565b915050608083015161799e6080860182617924565b508091505092915050565b600060a08201905081810360008301526179c381886178a3565b90506179d260208301876162d2565b81810360408301526179e481866178a3565b90506179f36060830185616157565b8181036080830152617a058184617933565b90509695505050505050565b60006080820190508181036000830152617a2b81876178a3565b9050617a3a6020830186616157565b617a4760408301856162d2565b8181036060830152617a598184617933565b905095945050505050565b6000617a6f826162e1565b9150617a7a836162e1565b9250828203905063ffffffff811115617a9657617a95616d3e565b5b92915050565b6000617aa7826169fe565b9150617ab2836169fe565b925082820190506fffffffffffffffffffffffffffffffff811115617ada57617ad9616d3e565b5b92915050565b6000617afb617af6617af1846169fe565b616bfa565b61614d565b9050919050565b617b0b81617ae0565b82525050565b6000602082019050617b266000830184617b02565b92915050565b6000617b37826169fe565b9150617b42836169fe565b925082820390506fffffffffffffffffffffffffffffffff811115617b6a57617b69616d3e565b5b92915050565b6000606082019050617b8560008301866162d2565b617b926020830185617b02565b617b9f6040830184617b02565b949350505050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302617c097fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82617bcc565b617c138683617bcc565b95508019841693508086168417925050509392505050565b6000617c46617c41617c3c8461614d565b616bfa565b61614d565b9050919050565b6000819050919050565b617c6083617c2b565b617c74617c6c82617c4d565b848454617bd9565b825550505050565b600090565b617c89617c7c565b617c94818484617c57565b505050565b5b81811015617cb857617cad600082617c81565b600181019050617c9a565b5050565b601f821115617cfd57617cce81617ba7565b617cd784617bbc565b81016020851015617ce6578190505b617cfa617cf285617bbc565b830182617c99565b50505b505050565b600082821c905092915050565b6000617d2060001984600802617d02565b1980831691505092915050565b6000617d398383617d0f565b9150826002028217905092915050565b617d5282616181565b67ffffffffffffffff811115617d6b57617d6a6164a7565b5b617d758254616dd0565b617d80828285617cbc565b600060209050601f831160018114617db35760008415617da1578287015190505b617dab8582617d2d565b865550617e13565b601f198416617dc186617ba7565b60005b82811015617de957848901518255600182019150602085019450602081019050617dc4565b86831015617e065784890151617e02601f891682617d0f565b8355505b6001600288020188555050505b505050505050565b6000617e2682616741565b9150617e3183616741565b9250828203905061ffff811115617e4b57617e4a616d3e565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000617eb88383616fb6565b60208301905092915050565b6000602082019050919050565b6000617edc82617e80565b617ee68185617e8b565b9350617ef183617e9c565b8060005b83811015617f22578151617f098882617eac565b9750617f1483617ec4565b925050600181019050617ef5565b5085935050505092915050565b600060a082019050617f446000830188616157565b617f516020830187616157565b8181036040830152617f638186617ed1565b9050617f7260608301856162d2565b617f7f6080830184616157565b9695505050505050565b600067ffffffffffffffff821115617fa457617fa36164a7565b5b602082029050602081019050919050565b6000617fc8617fc384617f89565b616507565b90508083825260208201905060208402830185811115617feb57617fea616857565b5b835b8181101561801457806180008882616cfc565b845260208401935050602081019050617fed565b5050509392505050565b600082601f8301126180335761803261649d565b5b8151618043848260208601617fb5565b91505092915050565b6000602082840312156180625761806161623d565b5b600082015167ffffffffffffffff8111156180805761807f616242565b5b61808c8482850161801e565b91505092915050565b600060c08201905081810360008301526180af81896178a3565b90506180be6020830188616157565b6180cb60408301876162d2565b81810360608301526180dd81866178a3565b90506180ec6080830185616157565b81810360a08301526180fe8184617933565b9050979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60006181458261641e565b91506181508361641e565b9250826181605761815f617296565b5b828206905092915050565b60006181778284617839565b91508190509291505056fe776974686472617728616464726573732c75696e743235362c75696e743235362c75696e7432353629a2646970667358221220c1e5c6f06556aa9ef9cc41284282e71f64627c64cfdb2d0c6ec2e625231239c664736f6c634300081a0033",
}

// UpgradeableVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeableVaultMetaData.ABI instead.
var UpgradeableVaultABI = UpgradeableVaultMetaData.ABI

// UpgradeableVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeableVaultMetaData.Bin instead.
var UpgradeableVaultBin = UpgradeableVaultMetaData.Bin

// DeployUpgradeableVault deploys a new Ethereum contract, binding an instance of UpgradeableVault to it.
func DeployUpgradeableVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UpgradeableVault, error) {
	parsed, err := UpgradeableVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeableVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeableVault{UpgradeableVaultCaller: UpgradeableVaultCaller{contract: contract}, UpgradeableVaultTransactor: UpgradeableVaultTransactor{contract: contract}, UpgradeableVaultFilterer: UpgradeableVaultFilterer{contract: contract}}, nil
}

// UpgradeableVault is an auto generated Go binding around an Ethereum contract.
type UpgradeableVault struct {
	UpgradeableVaultCaller     // Read-only binding to the contract
	UpgradeableVaultTransactor // Write-only binding to the contract
	UpgradeableVaultFilterer   // Log filterer for contract events
}

// UpgradeableVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableVaultSession struct {
	Contract     *UpgradeableVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeableVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableVaultCallerSession struct {
	Contract *UpgradeableVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UpgradeableVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableVaultTransactorSession struct {
	Contract     *UpgradeableVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UpgradeableVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableVaultRaw struct {
	Contract *UpgradeableVault // Generic contract binding to access the raw methods on
}

// UpgradeableVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableVaultCallerRaw struct {
	Contract *UpgradeableVaultCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableVaultTransactorRaw struct {
	Contract *UpgradeableVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableVault creates a new instance of UpgradeableVault, bound to a specific deployed contract.
func NewUpgradeableVault(address common.Address, backend bind.ContractBackend) (*UpgradeableVault, error) {
	contract, err := bindUpgradeableVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVault{UpgradeableVaultCaller: UpgradeableVaultCaller{contract: contract}, UpgradeableVaultTransactor: UpgradeableVaultTransactor{contract: contract}, UpgradeableVaultFilterer: UpgradeableVaultFilterer{contract: contract}}, nil
}

// NewUpgradeableVaultCaller creates a new read-only instance of UpgradeableVault, bound to a specific deployed contract.
func NewUpgradeableVaultCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableVaultCaller, error) {
	contract, err := bindUpgradeableVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultCaller{contract: contract}, nil
}

// NewUpgradeableVaultTransactor creates a new write-only instance of UpgradeableVault, bound to a specific deployed contract.
func NewUpgradeableVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableVaultTransactor, error) {
	contract, err := bindUpgradeableVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultTransactor{contract: contract}, nil
}

// NewUpgradeableVaultFilterer creates a new log filterer instance of UpgradeableVault, bound to a specific deployed contract.
func NewUpgradeableVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableVaultFilterer, error) {
	contract, err := bindUpgradeableVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultFilterer{contract: contract}, nil
}

// bindUpgradeableVault binds a generic wrapper to an already deployed contract.
func bindUpgradeableVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeableVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableVault *UpgradeableVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableVault.Contract.UpgradeableVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableVault *UpgradeableVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.UpgradeableVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableVault *UpgradeableVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.UpgradeableVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableVault *UpgradeableVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableVault *UpgradeableVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableVault *UpgradeableVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_UpgradeableVault *UpgradeableVaultCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_UpgradeableVault *UpgradeableVaultSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _UpgradeableVault.Contract.UPGRADEINTERFACEVERSION(&_UpgradeableVault.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_UpgradeableVault *UpgradeableVaultCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _UpgradeableVault.Contract.UPGRADEINTERFACEVERSION(&_UpgradeableVault.CallOpts)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address ) view returns(uint128 accumulated, uint128 checkpoint)
func (_UpgradeableVault *UpgradeableVaultCaller) AccumulatedRewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	Accumulated *big.Int
	Checkpoint  *big.Int
}, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "accumulatedRewards", arg0)

	outstruct := new(struct {
		Accumulated *big.Int
		Checkpoint  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accumulated = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Checkpoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address ) view returns(uint128 accumulated, uint128 checkpoint)
func (_UpgradeableVault *UpgradeableVaultSession) AccumulatedRewards(arg0 common.Address) (struct {
	Accumulated *big.Int
	Checkpoint  *big.Int
}, error) {
	return _UpgradeableVault.Contract.AccumulatedRewards(&_UpgradeableVault.CallOpts, arg0)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address ) view returns(uint128 accumulated, uint128 checkpoint)
func (_UpgradeableVault *UpgradeableVaultCallerSession) AccumulatedRewards(arg0 common.Address) (struct {
	Accumulated *big.Int
	Checkpoint  *big.Int
}, error) {
	return _UpgradeableVault.Contract.AccumulatedRewards(&_UpgradeableVault.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.Allowance(&_UpgradeableVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.Allowance(&_UpgradeableVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_UpgradeableVault *UpgradeableVaultSession) Asset() (common.Address, error) {
	return _UpgradeableVault.Contract.Asset(&_UpgradeableVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCallerSession) Asset() (common.Address, error) {
	return _UpgradeableVault.Contract.Asset(&_UpgradeableVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.BalanceOf(&_UpgradeableVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.BalanceOf(&_UpgradeableVault.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.ConvertToAssets(&_UpgradeableVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.ConvertToAssets(&_UpgradeableVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.ConvertToShares(&_UpgradeableVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.ConvertToShares(&_UpgradeableVault.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UpgradeableVault *UpgradeableVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UpgradeableVault *UpgradeableVaultSession) Decimals() (uint8, error) {
	return _UpgradeableVault.Contract.Decimals(&_UpgradeableVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UpgradeableVault *UpgradeableVaultCallerSession) Decimals() (uint8, error) {
	return _UpgradeableVault.Contract.Decimals(&_UpgradeableVault.CallOpts)
}

// GetPerfFee is a free data retrieval call binding the contract method 0xfaec96e1.
//
// Solidity: function getPerfFee() view returns(uint16)
func (_UpgradeableVault *UpgradeableVaultCaller) GetPerfFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "getPerfFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetPerfFee is a free data retrieval call binding the contract method 0xfaec96e1.
//
// Solidity: function getPerfFee() view returns(uint16)
func (_UpgradeableVault *UpgradeableVaultSession) GetPerfFee() (uint16, error) {
	return _UpgradeableVault.Contract.GetPerfFee(&_UpgradeableVault.CallOpts)
}

// GetPerfFee is a free data retrieval call binding the contract method 0xfaec96e1.
//
// Solidity: function getPerfFee() view returns(uint16)
func (_UpgradeableVault *UpgradeableVaultCallerSession) GetPerfFee() (uint16, error) {
	return _UpgradeableVault.Contract.GetPerfFee(&_UpgradeableVault.CallOpts)
}

// GetStrategy is a free data retrieval call binding the contract method 0x07da0603.
//
// Solidity: function getStrategy() view returns(address, uint32)
func (_UpgradeableVault *UpgradeableVaultCaller) GetStrategy(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "getStrategy")

	if err != nil {
		return *new(common.Address), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetStrategy is a free data retrieval call binding the contract method 0x07da0603.
//
// Solidity: function getStrategy() view returns(address, uint32)
func (_UpgradeableVault *UpgradeableVaultSession) GetStrategy() (common.Address, uint32, error) {
	return _UpgradeableVault.Contract.GetStrategy(&_UpgradeableVault.CallOpts)
}

// GetStrategy is a free data retrieval call binding the contract method 0x07da0603.
//
// Solidity: function getStrategy() view returns(address, uint32)
func (_UpgradeableVault *UpgradeableVaultCallerSession) GetStrategy() (common.Address, uint32, error) {
	return _UpgradeableVault.Contract.GetStrategy(&_UpgradeableVault.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_UpgradeableVault *UpgradeableVaultSession) GetTreasury() (common.Address, error) {
	return _UpgradeableVault.Contract.GetTreasury(&_UpgradeableVault.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCallerSession) GetTreasury() (common.Address, error) {
	return _UpgradeableVault.Contract.GetTreasury(&_UpgradeableVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxDeposit(&_UpgradeableVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxDeposit(&_UpgradeableVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxMint(&_UpgradeableVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxMint(&_UpgradeableVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxRedeem(&_UpgradeableVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxRedeem(&_UpgradeableVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxWithdraw(&_UpgradeableVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _UpgradeableVault.Contract.MaxWithdraw(&_UpgradeableVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UpgradeableVault *UpgradeableVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UpgradeableVault *UpgradeableVaultSession) Name() (string, error) {
	return _UpgradeableVault.Contract.Name(&_UpgradeableVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UpgradeableVault *UpgradeableVaultCallerSession) Name() (string, error) {
	return _UpgradeableVault.Contract.Name(&_UpgradeableVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableVault *UpgradeableVaultSession) Owner() (common.Address, error) {
	return _UpgradeableVault.Contract.Owner(&_UpgradeableVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCallerSession) Owner() (common.Address, error) {
	return _UpgradeableVault.Contract.Owner(&_UpgradeableVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewDeposit(&_UpgradeableVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewDeposit(&_UpgradeableVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewMint(&_UpgradeableVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewMint(&_UpgradeableVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewRedeem(&_UpgradeableVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewRedeem(&_UpgradeableVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewWithdraw(&_UpgradeableVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _UpgradeableVault.Contract.PreviewWithdraw(&_UpgradeableVault.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UpgradeableVault *UpgradeableVaultCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UpgradeableVault *UpgradeableVaultSession) ProxiableUUID() ([32]byte, error) {
	return _UpgradeableVault.Contract.ProxiableUUID(&_UpgradeableVault.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UpgradeableVault *UpgradeableVaultCallerSession) ProxiableUUID() ([32]byte, error) {
	return _UpgradeableVault.Contract.ProxiableUUID(&_UpgradeableVault.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_UpgradeableVault *UpgradeableVaultSession) RewardToken() (common.Address, error) {
	return _UpgradeableVault.Contract.RewardToken(&_UpgradeableVault.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_UpgradeableVault *UpgradeableVaultCallerSession) RewardToken() (common.Address, error) {
	return _UpgradeableVault.Contract.RewardToken(&_UpgradeableVault.CallOpts)
}

// RewardsInterval is a free data retrieval call binding the contract method 0xfc3c1dc3.
//
// Solidity: function rewardsInterval() view returns(uint32 start, uint32 end, uint96 rate)
func (_UpgradeableVault *UpgradeableVaultCaller) RewardsInterval(opts *bind.CallOpts) (struct {
	Start uint32
	End   uint32
	Rate  *big.Int
}, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "rewardsInterval")

	outstruct := new(struct {
		Start uint32
		End   uint32
		Rate  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.End = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Rate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardsInterval is a free data retrieval call binding the contract method 0xfc3c1dc3.
//
// Solidity: function rewardsInterval() view returns(uint32 start, uint32 end, uint96 rate)
func (_UpgradeableVault *UpgradeableVaultSession) RewardsInterval() (struct {
	Start uint32
	End   uint32
	Rate  *big.Int
}, error) {
	return _UpgradeableVault.Contract.RewardsInterval(&_UpgradeableVault.CallOpts)
}

// RewardsInterval is a free data retrieval call binding the contract method 0xfc3c1dc3.
//
// Solidity: function rewardsInterval() view returns(uint32 start, uint32 end, uint96 rate)
func (_UpgradeableVault *UpgradeableVaultCallerSession) RewardsInterval() (struct {
	Start uint32
	End   uint32
	Rate  *big.Int
}, error) {
	return _UpgradeableVault.Contract.RewardsInterval(&_UpgradeableVault.CallOpts)
}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint128 accumulated, uint32 lastUpdated)
func (_UpgradeableVault *UpgradeableVaultCaller) RewardsPerToken(opts *bind.CallOpts) (struct {
	Accumulated *big.Int
	LastUpdated uint32
}, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "rewardsPerToken")

	outstruct := new(struct {
		Accumulated *big.Int
		LastUpdated uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Accumulated = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdated = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint128 accumulated, uint32 lastUpdated)
func (_UpgradeableVault *UpgradeableVaultSession) RewardsPerToken() (struct {
	Accumulated *big.Int
	LastUpdated uint32
}, error) {
	return _UpgradeableVault.Contract.RewardsPerToken(&_UpgradeableVault.CallOpts)
}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint128 accumulated, uint32 lastUpdated)
func (_UpgradeableVault *UpgradeableVaultCallerSession) RewardsPerToken() (struct {
	Accumulated *big.Int
	LastUpdated uint32
}, error) {
	return _UpgradeableVault.Contract.RewardsPerToken(&_UpgradeableVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UpgradeableVault *UpgradeableVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UpgradeableVault *UpgradeableVaultSession) Symbol() (string, error) {
	return _UpgradeableVault.Contract.Symbol(&_UpgradeableVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UpgradeableVault *UpgradeableVaultCallerSession) Symbol() (string, error) {
	return _UpgradeableVault.Contract.Symbol(&_UpgradeableVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) TotalAssets() (*big.Int, error) {
	return _UpgradeableVault.Contract.TotalAssets(&_UpgradeableVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _UpgradeableVault.Contract.TotalAssets(&_UpgradeableVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) TotalSupply() (*big.Int, error) {
	return _UpgradeableVault.Contract.TotalSupply(&_UpgradeableVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UpgradeableVault *UpgradeableVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _UpgradeableVault.Contract.TotalSupply(&_UpgradeableVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Approve(&_UpgradeableVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Approve(&_UpgradeableVault.TransactOpts, spender, value)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address to) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "claimRewards", to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address to) returns()
func (_UpgradeableVault *UpgradeableVaultSession) ClaimRewards(to common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.ClaimRewards(&_UpgradeableVault.TransactOpts, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address to) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) ClaimRewards(to common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.ClaimRewards(&_UpgradeableVault.TransactOpts, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Deposit(&_UpgradeableVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Deposit(&_UpgradeableVault.TransactOpts, assets, receiver)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address _token) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "emergencyWithdraw", _token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address _token) returns()
func (_UpgradeableVault *UpgradeableVaultSession) EmergencyWithdraw(_token common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.EmergencyWithdraw(&_UpgradeableVault.TransactOpts, _token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address _token) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) EmergencyWithdraw(_token common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.EmergencyWithdraw(&_UpgradeableVault.TransactOpts, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x5a5c1823.
//
// Solidity: function initialize(string name_, string symbol_, address asset_, address treasury_, uint16 perfFee_) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, asset_ common.Address, treasury_ common.Address, perfFee_ uint16) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "initialize", name_, symbol_, asset_, treasury_, perfFee_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5a5c1823.
//
// Solidity: function initialize(string name_, string symbol_, address asset_, address treasury_, uint16 perfFee_) returns()
func (_UpgradeableVault *UpgradeableVaultSession) Initialize(name_ string, symbol_ string, asset_ common.Address, treasury_ common.Address, perfFee_ uint16) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Initialize(&_UpgradeableVault.TransactOpts, name_, symbol_, asset_, treasury_, perfFee_)
}

// Initialize is a paid mutator transaction binding the contract method 0x5a5c1823.
//
// Solidity: function initialize(string name_, string symbol_, address asset_, address treasury_, uint16 perfFee_) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Initialize(name_ string, symbol_ string, asset_ common.Address, treasury_ common.Address, perfFee_ uint16) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Initialize(&_UpgradeableVault.TransactOpts, name_, symbol_, asset_, treasury_, perfFee_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Mint(&_UpgradeableVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Mint(&_UpgradeableVault.TransactOpts, shares, receiver)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) OnCall(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "onCall", context, zrc20, amount, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_UpgradeableVault *UpgradeableVaultSession) OnCall(context MessageContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.OnCall(&_UpgradeableVault.TransactOpts, context, zrc20, amount, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) OnCall(context MessageContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.OnCall(&_UpgradeableVault.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x660b9de0.
//
// Solidity: function onRevert((address,uint64,bytes) revertContext) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) OnRevert(opts *bind.TransactOpts, revertContext RevertContext) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "onRevert", revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0x660b9de0.
//
// Solidity: function onRevert((address,uint64,bytes) revertContext) returns()
func (_UpgradeableVault *UpgradeableVaultSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.OnRevert(&_UpgradeableVault.TransactOpts, revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0x660b9de0.
//
// Solidity: function onRevert((address,uint64,bytes) revertContext) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.OnRevert(&_UpgradeableVault.TransactOpts, revertContext)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Redeem(&_UpgradeableVault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Redeem(&_UpgradeableVault.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableVault *UpgradeableVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradeableVault.Contract.RenounceOwnership(&_UpgradeableVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradeableVault.Contract.RenounceOwnership(&_UpgradeableVault.TransactOpts)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 newFeeRate) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) SetPerformanceFee(opts *bind.TransactOpts, newFeeRate uint16) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "setPerformanceFee", newFeeRate)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 newFeeRate) returns()
func (_UpgradeableVault *UpgradeableVaultSession) SetPerformanceFee(newFeeRate uint16) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetPerformanceFee(&_UpgradeableVault.TransactOpts, newFeeRate)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 newFeeRate) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) SetPerformanceFee(newFeeRate uint16) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetPerformanceFee(&_UpgradeableVault.TransactOpts, newFeeRate)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _rewardToken) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) SetRewardToken(opts *bind.TransactOpts, _rewardToken common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "setRewardToken", _rewardToken)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _rewardToken) returns()
func (_UpgradeableVault *UpgradeableVaultSession) SetRewardToken(_rewardToken common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetRewardToken(&_UpgradeableVault.TransactOpts, _rewardToken)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _rewardToken) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) SetRewardToken(_rewardToken common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetRewardToken(&_UpgradeableVault.TransactOpts, _rewardToken)
}

// SetRewardsInterval is a paid mutator transaction binding the contract method 0xa2cf74f7.
//
// Solidity: function setRewardsInterval(uint256 start, uint256 end, uint256 totalRewards) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) SetRewardsInterval(opts *bind.TransactOpts, start *big.Int, end *big.Int, totalRewards *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "setRewardsInterval", start, end, totalRewards)
}

// SetRewardsInterval is a paid mutator transaction binding the contract method 0xa2cf74f7.
//
// Solidity: function setRewardsInterval(uint256 start, uint256 end, uint256 totalRewards) returns()
func (_UpgradeableVault *UpgradeableVaultSession) SetRewardsInterval(start *big.Int, end *big.Int, totalRewards *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetRewardsInterval(&_UpgradeableVault.TransactOpts, start, end, totalRewards)
}

// SetRewardsInterval is a paid mutator transaction binding the contract method 0xa2cf74f7.
//
// Solidity: function setRewardsInterval(uint256 start, uint256 end, uint256 totalRewards) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) SetRewardsInterval(start *big.Int, end *big.Int, totalRewards *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetRewardsInterval(&_UpgradeableVault.TransactOpts, start, end, totalRewards)
}

// SetStrategy is a paid mutator transaction binding the contract method 0xb5be6f22.
//
// Solidity: function setStrategy(address _strategyAddress, uint32 _strategyChainId) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) SetStrategy(opts *bind.TransactOpts, _strategyAddress common.Address, _strategyChainId uint32) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "setStrategy", _strategyAddress, _strategyChainId)
}

// SetStrategy is a paid mutator transaction binding the contract method 0xb5be6f22.
//
// Solidity: function setStrategy(address _strategyAddress, uint32 _strategyChainId) returns()
func (_UpgradeableVault *UpgradeableVaultSession) SetStrategy(_strategyAddress common.Address, _strategyChainId uint32) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetStrategy(&_UpgradeableVault.TransactOpts, _strategyAddress, _strategyChainId)
}

// SetStrategy is a paid mutator transaction binding the contract method 0xb5be6f22.
//
// Solidity: function setStrategy(address _strategyAddress, uint32 _strategyChainId) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) SetStrategy(_strategyAddress common.Address, _strategyChainId uint32) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SetStrategy(&_UpgradeableVault.TransactOpts, _strategyAddress, _strategyChainId)
}

// SwitchStrategy is a paid mutator transaction binding the contract method 0xfe68c81f.
//
// Solidity: function switchStrategy(address newStrategyAddress, uint32 newStrategyChainId) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) SwitchStrategy(opts *bind.TransactOpts, newStrategyAddress common.Address, newStrategyChainId uint32) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "switchStrategy", newStrategyAddress, newStrategyChainId)
}

// SwitchStrategy is a paid mutator transaction binding the contract method 0xfe68c81f.
//
// Solidity: function switchStrategy(address newStrategyAddress, uint32 newStrategyChainId) returns()
func (_UpgradeableVault *UpgradeableVaultSession) SwitchStrategy(newStrategyAddress common.Address, newStrategyChainId uint32) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SwitchStrategy(&_UpgradeableVault.TransactOpts, newStrategyAddress, newStrategyChainId)
}

// SwitchStrategy is a paid mutator transaction binding the contract method 0xfe68c81f.
//
// Solidity: function switchStrategy(address newStrategyAddress, uint32 newStrategyChainId) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) SwitchStrategy(newStrategyAddress common.Address, newStrategyChainId uint32) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.SwitchStrategy(&_UpgradeableVault.TransactOpts, newStrategyAddress, newStrategyChainId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Transfer(&_UpgradeableVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Transfer(&_UpgradeableVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.TransferFrom(&_UpgradeableVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.TransferFrom(&_UpgradeableVault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableVault *UpgradeableVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.TransferOwnership(&_UpgradeableVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.TransferOwnership(&_UpgradeableVault.TransactOpts, newOwner)
}

// UpdateTreasuryAddress is a paid mutator transaction binding the contract method 0x841e4561.
//
// Solidity: function updateTreasuryAddress(address _treasury) returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) UpdateTreasuryAddress(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "updateTreasuryAddress", _treasury)
}

// UpdateTreasuryAddress is a paid mutator transaction binding the contract method 0x841e4561.
//
// Solidity: function updateTreasuryAddress(address _treasury) returns()
func (_UpgradeableVault *UpgradeableVaultSession) UpdateTreasuryAddress(_treasury common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.UpdateTreasuryAddress(&_UpgradeableVault.TransactOpts, _treasury)
}

// UpdateTreasuryAddress is a paid mutator transaction binding the contract method 0x841e4561.
//
// Solidity: function updateTreasuryAddress(address _treasury) returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) UpdateTreasuryAddress(_treasury common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.UpdateTreasuryAddress(&_UpgradeableVault.TransactOpts, _treasury)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UpgradeableVault *UpgradeableVaultTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UpgradeableVault *UpgradeableVaultSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.UpgradeToAndCall(&_UpgradeableVault.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UpgradeableVault *UpgradeableVaultTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.UpgradeToAndCall(&_UpgradeableVault.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Withdraw(&_UpgradeableVault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_UpgradeableVault *UpgradeableVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _UpgradeableVault.Contract.Withdraw(&_UpgradeableVault.TransactOpts, assets, receiver, owner)
}

// UpgradeableVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UpgradeableVault contract.
type UpgradeableVaultApprovalIterator struct {
	Event *UpgradeableVaultApproval // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultApproval)
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
		it.Event = new(UpgradeableVaultApproval)
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
func (it *UpgradeableVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultApproval represents a Approval event raised by the UpgradeableVault contract.
type UpgradeableVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UpgradeableVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultApprovalIterator{contract: _UpgradeableVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultApproval)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseApproval(log types.Log) (*UpgradeableVaultApproval, error) {
	event := new(UpgradeableVaultApproval)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the UpgradeableVault contract.
type UpgradeableVaultClaimedIterator struct {
	Event *UpgradeableVaultClaimed // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultClaimed)
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
		it.Event = new(UpgradeableVaultClaimed)
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
func (it *UpgradeableVaultClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultClaimed represents a Claimed event raised by the UpgradeableVault contract.
type UpgradeableVaultClaimed struct {
	User     common.Address
	Receiver common.Address
	Claimed  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address user, address receiver, uint256 claimed)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterClaimed(opts *bind.FilterOpts) (*UpgradeableVaultClaimedIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultClaimedIterator{contract: _UpgradeableVault.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address user, address receiver, uint256 claimed)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultClaimed) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultClaimed)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address user, address receiver, uint256 claimed)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseClaimed(log types.Log) (*UpgradeableVaultClaimed, error) {
	event := new(UpgradeableVaultClaimed)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the UpgradeableVault contract.
type UpgradeableVaultContextDataRevertIterator struct {
	Event *UpgradeableVaultContextDataRevert // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultContextDataRevert)
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
		it.Event = new(UpgradeableVaultContextDataRevert)
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
func (it *UpgradeableVaultContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultContextDataRevert represents a ContextDataRevert event raised by the UpgradeableVault contract.
type UpgradeableVaultContextDataRevert struct {
	Arg0 RevertContext
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0x35a9324413457251c1059312318f6f1cec6bd0da4105d01315f3151b1e3a2c76.
//
// Solidity: event ContextDataRevert((address,uint64,bytes) arg0)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*UpgradeableVaultContextDataRevertIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultContextDataRevertIterator{contract: _UpgradeableVault.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0x35a9324413457251c1059312318f6f1cec6bd0da4105d01315f3151b1e3a2c76.
//
// Solidity: event ContextDataRevert((address,uint64,bytes) arg0)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultContextDataRevert)
				if err := _UpgradeableVault.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0x35a9324413457251c1059312318f6f1cec6bd0da4105d01315f3151b1e3a2c76.
//
// Solidity: event ContextDataRevert((address,uint64,bytes) arg0)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseContextDataRevert(log types.Log) (*UpgradeableVaultContextDataRevert, error) {
	event := new(UpgradeableVaultContextDataRevert)
	if err := _UpgradeableVault.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the UpgradeableVault contract.
type UpgradeableVaultDepositIterator struct {
	Event *UpgradeableVaultDeposit // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultDeposit)
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
		it.Event = new(UpgradeableVaultDeposit)
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
func (it *UpgradeableVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultDeposit represents a Deposit event raised by the UpgradeableVault contract.
type UpgradeableVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*UpgradeableVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultDepositIterator{contract: _UpgradeableVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultDeposit)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseDeposit(log types.Log) (*UpgradeableVaultDeposit, error) {
	event := new(UpgradeableVaultDeposit)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpgradeableVault contract.
type UpgradeableVaultInitializedIterator struct {
	Event *UpgradeableVaultInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultInitialized)
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
		it.Event = new(UpgradeableVaultInitialized)
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
func (it *UpgradeableVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultInitialized represents a Initialized event raised by the UpgradeableVault contract.
type UpgradeableVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeableVaultInitializedIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultInitializedIterator{contract: _UpgradeableVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultInitialized)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseInitialized(log types.Log) (*UpgradeableVaultInitialized, error) {
	event := new(UpgradeableVaultInitialized)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpgradeableVault contract.
type UpgradeableVaultOwnershipTransferredIterator struct {
	Event *UpgradeableVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultOwnershipTransferred)
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
		it.Event = new(UpgradeableVaultOwnershipTransferred)
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
func (it *UpgradeableVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultOwnershipTransferred represents a OwnershipTransferred event raised by the UpgradeableVault contract.
type UpgradeableVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradeableVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultOwnershipTransferredIterator{contract: _UpgradeableVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultOwnershipTransferred)
				if err := _UpgradeableVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradeableVaultOwnershipTransferred, error) {
	event := new(UpgradeableVaultOwnershipTransferred)
	if err := _UpgradeableVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultPerformanceFeePaidIterator is returned from FilterPerformanceFeePaid and is used to iterate over the raw logs and unpacked data for PerformanceFeePaid events raised by the UpgradeableVault contract.
type UpgradeableVaultPerformanceFeePaidIterator struct {
	Event *UpgradeableVaultPerformanceFeePaid // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultPerformanceFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultPerformanceFeePaid)
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
		it.Event = new(UpgradeableVaultPerformanceFeePaid)
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
func (it *UpgradeableVaultPerformanceFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultPerformanceFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultPerformanceFeePaid represents a PerformanceFeePaid event raised by the UpgradeableVault contract.
type UpgradeableVaultPerformanceFeePaid struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPerformanceFeePaid is a free log retrieval operation binding the contract event 0xafdbf88d8a6e3bd4dbd770e775c7e168f61db74aea5e21852ac1061bf1cb625d.
//
// Solidity: event PerformanceFeePaid(address indexed user, uint256 amount)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterPerformanceFeePaid(opts *bind.FilterOpts, user []common.Address) (*UpgradeableVaultPerformanceFeePaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "PerformanceFeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultPerformanceFeePaidIterator{contract: _UpgradeableVault.contract, event: "PerformanceFeePaid", logs: logs, sub: sub}, nil
}

// WatchPerformanceFeePaid is a free log subscription operation binding the contract event 0xafdbf88d8a6e3bd4dbd770e775c7e168f61db74aea5e21852ac1061bf1cb625d.
//
// Solidity: event PerformanceFeePaid(address indexed user, uint256 amount)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchPerformanceFeePaid(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultPerformanceFeePaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "PerformanceFeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultPerformanceFeePaid)
				if err := _UpgradeableVault.contract.UnpackLog(event, "PerformanceFeePaid", log); err != nil {
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

// ParsePerformanceFeePaid is a log parse operation binding the contract event 0xafdbf88d8a6e3bd4dbd770e775c7e168f61db74aea5e21852ac1061bf1cb625d.
//
// Solidity: event PerformanceFeePaid(address indexed user, uint256 amount)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParsePerformanceFeePaid(log types.Log) (*UpgradeableVaultPerformanceFeePaid, error) {
	event := new(UpgradeableVaultPerformanceFeePaid)
	if err := _UpgradeableVault.contract.UnpackLog(event, "PerformanceFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultPerformanceFeeUpdatedIterator is returned from FilterPerformanceFeeUpdated and is used to iterate over the raw logs and unpacked data for PerformanceFeeUpdated events raised by the UpgradeableVault contract.
type UpgradeableVaultPerformanceFeeUpdatedIterator struct {
	Event *UpgradeableVaultPerformanceFeeUpdated // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultPerformanceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultPerformanceFeeUpdated)
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
		it.Event = new(UpgradeableVaultPerformanceFeeUpdated)
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
func (it *UpgradeableVaultPerformanceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultPerformanceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultPerformanceFeeUpdated represents a PerformanceFeeUpdated event raised by the UpgradeableVault contract.
type UpgradeableVaultPerformanceFeeUpdated struct {
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPerformanceFeeUpdated is a free log retrieval operation binding the contract event 0x9b49d0cd76012d9c67241c2f68f836efbaf50ea29901a250040671402d5263f5.
//
// Solidity: event PerformanceFeeUpdated(uint256 newFeeRate)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterPerformanceFeeUpdated(opts *bind.FilterOpts) (*UpgradeableVaultPerformanceFeeUpdatedIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "PerformanceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultPerformanceFeeUpdatedIterator{contract: _UpgradeableVault.contract, event: "PerformanceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPerformanceFeeUpdated is a free log subscription operation binding the contract event 0x9b49d0cd76012d9c67241c2f68f836efbaf50ea29901a250040671402d5263f5.
//
// Solidity: event PerformanceFeeUpdated(uint256 newFeeRate)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchPerformanceFeeUpdated(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultPerformanceFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "PerformanceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultPerformanceFeeUpdated)
				if err := _UpgradeableVault.contract.UnpackLog(event, "PerformanceFeeUpdated", log); err != nil {
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

// ParsePerformanceFeeUpdated is a log parse operation binding the contract event 0x9b49d0cd76012d9c67241c2f68f836efbaf50ea29901a250040671402d5263f5.
//
// Solidity: event PerformanceFeeUpdated(uint256 newFeeRate)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParsePerformanceFeeUpdated(log types.Log) (*UpgradeableVaultPerformanceFeeUpdated, error) {
	event := new(UpgradeableVaultPerformanceFeeUpdated)
	if err := _UpgradeableVault.contract.UnpackLog(event, "PerformanceFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultRewardsPerTokenUpdatedIterator is returned from FilterRewardsPerTokenUpdated and is used to iterate over the raw logs and unpacked data for RewardsPerTokenUpdated events raised by the UpgradeableVault contract.
type UpgradeableVaultRewardsPerTokenUpdatedIterator struct {
	Event *UpgradeableVaultRewardsPerTokenUpdated // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultRewardsPerTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultRewardsPerTokenUpdated)
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
		it.Event = new(UpgradeableVaultRewardsPerTokenUpdated)
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
func (it *UpgradeableVaultRewardsPerTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultRewardsPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultRewardsPerTokenUpdated represents a RewardsPerTokenUpdated event raised by the UpgradeableVault contract.
type UpgradeableVaultRewardsPerTokenUpdated struct {
	Accumulated *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsPerTokenUpdated is a free log retrieval operation binding the contract event 0xe972555b20cae8150e291bb40efce3ef4e3ed6b6b2c39c029346600e95469d48.
//
// Solidity: event RewardsPerTokenUpdated(uint256 accumulated)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterRewardsPerTokenUpdated(opts *bind.FilterOpts) (*UpgradeableVaultRewardsPerTokenUpdatedIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "RewardsPerTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultRewardsPerTokenUpdatedIterator{contract: _UpgradeableVault.contract, event: "RewardsPerTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsPerTokenUpdated is a free log subscription operation binding the contract event 0xe972555b20cae8150e291bb40efce3ef4e3ed6b6b2c39c029346600e95469d48.
//
// Solidity: event RewardsPerTokenUpdated(uint256 accumulated)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchRewardsPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultRewardsPerTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "RewardsPerTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultRewardsPerTokenUpdated)
				if err := _UpgradeableVault.contract.UnpackLog(event, "RewardsPerTokenUpdated", log); err != nil {
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

// ParseRewardsPerTokenUpdated is a log parse operation binding the contract event 0xe972555b20cae8150e291bb40efce3ef4e3ed6b6b2c39c029346600e95469d48.
//
// Solidity: event RewardsPerTokenUpdated(uint256 accumulated)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseRewardsPerTokenUpdated(log types.Log) (*UpgradeableVaultRewardsPerTokenUpdated, error) {
	event := new(UpgradeableVaultRewardsPerTokenUpdated)
	if err := _UpgradeableVault.contract.UnpackLog(event, "RewardsPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultRewardsSetIterator is returned from FilterRewardsSet and is used to iterate over the raw logs and unpacked data for RewardsSet events raised by the UpgradeableVault contract.
type UpgradeableVaultRewardsSetIterator struct {
	Event *UpgradeableVaultRewardsSet // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultRewardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultRewardsSet)
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
		it.Event = new(UpgradeableVaultRewardsSet)
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
func (it *UpgradeableVaultRewardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultRewardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultRewardsSet represents a RewardsSet event raised by the UpgradeableVault contract.
type UpgradeableVaultRewardsSet struct {
	Start uint32
	End   uint32
	Rate  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSet is a free log retrieval operation binding the contract event 0x95efd8a2a0aa591f48fd9673cf5d13c2150ca7a1fe1cbe438dd3f0ae47064663.
//
// Solidity: event RewardsSet(uint32 start, uint32 end, uint256 rate)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterRewardsSet(opts *bind.FilterOpts) (*UpgradeableVaultRewardsSetIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "RewardsSet")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultRewardsSetIterator{contract: _UpgradeableVault.contract, event: "RewardsSet", logs: logs, sub: sub}, nil
}

// WatchRewardsSet is a free log subscription operation binding the contract event 0x95efd8a2a0aa591f48fd9673cf5d13c2150ca7a1fe1cbe438dd3f0ae47064663.
//
// Solidity: event RewardsSet(uint32 start, uint32 end, uint256 rate)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchRewardsSet(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultRewardsSet) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "RewardsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultRewardsSet)
				if err := _UpgradeableVault.contract.UnpackLog(event, "RewardsSet", log); err != nil {
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

// ParseRewardsSet is a log parse operation binding the contract event 0x95efd8a2a0aa591f48fd9673cf5d13c2150ca7a1fe1cbe438dd3f0ae47064663.
//
// Solidity: event RewardsSet(uint32 start, uint32 end, uint256 rate)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseRewardsSet(log types.Log) (*UpgradeableVaultRewardsSet, error) {
	event := new(UpgradeableVaultRewardsSet)
	if err := _UpgradeableVault.contract.UnpackLog(event, "RewardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultStrategyUpdatedIterator is returned from FilterStrategyUpdated and is used to iterate over the raw logs and unpacked data for StrategyUpdated events raised by the UpgradeableVault contract.
type UpgradeableVaultStrategyUpdatedIterator struct {
	Event *UpgradeableVaultStrategyUpdated // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultStrategyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultStrategyUpdated)
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
		it.Event = new(UpgradeableVaultStrategyUpdated)
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
func (it *UpgradeableVaultStrategyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultStrategyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultStrategyUpdated represents a StrategyUpdated event raised by the UpgradeableVault contract.
type UpgradeableVaultStrategyUpdated struct {
	NewStrategyAddress common.Address
	NewStrategyChainId uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdated is a free log retrieval operation binding the contract event 0x3f1cc356d56ba3d612ed5b8fa9d919d35bfc89e02529efd6723156bffbd24042.
//
// Solidity: event StrategyUpdated(address indexed newStrategyAddress, uint32 newStrategyChainId)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterStrategyUpdated(opts *bind.FilterOpts, newStrategyAddress []common.Address) (*UpgradeableVaultStrategyUpdatedIterator, error) {

	var newStrategyAddressRule []interface{}
	for _, newStrategyAddressItem := range newStrategyAddress {
		newStrategyAddressRule = append(newStrategyAddressRule, newStrategyAddressItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "StrategyUpdated", newStrategyAddressRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultStrategyUpdatedIterator{contract: _UpgradeableVault.contract, event: "StrategyUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdated is a free log subscription operation binding the contract event 0x3f1cc356d56ba3d612ed5b8fa9d919d35bfc89e02529efd6723156bffbd24042.
//
// Solidity: event StrategyUpdated(address indexed newStrategyAddress, uint32 newStrategyChainId)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchStrategyUpdated(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultStrategyUpdated, newStrategyAddress []common.Address) (event.Subscription, error) {

	var newStrategyAddressRule []interface{}
	for _, newStrategyAddressItem := range newStrategyAddress {
		newStrategyAddressRule = append(newStrategyAddressRule, newStrategyAddressItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "StrategyUpdated", newStrategyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultStrategyUpdated)
				if err := _UpgradeableVault.contract.UnpackLog(event, "StrategyUpdated", log); err != nil {
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

// ParseStrategyUpdated is a log parse operation binding the contract event 0x3f1cc356d56ba3d612ed5b8fa9d919d35bfc89e02529efd6723156bffbd24042.
//
// Solidity: event StrategyUpdated(address indexed newStrategyAddress, uint32 newStrategyChainId)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseStrategyUpdated(log types.Log) (*UpgradeableVaultStrategyUpdated, error) {
	event := new(UpgradeableVaultStrategyUpdated)
	if err := _UpgradeableVault.contract.UnpackLog(event, "StrategyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UpgradeableVault contract.
type UpgradeableVaultTransferIterator struct {
	Event *UpgradeableVaultTransfer // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultTransfer)
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
		it.Event = new(UpgradeableVaultTransfer)
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
func (it *UpgradeableVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultTransfer represents a Transfer event raised by the UpgradeableVault contract.
type UpgradeableVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UpgradeableVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultTransferIterator{contract: _UpgradeableVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultTransfer)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseTransfer(log types.Log) (*UpgradeableVaultTransfer, error) {
	event := new(UpgradeableVaultTransfer)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UpgradeableVault contract.
type UpgradeableVaultUpgradedIterator struct {
	Event *UpgradeableVaultUpgraded // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultUpgraded)
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
		it.Event = new(UpgradeableVaultUpgraded)
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
func (it *UpgradeableVaultUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultUpgraded represents a Upgraded event raised by the UpgradeableVault contract.
type UpgradeableVaultUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UpgradeableVaultUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultUpgradedIterator{contract: _UpgradeableVault.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultUpgraded)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseUpgraded(log types.Log) (*UpgradeableVaultUpgraded, error) {
	event := new(UpgradeableVaultUpgraded)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultUserRewardsUpdatedIterator is returned from FilterUserRewardsUpdated and is used to iterate over the raw logs and unpacked data for UserRewardsUpdated events raised by the UpgradeableVault contract.
type UpgradeableVaultUserRewardsUpdatedIterator struct {
	Event *UpgradeableVaultUserRewardsUpdated // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultUserRewardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultUserRewardsUpdated)
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
		it.Event = new(UpgradeableVaultUserRewardsUpdated)
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
func (it *UpgradeableVaultUserRewardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultUserRewardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultUserRewardsUpdated represents a UserRewardsUpdated event raised by the UpgradeableVault contract.
type UpgradeableVaultUserRewardsUpdated struct {
	User               common.Address
	UserRewards        *big.Int
	PaidRewardPerToken *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUserRewardsUpdated is a free log retrieval operation binding the contract event 0x5b9aaf4cc5141c090a75f8b8a627863eba92df81f0c83c096350da9b79aedd04.
//
// Solidity: event UserRewardsUpdated(address user, uint256 userRewards, uint256 paidRewardPerToken)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterUserRewardsUpdated(opts *bind.FilterOpts) (*UpgradeableVaultUserRewardsUpdatedIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "UserRewardsUpdated")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultUserRewardsUpdatedIterator{contract: _UpgradeableVault.contract, event: "UserRewardsUpdated", logs: logs, sub: sub}, nil
}

// WatchUserRewardsUpdated is a free log subscription operation binding the contract event 0x5b9aaf4cc5141c090a75f8b8a627863eba92df81f0c83c096350da9b79aedd04.
//
// Solidity: event UserRewardsUpdated(address user, uint256 userRewards, uint256 paidRewardPerToken)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchUserRewardsUpdated(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultUserRewardsUpdated) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "UserRewardsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultUserRewardsUpdated)
				if err := _UpgradeableVault.contract.UnpackLog(event, "UserRewardsUpdated", log); err != nil {
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

// ParseUserRewardsUpdated is a log parse operation binding the contract event 0x5b9aaf4cc5141c090a75f8b8a627863eba92df81f0c83c096350da9b79aedd04.
//
// Solidity: event UserRewardsUpdated(address user, uint256 userRewards, uint256 paidRewardPerToken)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseUserRewardsUpdated(log types.Log) (*UpgradeableVaultUserRewardsUpdated, error) {
	event := new(UpgradeableVaultUserRewardsUpdated)
	if err := _UpgradeableVault.contract.UnpackLog(event, "UserRewardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultVaultInitializedIterator is returned from FilterVaultInitialized and is used to iterate over the raw logs and unpacked data for VaultInitialized events raised by the UpgradeableVault contract.
type UpgradeableVaultVaultInitializedIterator struct {
	Event *UpgradeableVaultVaultInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultVaultInitialized)
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
		it.Event = new(UpgradeableVaultVaultInitialized)
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
func (it *UpgradeableVaultVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultVaultInitialized represents a VaultInitialized event raised by the UpgradeableVault contract.
type UpgradeableVaultVaultInitialized struct {
	Decimals uint8
	PerfFee  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVaultInitialized is a free log retrieval operation binding the contract event 0x7203e25960677692487f12bc9315b5b0cc817a4cbc645714efcb9bcef4b74dca.
//
// Solidity: event VaultInitialized(uint8 decimals, uint256 perfFee)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterVaultInitialized(opts *bind.FilterOpts) (*UpgradeableVaultVaultInitializedIterator, error) {

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "VaultInitialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultVaultInitializedIterator{contract: _UpgradeableVault.contract, event: "VaultInitialized", logs: logs, sub: sub}, nil
}

// WatchVaultInitialized is a free log subscription operation binding the contract event 0x7203e25960677692487f12bc9315b5b0cc817a4cbc645714efcb9bcef4b74dca.
//
// Solidity: event VaultInitialized(uint8 decimals, uint256 perfFee)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchVaultInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "VaultInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultVaultInitialized)
				if err := _UpgradeableVault.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
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

// ParseVaultInitialized is a log parse operation binding the contract event 0x7203e25960677692487f12bc9315b5b0cc817a4cbc645714efcb9bcef4b74dca.
//
// Solidity: event VaultInitialized(uint8 decimals, uint256 perfFee)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseVaultInitialized(log types.Log) (*UpgradeableVaultVaultInitialized, error) {
	event := new(UpgradeableVaultVaultInitialized)
	if err := _UpgradeableVault.contract.UnpackLog(event, "VaultInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the UpgradeableVault contract.
type UpgradeableVaultWithdrawIterator struct {
	Event *UpgradeableVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *UpgradeableVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableVaultWithdraw)
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
		it.Event = new(UpgradeableVaultWithdraw)
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
func (it *UpgradeableVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableVaultWithdraw represents a Withdraw event raised by the UpgradeableVault contract.
type UpgradeableVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_UpgradeableVault *UpgradeableVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*UpgradeableVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UpgradeableVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableVaultWithdrawIterator{contract: _UpgradeableVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_UpgradeableVault *UpgradeableVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *UpgradeableVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _UpgradeableVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableVaultWithdraw)
				if err := _UpgradeableVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_UpgradeableVault *UpgradeableVaultFilterer) ParseWithdraw(log types.Log) (*UpgradeableVaultWithdraw, error) {
	event := new(UpgradeableVaultWithdraw)
	if err := _UpgradeableVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
