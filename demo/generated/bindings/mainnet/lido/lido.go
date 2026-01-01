// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// LidoMetaData contains all meta data concerning the Lido contract.
var LidoMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthBySharesRoundUp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amountOfShares\",\"type\":\"uint256\"}],\"name\":\"mintExternalShares\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CONTROL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportTimestamp\",\"type\":\"uint256\"},{\"name\":\"_timeElapsed\",\"type\":\"uint256\"},{\"name\":\"_preTotalShares\",\"type\":\"uint256\"},{\"name\":\"_preTotalEther\",\"type\":\"uint256\"},{\"name\":\"_postTotalShares\",\"type\":\"uint256\"},{\"name\":\"_postTotalEther\",\"type\":\"uint256\"},{\"name\":\"_postInternalShares\",\"type\":\"uint256\"},{\"name\":\"_postInternalEther\",\"type\":\"uint256\"},{\"name\":\"_sharesMintedAsFees\",\"type\":\"uint256\"}],\"name\":\"emitTokenRebase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStakingPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"_stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDepositedValidators\",\"type\":\"uint256\"}],\"name\":\"unsafeChangeDepositedValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBufferedEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amountOfShares\",\"type\":\"uint256\"}],\"name\":\"mintShares\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportTimestamp\",\"type\":\"uint256\"},{\"name\":\"_preClValidators\",\"type\":\"uint256\"},{\"name\":\"_reportClValidators\",\"type\":\"uint256\"},{\"name\":\"_reportClBalance\",\"type\":\"uint256\"}],\"name\":\"processClStateUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentStakeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExternalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfShares\",\"type\":\"uint256\"}],\"name\":\"internalizeExternalBadDebt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeLimitFullInfo\",\"outputs\":[{\"name\":\"isStakingPaused_\",\"type\":\"bool\"},{\"name\":\"isStakingLimitSet\",\"type\":\"bool\"},{\"name\":\"currentStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimitGrowthBlocks\",\"type\":\"uint256\"},{\"name\":\"prevStakeLimit\",\"type\":\"uint256\"},{\"name\":\"prevStakeBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferSharesFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfShares\",\"type\":\"uint256\"}],\"name\":\"burnExternalShares\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeDistribution\",\"outputs\":[{\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveWithdrawals\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfShares\",\"type\":\"uint256\"}],\"name\":\"rebalanceExternalEtherToInternal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountOfShares\",\"type\":\"uint256\"}],\"name\":\"burnShares\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldBurner\",\"type\":\"address\"},{\"name\":\"_contractsWithBurnerAllowances\",\"type\":\"address[]\"},{\"name\":\"_initialMaxExternalRatioBP\",\"type\":\"uint256\"}],\"name\":\"finalizeUpgrade_v3\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxMintableExternalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportTimestamp\",\"type\":\"uint256\"},{\"name\":\"_reportClBalance\",\"type\":\"uint256\"},{\"name\":\"_principalCLBalance\",\"type\":\"uint256\"},{\"name\":\"_withdrawalsToWithdraw\",\"type\":\"uint256\"},{\"name\":\"_elRewardsToWithdraw\",\"type\":\"uint256\"},{\"name\":\"_lastWithdrawalRequestToFinalize\",\"type\":\"uint256\"},{\"name\":\"_withdrawalsShareRate\",\"type\":\"uint256\"},{\"name\":\"_etherToLockOnWithdrawalQueue\",\"type\":\"uint256\"}],\"name\":\"collectRewardsAndProcessWithdrawals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEIP712StETH\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxExternalRatioBP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDepositsCount\",\"type\":\"uint256\"},{\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"name\":\"_depositCalldata\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBeaconStat\",\"outputs\":[{\"name\":\"depositedValidators\",\"type\":\"uint256\"},{\"name\":\"beaconValidators\",\"type\":\"uint256\"},{\"name\":\"beaconBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"totalFee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExternalEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLidoLocator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositableEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxExternalRatioBP\",\"type\":\"uint256\"}],\"name\":\"setMaxExternalRatioBP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalELRewardsCollected\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"StakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLValidators\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLValidators\",\"type\":\"uint256\"}],\"name\":\"CLValidatorsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositedValidators\",\"type\":\"uint256\"}],\"name\":\"DepositedValidatorsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"executionLayerRewardsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postBufferedEther\",\"type\":\"uint256\"}],\"name\":\"ETHDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeElapsed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesMintedAsFees\",\"type\":\"uint256\"}],\"name\":\"TokenRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lidoLocator\",\"type\":\"address\"}],\"name\":\"LidoLocatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postInternalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postInternalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesMintedAsFees\",\"type\":\"uint256\"}],\"name\":\"InternalShareRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountOfShares\",\"type\":\"uint256\"}],\"name\":\"ExternalSharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amountOfShares\",\"type\":\"uint256\"}],\"name\":\"ExternalSharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxExternalRatioBP\",\"type\":\"uint256\"}],\"name\":\"MaxExternalRatioBPSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExternalEtherTransferredToBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amountOfShares\",\"type\":\"uint256\"}],\"name\":\"ExternalBadDebtInternalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eip712StETH\",\"type\":\"address\"}],\"name\":\"EIP712StETHInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"TransferShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"SharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"}]",
}

// LidoABI is the input ABI used to generate the binding from.
// Deprecated: Use LidoMetaData.ABI instead.
var LidoABI = LidoMetaData.ABI

// Lido is an auto generated Go binding around an Ethereum contract.
type Lido struct {
	LidoCaller     // Read-only binding to the contract
	LidoTransactor // Write-only binding to the contract
	LidoFilterer   // Log filterer for contract events
}

// LidoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoSession struct {
	Contract     *Lido             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoCallerSession struct {
	Contract *LidoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LidoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoTransactorSession struct {
	Contract     *LidoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoRaw struct {
	Contract *Lido // Generic contract binding to access the raw methods on
}

// LidoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoCallerRaw struct {
	Contract *LidoCaller // Generic read-only contract binding to access the raw methods on
}

// LidoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoTransactorRaw struct {
	Contract *LidoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLido creates a new instance of Lido, bound to a specific deployed contract.
func NewLido(address common.Address, backend bind.ContractBackend) (*Lido, error) {
	contract, err := bindLido(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lido{LidoCaller: LidoCaller{contract: contract}, LidoTransactor: LidoTransactor{contract: contract}, LidoFilterer: LidoFilterer{contract: contract}}, nil
}

// NewLidoCaller creates a new read-only instance of Lido, bound to a specific deployed contract.
func NewLidoCaller(address common.Address, caller bind.ContractCaller) (*LidoCaller, error) {
	contract, err := bindLido(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoCaller{contract: contract}, nil
}

// NewLidoTransactor creates a new write-only instance of Lido, bound to a specific deployed contract.
func NewLidoTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoTransactor, error) {
	contract, err := bindLido(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoTransactor{contract: contract}, nil
}

// NewLidoFilterer creates a new log filterer instance of Lido, bound to a specific deployed contract.
func NewLidoFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoFilterer, error) {
	contract, err := bindLido(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoFilterer{contract: contract}, nil
}

// bindLido binds a generic wrapper to an already deployed contract.
func bindLido(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LidoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lido *LidoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lido.Contract.LidoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lido *LidoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.Contract.LidoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lido *LidoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lido.Contract.LidoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lido *LidoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lido.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lido *LidoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lido *LidoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lido.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Lido *LidoCaller) DOMAIN_SEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Lido *LidoSession) DOMAIN_SEPARATOR() ([32]byte, error) {
	return _Lido.Contract.DOMAIN_SEPARATOR(&_Lido.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Lido *LidoCallerSession) DOMAIN_SEPARATOR() ([32]byte, error) {
	return _Lido.Contract.DOMAIN_SEPARATOR(&_Lido.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) PAUSE_ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoSession) PAUSE_ROLE() ([32]byte, error) {
	return _Lido.Contract.PAUSE_ROLE(&_Lido.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) PAUSE_ROLE() ([32]byte, error) {
	return _Lido.Contract.PAUSE_ROLE(&_Lido.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) RESUME_ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Lido *LidoSession) RESUME_ROLE() ([32]byte, error) {
	return _Lido.Contract.RESUME_ROLE(&_Lido.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) RESUME_ROLE() ([32]byte, error) {
	return _Lido.Contract.RESUME_ROLE(&_Lido.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) STAKING_CONTROL_ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "STAKING_CONTROL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_Lido *LidoSession) STAKING_CONTROL_ROLE() ([32]byte, error) {
	return _Lido.Contract.STAKING_CONTROL_ROLE(&_Lido.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) STAKING_CONTROL_ROLE() ([32]byte, error) {
	return _Lido.Contract.STAKING_CONTROL_ROLE(&_Lido.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) STAKING_PAUSE_ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "STAKING_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoSession) STAKING_PAUSE_ROLE() ([32]byte, error) {
	return _Lido.Contract.STAKING_PAUSE_ROLE(&_Lido.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) STAKING_PAUSE_ROLE() ([32]byte, error) {
	return _Lido.Contract.STAKING_PAUSE_ROLE(&_Lido.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Lido *LidoSession) UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() ([32]byte, error) {
	return _Lido.Contract.UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE(&_Lido.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() ([32]byte, error) {
	return _Lido.Contract.UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE(&_Lido.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Lido *LidoCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Lido *LidoSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Lido.Contract.AllowRecoverability(&_Lido.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Lido *LidoCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Lido.Contract.AllowRecoverability(&_Lido.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lido *LidoCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lido *LidoSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lido.Contract.Allowance(&_Lido.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lido *LidoCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lido.Contract.Allowance(&_Lido.CallOpts, _owner, _spender)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Lido *LidoCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Lido *LidoSession) AppId() ([32]byte, error) {
	return _Lido.Contract.AppId(&_Lido.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Lido *LidoCallerSession) AppId() ([32]byte, error) {
	return _Lido.Contract.AppId(&_Lido.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lido *LidoCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lido *LidoSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.BalanceOf(&_Lido.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lido *LidoCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.BalanceOf(&_Lido.CallOpts, _account)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_Lido *LidoCaller) CanDeposit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "canDeposit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_Lido *LidoSession) CanDeposit() (bool, error) {
	return _Lido.Contract.CanDeposit(&_Lido.CallOpts)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_Lido *LidoCallerSession) CanDeposit() (bool, error) {
	return _Lido.Contract.CanDeposit(&_Lido.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Lido *LidoCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Lido *LidoSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Lido.Contract.CanPerform(&_Lido.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Lido *LidoCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Lido.Contract.CanPerform(&_Lido.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Lido *LidoCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Lido *LidoSession) Decimals() (uint8, error) {
	return _Lido.Contract.Decimals(&_Lido.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Lido *LidoCallerSession) Decimals() (uint8, error) {
	return _Lido.Contract.Decimals(&_Lido.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_Lido *LidoCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_Lido *LidoSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _Lido.Contract.Eip712Domain(&_Lido.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_Lido *LidoCallerSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _Lido.Contract.Eip712Domain(&_Lido.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_Lido *LidoCaller) GetBeaconStat(opts *bind.CallOpts) (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getBeaconStat")

	outstruct := new(struct {
		DepositedValidators *big.Int
		BeaconValidators    *big.Int
		BeaconBalance       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BeaconValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeaconBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_Lido *LidoSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _Lido.Contract.GetBeaconStat(&_Lido.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_Lido *LidoCallerSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _Lido.Contract.GetBeaconStat(&_Lido.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_Lido *LidoCaller) GetBufferedEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getBufferedEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_Lido *LidoSession) GetBufferedEther() (*big.Int, error) {
	return _Lido.Contract.GetBufferedEther(&_Lido.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_Lido *LidoCallerSession) GetBufferedEther() (*big.Int, error) {
	return _Lido.Contract.GetBufferedEther(&_Lido.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Lido *LidoCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Lido *LidoSession) GetContractVersion() (*big.Int, error) {
	return _Lido.Contract.GetContractVersion(&_Lido.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Lido *LidoCallerSession) GetContractVersion() (*big.Int, error) {
	return _Lido.Contract.GetContractVersion(&_Lido.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_Lido *LidoCaller) GetCurrentStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getCurrentStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_Lido *LidoSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _Lido.Contract.GetCurrentStakeLimit(&_Lido.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_Lido *LidoCallerSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _Lido.Contract.GetCurrentStakeLimit(&_Lido.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_Lido *LidoCaller) GetDepositableEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getDepositableEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_Lido *LidoSession) GetDepositableEther() (*big.Int, error) {
	return _Lido.Contract.GetDepositableEther(&_Lido.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_Lido *LidoCallerSession) GetDepositableEther() (*big.Int, error) {
	return _Lido.Contract.GetDepositableEther(&_Lido.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_Lido *LidoCaller) GetEIP712StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getEIP712StETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_Lido *LidoSession) GetEIP712StETH() (common.Address, error) {
	return _Lido.Contract.GetEIP712StETH(&_Lido.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_Lido *LidoCallerSession) GetEIP712StETH() (common.Address, error) {
	return _Lido.Contract.GetEIP712StETH(&_Lido.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Lido *LidoCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Lido *LidoSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Lido.Contract.GetEVMScriptExecutor(&_Lido.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Lido *LidoCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Lido.Contract.GetEVMScriptExecutor(&_Lido.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Lido *LidoCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Lido *LidoSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Lido.Contract.GetEVMScriptRegistry(&_Lido.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Lido *LidoCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Lido.Contract.GetEVMScriptRegistry(&_Lido.CallOpts)
}

// GetExternalEther is a free data retrieval call binding the contract method 0xe16a9065.
//
// Solidity: function getExternalEther() view returns(uint256)
func (_Lido *LidoCaller) GetExternalEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getExternalEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExternalEther is a free data retrieval call binding the contract method 0xe16a9065.
//
// Solidity: function getExternalEther() view returns(uint256)
func (_Lido *LidoSession) GetExternalEther() (*big.Int, error) {
	return _Lido.Contract.GetExternalEther(&_Lido.CallOpts)
}

// GetExternalEther is a free data retrieval call binding the contract method 0xe16a9065.
//
// Solidity: function getExternalEther() view returns(uint256)
func (_Lido *LidoCallerSession) GetExternalEther() (*big.Int, error) {
	return _Lido.Contract.GetExternalEther(&_Lido.CallOpts)
}

// GetExternalShares is a free data retrieval call binding the contract method 0x63021d8b.
//
// Solidity: function getExternalShares() view returns(uint256)
func (_Lido *LidoCaller) GetExternalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getExternalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExternalShares is a free data retrieval call binding the contract method 0x63021d8b.
//
// Solidity: function getExternalShares() view returns(uint256)
func (_Lido *LidoSession) GetExternalShares() (*big.Int, error) {
	return _Lido.Contract.GetExternalShares(&_Lido.CallOpts)
}

// GetExternalShares is a free data retrieval call binding the contract method 0x63021d8b.
//
// Solidity: function getExternalShares() view returns(uint256)
func (_Lido *LidoCallerSession) GetExternalShares() (*big.Int, error) {
	return _Lido.Contract.GetExternalShares(&_Lido.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_Lido *LidoCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_Lido *LidoSession) GetFee() (uint16, error) {
	return _Lido.Contract.GetFee(&_Lido.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_Lido *LidoCallerSession) GetFee() (uint16, error) {
	return _Lido.Contract.GetFee(&_Lido.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoCaller) GetFeeDistribution(opts *bind.CallOpts) (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getFeeDistribution")

	outstruct := new(struct {
		TreasuryFeeBasisPoints  uint16
		InsuranceFeeBasisPoints uint16
		OperatorsFeeBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TreasuryFeeBasisPoints = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InsuranceFeeBasisPoints = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.OperatorsFeeBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _Lido.Contract.GetFeeDistribution(&_Lido.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoCallerSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _Lido.Contract.GetFeeDistribution(&_Lido.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Lido *LidoCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Lido *LidoSession) GetInitializationBlock() (*big.Int, error) {
	return _Lido.Contract.GetInitializationBlock(&_Lido.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Lido *LidoCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _Lido.Contract.GetInitializationBlock(&_Lido.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_Lido *LidoCaller) GetLidoLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getLidoLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_Lido *LidoSession) GetLidoLocator() (common.Address, error) {
	return _Lido.Contract.GetLidoLocator(&_Lido.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_Lido *LidoCallerSession) GetLidoLocator() (common.Address, error) {
	return _Lido.Contract.GetLidoLocator(&_Lido.CallOpts)
}

// GetMaxExternalRatioBP is a free data retrieval call binding the contract method 0x9ca5a3d0.
//
// Solidity: function getMaxExternalRatioBP() view returns(uint256)
func (_Lido *LidoCaller) GetMaxExternalRatioBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getMaxExternalRatioBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxExternalRatioBP is a free data retrieval call binding the contract method 0x9ca5a3d0.
//
// Solidity: function getMaxExternalRatioBP() view returns(uint256)
func (_Lido *LidoSession) GetMaxExternalRatioBP() (*big.Int, error) {
	return _Lido.Contract.GetMaxExternalRatioBP(&_Lido.CallOpts)
}

// GetMaxExternalRatioBP is a free data retrieval call binding the contract method 0x9ca5a3d0.
//
// Solidity: function getMaxExternalRatioBP() view returns(uint256)
func (_Lido *LidoCallerSession) GetMaxExternalRatioBP() (*big.Int, error) {
	return _Lido.Contract.GetMaxExternalRatioBP(&_Lido.CallOpts)
}

// GetMaxMintableExternalShares is a free data retrieval call binding the contract method 0x8831f09e.
//
// Solidity: function getMaxMintableExternalShares() view returns(uint256)
func (_Lido *LidoCaller) GetMaxMintableExternalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getMaxMintableExternalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMintableExternalShares is a free data retrieval call binding the contract method 0x8831f09e.
//
// Solidity: function getMaxMintableExternalShares() view returns(uint256)
func (_Lido *LidoSession) GetMaxMintableExternalShares() (*big.Int, error) {
	return _Lido.Contract.GetMaxMintableExternalShares(&_Lido.CallOpts)
}

// GetMaxMintableExternalShares is a free data retrieval call binding the contract method 0x8831f09e.
//
// Solidity: function getMaxMintableExternalShares() view returns(uint256)
func (_Lido *LidoCallerSession) GetMaxMintableExternalShares() (*big.Int, error) {
	return _Lido.Contract.GetMaxMintableExternalShares(&_Lido.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetPooledEthByShares(&_Lido.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetPooledEthByShares(&_Lido.CallOpts, _sharesAmount)
}

// GetPooledEthBySharesRoundUp is a free data retrieval call binding the contract method 0x0103349c.
//
// Solidity: function getPooledEthBySharesRoundUp(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoCaller) GetPooledEthBySharesRoundUp(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getPooledEthBySharesRoundUp", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthBySharesRoundUp is a free data retrieval call binding the contract method 0x0103349c.
//
// Solidity: function getPooledEthBySharesRoundUp(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoSession) GetPooledEthBySharesRoundUp(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetPooledEthBySharesRoundUp(&_Lido.CallOpts, _sharesAmount)
}

// GetPooledEthBySharesRoundUp is a free data retrieval call binding the contract method 0x0103349c.
//
// Solidity: function getPooledEthBySharesRoundUp(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoCallerSession) GetPooledEthBySharesRoundUp(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetPooledEthBySharesRoundUp(&_Lido.CallOpts, _sharesAmount)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Lido *LidoCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Lido *LidoSession) GetRecoveryVault() (common.Address, error) {
	return _Lido.Contract.GetRecoveryVault(&_Lido.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Lido *LidoCallerSession) GetRecoveryVault() (common.Address, error) {
	return _Lido.Contract.GetRecoveryVault(&_Lido.CallOpts)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lido *LidoCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lido *LidoSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetSharesByPooledEth(&_Lido.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lido *LidoCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetSharesByPooledEth(&_Lido.CallOpts, _ethAmount)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused_, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_Lido *LidoCaller) GetStakeLimitFullInfo(opts *bind.CallOpts) (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getStakeLimitFullInfo")

	outstruct := new(struct {
		IsStakingPaused           bool
		IsStakingLimitSet         bool
		CurrentStakeLimit         *big.Int
		MaxStakeLimit             *big.Int
		MaxStakeLimitGrowthBlocks *big.Int
		PrevStakeLimit            *big.Int
		PrevStakeBlockNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStakingPaused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsStakingLimitSet = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentStakeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimitGrowthBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeLimit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused_, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_Lido *LidoSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _Lido.Contract.GetStakeLimitFullInfo(&_Lido.CallOpts)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused_, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_Lido *LidoCallerSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _Lido.Contract.GetStakeLimitFullInfo(&_Lido.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_Lido *LidoCaller) GetTotalELRewardsCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTotalELRewardsCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_Lido *LidoSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _Lido.Contract.GetTotalELRewardsCollected(&_Lido.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_Lido *LidoCallerSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _Lido.Contract.GetTotalELRewardsCollected(&_Lido.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lido *LidoCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lido *LidoSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lido.Contract.GetTotalPooledEther(&_Lido.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lido *LidoCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lido.Contract.GetTotalPooledEther(&_Lido.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lido *LidoCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lido *LidoSession) GetTotalShares() (*big.Int, error) {
	return _Lido.Contract.GetTotalShares(&_Lido.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lido *LidoCallerSession) GetTotalShares() (*big.Int, error) {
	return _Lido.Contract.GetTotalShares(&_Lido.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_Lido *LidoCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_Lido *LidoSession) GetTreasury() (common.Address, error) {
	return _Lido.Contract.GetTreasury(&_Lido.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_Lido *LidoCallerSession) GetTreasury() (common.Address, error) {
	return _Lido.Contract.GetTreasury(&_Lido.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Lido *LidoCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Lido *LidoSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Lido.Contract.GetWithdrawalCredentials(&_Lido.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Lido *LidoCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Lido.Contract.GetWithdrawalCredentials(&_Lido.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Lido *LidoCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Lido *LidoSession) HasInitialized() (bool, error) {
	return _Lido.Contract.HasInitialized(&_Lido.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Lido *LidoCallerSession) HasInitialized() (bool, error) {
	return _Lido.Contract.HasInitialized(&_Lido.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Lido *LidoCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Lido *LidoSession) IsPetrified() (bool, error) {
	return _Lido.Contract.IsPetrified(&_Lido.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Lido *LidoCallerSession) IsPetrified() (bool, error) {
	return _Lido.Contract.IsPetrified(&_Lido.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_Lido *LidoCaller) IsStakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "isStakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_Lido *LidoSession) IsStakingPaused() (bool, error) {
	return _Lido.Contract.IsStakingPaused(&_Lido.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_Lido *LidoCallerSession) IsStakingPaused() (bool, error) {
	return _Lido.Contract.IsStakingPaused(&_Lido.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_Lido *LidoCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "isStopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_Lido *LidoSession) IsStopped() (bool, error) {
	return _Lido.Contract.IsStopped(&_Lido.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_Lido *LidoCallerSession) IsStopped() (bool, error) {
	return _Lido.Contract.IsStopped(&_Lido.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Lido *LidoCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Lido *LidoSession) Kernel() (common.Address, error) {
	return _Lido.Contract.Kernel(&_Lido.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Lido *LidoCallerSession) Kernel() (common.Address, error) {
	return _Lido.Contract.Kernel(&_Lido.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Lido *LidoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Lido *LidoSession) Name() (string, error) {
	return _Lido.Contract.Name(&_Lido.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Lido *LidoCallerSession) Name() (string, error) {
	return _Lido.Contract.Name(&_Lido.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Lido *LidoCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Lido *LidoSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Lido.Contract.Nonces(&_Lido.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Lido *LidoCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Lido.Contract.Nonces(&_Lido.CallOpts, owner)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lido *LidoCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lido *LidoSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.SharesOf(&_Lido.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lido *LidoCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.SharesOf(&_Lido.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Lido *LidoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Lido *LidoSession) Symbol() (string, error) {
	return _Lido.Contract.Symbol(&_Lido.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Lido *LidoCallerSession) Symbol() (string, error) {
	return _Lido.Contract.Symbol(&_Lido.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lido *LidoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lido *LidoSession) TotalSupply() (*big.Int, error) {
	return _Lido.Contract.TotalSupply(&_Lido.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lido *LidoCallerSession) TotalSupply() (*big.Int, error) {
	return _Lido.Contract.TotalSupply(&_Lido.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Lido *LidoTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Lido *LidoSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Approve(&_Lido.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Lido *LidoTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Approve(&_Lido.TransactOpts, _spender, _amount)
}

// BurnExternalShares is a paid mutator transaction binding the contract method 0x72e62e56.
//
// Solidity: function burnExternalShares(uint256 _amountOfShares) returns()
func (_Lido *LidoTransactor) BurnExternalShares(opts *bind.TransactOpts, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "burnExternalShares", _amountOfShares)
}

// BurnExternalShares is a paid mutator transaction binding the contract method 0x72e62e56.
//
// Solidity: function burnExternalShares(uint256 _amountOfShares) returns()
func (_Lido *LidoSession) BurnExternalShares(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.BurnExternalShares(&_Lido.TransactOpts, _amountOfShares)
}

// BurnExternalShares is a paid mutator transaction binding the contract method 0x72e62e56.
//
// Solidity: function burnExternalShares(uint256 _amountOfShares) returns()
func (_Lido *LidoTransactorSession) BurnExternalShares(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.BurnExternalShares(&_Lido.TransactOpts, _amountOfShares)
}

// BurnShares is a paid mutator transaction binding the contract method 0x853c637d.
//
// Solidity: function burnShares(uint256 _amountOfShares) returns()
func (_Lido *LidoTransactor) BurnShares(opts *bind.TransactOpts, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "burnShares", _amountOfShares)
}

// BurnShares is a paid mutator transaction binding the contract method 0x853c637d.
//
// Solidity: function burnShares(uint256 _amountOfShares) returns()
func (_Lido *LidoSession) BurnShares(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.BurnShares(&_Lido.TransactOpts, _amountOfShares)
}

// BurnShares is a paid mutator transaction binding the contract method 0x853c637d.
//
// Solidity: function burnShares(uint256 _amountOfShares) returns()
func (_Lido *LidoTransactorSession) BurnShares(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.BurnShares(&_Lido.TransactOpts, _amountOfShares)
}

// CollectRewardsAndProcessWithdrawals is a paid mutator transaction binding the contract method 0x9271e3e6.
//
// Solidity: function collectRewardsAndProcessWithdrawals(uint256 _reportTimestamp, uint256 _reportClBalance, uint256 _principalCLBalance, uint256 _withdrawalsToWithdraw, uint256 _elRewardsToWithdraw, uint256 _lastWithdrawalRequestToFinalize, uint256 _withdrawalsShareRate, uint256 _etherToLockOnWithdrawalQueue) returns()
func (_Lido *LidoTransactor) CollectRewardsAndProcessWithdrawals(opts *bind.TransactOpts, _reportTimestamp *big.Int, _reportClBalance *big.Int, _principalCLBalance *big.Int, _withdrawalsToWithdraw *big.Int, _elRewardsToWithdraw *big.Int, _lastWithdrawalRequestToFinalize *big.Int, _withdrawalsShareRate *big.Int, _etherToLockOnWithdrawalQueue *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "collectRewardsAndProcessWithdrawals", _reportTimestamp, _reportClBalance, _principalCLBalance, _withdrawalsToWithdraw, _elRewardsToWithdraw, _lastWithdrawalRequestToFinalize, _withdrawalsShareRate, _etherToLockOnWithdrawalQueue)
}

// CollectRewardsAndProcessWithdrawals is a paid mutator transaction binding the contract method 0x9271e3e6.
//
// Solidity: function collectRewardsAndProcessWithdrawals(uint256 _reportTimestamp, uint256 _reportClBalance, uint256 _principalCLBalance, uint256 _withdrawalsToWithdraw, uint256 _elRewardsToWithdraw, uint256 _lastWithdrawalRequestToFinalize, uint256 _withdrawalsShareRate, uint256 _etherToLockOnWithdrawalQueue) returns()
func (_Lido *LidoSession) CollectRewardsAndProcessWithdrawals(_reportTimestamp *big.Int, _reportClBalance *big.Int, _principalCLBalance *big.Int, _withdrawalsToWithdraw *big.Int, _elRewardsToWithdraw *big.Int, _lastWithdrawalRequestToFinalize *big.Int, _withdrawalsShareRate *big.Int, _etherToLockOnWithdrawalQueue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.CollectRewardsAndProcessWithdrawals(&_Lido.TransactOpts, _reportTimestamp, _reportClBalance, _principalCLBalance, _withdrawalsToWithdraw, _elRewardsToWithdraw, _lastWithdrawalRequestToFinalize, _withdrawalsShareRate, _etherToLockOnWithdrawalQueue)
}

// CollectRewardsAndProcessWithdrawals is a paid mutator transaction binding the contract method 0x9271e3e6.
//
// Solidity: function collectRewardsAndProcessWithdrawals(uint256 _reportTimestamp, uint256 _reportClBalance, uint256 _principalCLBalance, uint256 _withdrawalsToWithdraw, uint256 _elRewardsToWithdraw, uint256 _lastWithdrawalRequestToFinalize, uint256 _withdrawalsShareRate, uint256 _etherToLockOnWithdrawalQueue) returns()
func (_Lido *LidoTransactorSession) CollectRewardsAndProcessWithdrawals(_reportTimestamp *big.Int, _reportClBalance *big.Int, _principalCLBalance *big.Int, _withdrawalsToWithdraw *big.Int, _elRewardsToWithdraw *big.Int, _lastWithdrawalRequestToFinalize *big.Int, _withdrawalsShareRate *big.Int, _etherToLockOnWithdrawalQueue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.CollectRewardsAndProcessWithdrawals(&_Lido.TransactOpts, _reportTimestamp, _reportClBalance, _principalCLBalance, _withdrawalsToWithdraw, _elRewardsToWithdraw, _lastWithdrawalRequestToFinalize, _withdrawalsShareRate, _etherToLockOnWithdrawalQueue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Lido *LidoTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Lido *LidoSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.DecreaseAllowance(&_Lido.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Lido *LidoTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.DecreaseAllowance(&_Lido.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_Lido *LidoTransactor) Deposit(opts *bind.TransactOpts, _maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "deposit", _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_Lido *LidoSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Lido.Contract.Deposit(&_Lido.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_Lido *LidoTransactorSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Lido.Contract.Deposit(&_Lido.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// EmitTokenRebase is a paid mutator transaction binding the contract method 0x1b250097.
//
// Solidity: function emitTokenRebase(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _preTotalShares, uint256 _preTotalEther, uint256 _postTotalShares, uint256 _postTotalEther, uint256 _postInternalShares, uint256 _postInternalEther, uint256 _sharesMintedAsFees) returns()
func (_Lido *LidoTransactor) EmitTokenRebase(opts *bind.TransactOpts, _reportTimestamp *big.Int, _timeElapsed *big.Int, _preTotalShares *big.Int, _preTotalEther *big.Int, _postTotalShares *big.Int, _postTotalEther *big.Int, _postInternalShares *big.Int, _postInternalEther *big.Int, _sharesMintedAsFees *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "emitTokenRebase", _reportTimestamp, _timeElapsed, _preTotalShares, _preTotalEther, _postTotalShares, _postTotalEther, _postInternalShares, _postInternalEther, _sharesMintedAsFees)
}

// EmitTokenRebase is a paid mutator transaction binding the contract method 0x1b250097.
//
// Solidity: function emitTokenRebase(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _preTotalShares, uint256 _preTotalEther, uint256 _postTotalShares, uint256 _postTotalEther, uint256 _postInternalShares, uint256 _postInternalEther, uint256 _sharesMintedAsFees) returns()
func (_Lido *LidoSession) EmitTokenRebase(_reportTimestamp *big.Int, _timeElapsed *big.Int, _preTotalShares *big.Int, _preTotalEther *big.Int, _postTotalShares *big.Int, _postTotalEther *big.Int, _postInternalShares *big.Int, _postInternalEther *big.Int, _sharesMintedAsFees *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.EmitTokenRebase(&_Lido.TransactOpts, _reportTimestamp, _timeElapsed, _preTotalShares, _preTotalEther, _postTotalShares, _postTotalEther, _postInternalShares, _postInternalEther, _sharesMintedAsFees)
}

// EmitTokenRebase is a paid mutator transaction binding the contract method 0x1b250097.
//
// Solidity: function emitTokenRebase(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _preTotalShares, uint256 _preTotalEther, uint256 _postTotalShares, uint256 _postTotalEther, uint256 _postInternalShares, uint256 _postInternalEther, uint256 _sharesMintedAsFees) returns()
func (_Lido *LidoTransactorSession) EmitTokenRebase(_reportTimestamp *big.Int, _timeElapsed *big.Int, _preTotalShares *big.Int, _preTotalEther *big.Int, _postTotalShares *big.Int, _postTotalEther *big.Int, _postInternalShares *big.Int, _postInternalEther *big.Int, _sharesMintedAsFees *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.EmitTokenRebase(&_Lido.TransactOpts, _reportTimestamp, _timeElapsed, _preTotalShares, _preTotalEther, _postTotalShares, _postTotalEther, _postInternalShares, _postInternalEther, _sharesMintedAsFees)
}

// FinalizeUpgradeV3 is a paid mutator transaction binding the contract method 0x878a209c.
//
// Solidity: function finalizeUpgrade_v3(address _oldBurner, address[] _contractsWithBurnerAllowances, uint256 _initialMaxExternalRatioBP) returns()
func (_Lido *LidoTransactor) FinalizeUpgradeV3(opts *bind.TransactOpts, _oldBurner common.Address, _contractsWithBurnerAllowances []common.Address, _initialMaxExternalRatioBP *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "finalizeUpgrade_v3", _oldBurner, _contractsWithBurnerAllowances, _initialMaxExternalRatioBP)
}

// FinalizeUpgradeV3 is a paid mutator transaction binding the contract method 0x878a209c.
//
// Solidity: function finalizeUpgrade_v3(address _oldBurner, address[] _contractsWithBurnerAllowances, uint256 _initialMaxExternalRatioBP) returns()
func (_Lido *LidoSession) FinalizeUpgradeV3(_oldBurner common.Address, _contractsWithBurnerAllowances []common.Address, _initialMaxExternalRatioBP *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.FinalizeUpgradeV3(&_Lido.TransactOpts, _oldBurner, _contractsWithBurnerAllowances, _initialMaxExternalRatioBP)
}

// FinalizeUpgradeV3 is a paid mutator transaction binding the contract method 0x878a209c.
//
// Solidity: function finalizeUpgrade_v3(address _oldBurner, address[] _contractsWithBurnerAllowances, uint256 _initialMaxExternalRatioBP) returns()
func (_Lido *LidoTransactorSession) FinalizeUpgradeV3(_oldBurner common.Address, _contractsWithBurnerAllowances []common.Address, _initialMaxExternalRatioBP *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.FinalizeUpgradeV3(&_Lido.TransactOpts, _oldBurner, _contractsWithBurnerAllowances, _initialMaxExternalRatioBP)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Lido *LidoTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Lido *LidoSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.IncreaseAllowance(&_Lido.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Lido *LidoTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.IncreaseAllowance(&_Lido.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_Lido *LidoTransactor) Initialize(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "initialize", _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_Lido *LidoSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Initialize(&_Lido.TransactOpts, _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_Lido *LidoTransactorSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Initialize(&_Lido.TransactOpts, _lidoLocator, _eip712StETH)
}

// InternalizeExternalBadDebt is a paid mutator transaction binding the contract method 0x648c51e7.
//
// Solidity: function internalizeExternalBadDebt(uint256 _amountOfShares) returns()
func (_Lido *LidoTransactor) InternalizeExternalBadDebt(opts *bind.TransactOpts, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "internalizeExternalBadDebt", _amountOfShares)
}

// InternalizeExternalBadDebt is a paid mutator transaction binding the contract method 0x648c51e7.
//
// Solidity: function internalizeExternalBadDebt(uint256 _amountOfShares) returns()
func (_Lido *LidoSession) InternalizeExternalBadDebt(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.InternalizeExternalBadDebt(&_Lido.TransactOpts, _amountOfShares)
}

// InternalizeExternalBadDebt is a paid mutator transaction binding the contract method 0x648c51e7.
//
// Solidity: function internalizeExternalBadDebt(uint256 _amountOfShares) returns()
func (_Lido *LidoTransactorSession) InternalizeExternalBadDebt(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.InternalizeExternalBadDebt(&_Lido.TransactOpts, _amountOfShares)
}

// MintExternalShares is a paid mutator transaction binding the contract method 0x06f187a4.
//
// Solidity: function mintExternalShares(address _recipient, uint256 _amountOfShares) returns()
func (_Lido *LidoTransactor) MintExternalShares(opts *bind.TransactOpts, _recipient common.Address, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "mintExternalShares", _recipient, _amountOfShares)
}

// MintExternalShares is a paid mutator transaction binding the contract method 0x06f187a4.
//
// Solidity: function mintExternalShares(address _recipient, uint256 _amountOfShares) returns()
func (_Lido *LidoSession) MintExternalShares(_recipient common.Address, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.MintExternalShares(&_Lido.TransactOpts, _recipient, _amountOfShares)
}

// MintExternalShares is a paid mutator transaction binding the contract method 0x06f187a4.
//
// Solidity: function mintExternalShares(address _recipient, uint256 _amountOfShares) returns()
func (_Lido *LidoTransactorSession) MintExternalShares(_recipient common.Address, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.MintExternalShares(&_Lido.TransactOpts, _recipient, _amountOfShares)
}

// MintShares is a paid mutator transaction binding the contract method 0x528c198a.
//
// Solidity: function mintShares(address _recipient, uint256 _amountOfShares) returns()
func (_Lido *LidoTransactor) MintShares(opts *bind.TransactOpts, _recipient common.Address, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "mintShares", _recipient, _amountOfShares)
}

// MintShares is a paid mutator transaction binding the contract method 0x528c198a.
//
// Solidity: function mintShares(address _recipient, uint256 _amountOfShares) returns()
func (_Lido *LidoSession) MintShares(_recipient common.Address, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.MintShares(&_Lido.TransactOpts, _recipient, _amountOfShares)
}

// MintShares is a paid mutator transaction binding the contract method 0x528c198a.
//
// Solidity: function mintShares(address _recipient, uint256 _amountOfShares) returns()
func (_Lido *LidoTransactorSession) MintShares(_recipient common.Address, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.MintShares(&_Lido.TransactOpts, _recipient, _amountOfShares)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Lido *LidoTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Lido *LidoSession) PauseStaking() (*types.Transaction, error) {
	return _Lido.Contract.PauseStaking(&_Lido.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Lido *LidoTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _Lido.Contract.PauseStaking(&_Lido.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Lido *LidoTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Lido *LidoSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Lido.Contract.Permit(&_Lido.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Lido *LidoTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Lido.Contract.Permit(&_Lido.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// ProcessClStateUpdate is a paid mutator transaction binding the contract method 0x603540cd.
//
// Solidity: function processClStateUpdate(uint256 _reportTimestamp, uint256 _preClValidators, uint256 _reportClValidators, uint256 _reportClBalance) returns()
func (_Lido *LidoTransactor) ProcessClStateUpdate(opts *bind.TransactOpts, _reportTimestamp *big.Int, _preClValidators *big.Int, _reportClValidators *big.Int, _reportClBalance *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "processClStateUpdate", _reportTimestamp, _preClValidators, _reportClValidators, _reportClBalance)
}

// ProcessClStateUpdate is a paid mutator transaction binding the contract method 0x603540cd.
//
// Solidity: function processClStateUpdate(uint256 _reportTimestamp, uint256 _preClValidators, uint256 _reportClValidators, uint256 _reportClBalance) returns()
func (_Lido *LidoSession) ProcessClStateUpdate(_reportTimestamp *big.Int, _preClValidators *big.Int, _reportClValidators *big.Int, _reportClBalance *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.ProcessClStateUpdate(&_Lido.TransactOpts, _reportTimestamp, _preClValidators, _reportClValidators, _reportClBalance)
}

// ProcessClStateUpdate is a paid mutator transaction binding the contract method 0x603540cd.
//
// Solidity: function processClStateUpdate(uint256 _reportTimestamp, uint256 _preClValidators, uint256 _reportClValidators, uint256 _reportClBalance) returns()
func (_Lido *LidoTransactorSession) ProcessClStateUpdate(_reportTimestamp *big.Int, _preClValidators *big.Int, _reportClValidators *big.Int, _reportClBalance *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.ProcessClStateUpdate(&_Lido.TransactOpts, _reportTimestamp, _preClValidators, _reportClValidators, _reportClBalance)
}

// RebalanceExternalEtherToInternal is a paid mutator transaction binding the contract method 0x7c8d9e38.
//
// Solidity: function rebalanceExternalEtherToInternal(uint256 _amountOfShares) payable returns()
func (_Lido *LidoTransactor) RebalanceExternalEtherToInternal(opts *bind.TransactOpts, _amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "rebalanceExternalEtherToInternal", _amountOfShares)
}

// RebalanceExternalEtherToInternal is a paid mutator transaction binding the contract method 0x7c8d9e38.
//
// Solidity: function rebalanceExternalEtherToInternal(uint256 _amountOfShares) payable returns()
func (_Lido *LidoSession) RebalanceExternalEtherToInternal(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.RebalanceExternalEtherToInternal(&_Lido.TransactOpts, _amountOfShares)
}

// RebalanceExternalEtherToInternal is a paid mutator transaction binding the contract method 0x7c8d9e38.
//
// Solidity: function rebalanceExternalEtherToInternal(uint256 _amountOfShares) payable returns()
func (_Lido *LidoTransactorSession) RebalanceExternalEtherToInternal(_amountOfShares *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.RebalanceExternalEtherToInternal(&_Lido.TransactOpts, _amountOfShares)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_Lido *LidoTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_Lido *LidoSession) ReceiveELRewards() (*types.Transaction, error) {
	return _Lido.Contract.ReceiveELRewards(&_Lido.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_Lido *LidoTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _Lido.Contract.ReceiveELRewards(&_Lido.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_Lido *LidoTransactor) ReceiveWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "receiveWithdrawals")
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_Lido *LidoSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _Lido.Contract.ReceiveWithdrawals(&_Lido.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_Lido *LidoTransactorSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _Lido.Contract.ReceiveWithdrawals(&_Lido.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_Lido *LidoTransactor) RemoveStakingLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "removeStakingLimit")
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_Lido *LidoSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _Lido.Contract.RemoveStakingLimit(&_Lido.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_Lido *LidoTransactorSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _Lido.Contract.RemoveStakingLimit(&_Lido.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Lido *LidoTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Lido *LidoSession) Resume() (*types.Transaction, error) {
	return _Lido.Contract.Resume(&_Lido.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Lido *LidoTransactorSession) Resume() (*types.Transaction, error) {
	return _Lido.Contract.Resume(&_Lido.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Lido *LidoTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Lido *LidoSession) ResumeStaking() (*types.Transaction, error) {
	return _Lido.Contract.ResumeStaking(&_Lido.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Lido *LidoTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _Lido.Contract.ResumeStaking(&_Lido.TransactOpts)
}

// SetMaxExternalRatioBP is a paid mutator transaction binding the contract method 0xf352e17e.
//
// Solidity: function setMaxExternalRatioBP(uint256 _maxExternalRatioBP) returns()
func (_Lido *LidoTransactor) SetMaxExternalRatioBP(opts *bind.TransactOpts, _maxExternalRatioBP *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setMaxExternalRatioBP", _maxExternalRatioBP)
}

// SetMaxExternalRatioBP is a paid mutator transaction binding the contract method 0xf352e17e.
//
// Solidity: function setMaxExternalRatioBP(uint256 _maxExternalRatioBP) returns()
func (_Lido *LidoSession) SetMaxExternalRatioBP(_maxExternalRatioBP *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.SetMaxExternalRatioBP(&_Lido.TransactOpts, _maxExternalRatioBP)
}

// SetMaxExternalRatioBP is a paid mutator transaction binding the contract method 0xf352e17e.
//
// Solidity: function setMaxExternalRatioBP(uint256 _maxExternalRatioBP) returns()
func (_Lido *LidoTransactorSession) SetMaxExternalRatioBP(_maxExternalRatioBP *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.SetMaxExternalRatioBP(&_Lido.TransactOpts, _maxExternalRatioBP)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_Lido *LidoTransactor) SetStakingLimit(opts *bind.TransactOpts, _maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setStakingLimit", _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_Lido *LidoSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.SetStakingLimit(&_Lido.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_Lido *LidoTransactorSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.SetStakingLimit(&_Lido.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Lido *LidoTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Lido *LidoSession) Stop() (*types.Transaction, error) {
	return _Lido.Contract.Stop(&_Lido.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Lido *LidoTransactorSession) Stop() (*types.Transaction, error) {
	return _Lido.Contract.Stop(&_Lido.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Submit(&_Lido.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Submit(&_Lido.TransactOpts, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Transfer(&_Lido.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Transfer(&_Lido.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferFrom(&_Lido.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferFrom(&_Lido.TransactOpts, _sender, _recipient, _amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoTransactor) TransferShares(opts *bind.TransactOpts, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferShares", _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferShares(&_Lido.TransactOpts, _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoTransactorSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferShares(&_Lido.TransactOpts, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoTransactor) TransferSharesFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferSharesFrom", _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferSharesFrom(&_Lido.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoTransactorSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferSharesFrom(&_Lido.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_Lido *LidoTransactor) TransferToVault(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferToVault", arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_Lido *LidoSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _Lido.Contract.TransferToVault(&_Lido.TransactOpts, arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_Lido *LidoTransactorSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _Lido.Contract.TransferToVault(&_Lido.TransactOpts, arg0)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_Lido *LidoTransactor) UnsafeChangeDepositedValidators(opts *bind.TransactOpts, _newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "unsafeChangeDepositedValidators", _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_Lido *LidoSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.UnsafeChangeDepositedValidators(&_Lido.TransactOpts, _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_Lido *LidoTransactorSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.UnsafeChangeDepositedValidators(&_Lido.TransactOpts, _newDepositedValidators)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lido *LidoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Lido.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lido *LidoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lido.Contract.Fallback(&_Lido.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lido *LidoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lido.Contract.Fallback(&_Lido.TransactOpts, calldata)
}

// LidoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lido contract.
type LidoApprovalIterator struct {
	Event *LidoApproval // Event containing the contract specifics and raw log

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
func (it *LidoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoApproval)
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
		it.Event = new(LidoApproval)
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
func (it *LidoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoApproval represents a Approval event raised by the Lido contract.
type LidoApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lido *LidoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LidoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LidoApprovalIterator{contract: _Lido.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lido *LidoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LidoApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoApproval)
				if err := _Lido.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Lido *LidoFilterer) ParseApproval(log types.Log) (*LidoApproval, error) {
	event := new(LidoApproval)
	if err := _Lido.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoCLValidatorsUpdatedIterator is returned from FilterCLValidatorsUpdated and is used to iterate over the raw logs and unpacked data for CLValidatorsUpdated events raised by the Lido contract.
type LidoCLValidatorsUpdatedIterator struct {
	Event *LidoCLValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *LidoCLValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoCLValidatorsUpdated)
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
		it.Event = new(LidoCLValidatorsUpdated)
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
func (it *LidoCLValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoCLValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoCLValidatorsUpdated represents a CLValidatorsUpdated event raised by the Lido contract.
type LidoCLValidatorsUpdated struct {
	ReportTimestamp  *big.Int
	PreCLValidators  *big.Int
	PostCLValidators *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCLValidatorsUpdated is a free log retrieval operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_Lido *LidoFilterer) FilterCLValidatorsUpdated(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*LidoCLValidatorsUpdatedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &LidoCLValidatorsUpdatedIterator{contract: _Lido.contract, event: "CLValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchCLValidatorsUpdated is a free log subscription operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_Lido *LidoFilterer) WatchCLValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *LidoCLValidatorsUpdated, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoCLValidatorsUpdated)
				if err := _Lido.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
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

// ParseCLValidatorsUpdated is a log parse operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_Lido *LidoFilterer) ParseCLValidatorsUpdated(log types.Log) (*LidoCLValidatorsUpdated, error) {
	event := new(LidoCLValidatorsUpdated)
	if err := _Lido.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Lido contract.
type LidoContractVersionSetIterator struct {
	Event *LidoContractVersionSet // Event containing the contract specifics and raw log

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
func (it *LidoContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoContractVersionSet)
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
		it.Event = new(LidoContractVersionSet)
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
func (it *LidoContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoContractVersionSet represents a ContractVersionSet event raised by the Lido contract.
type LidoContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Lido *LidoFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*LidoContractVersionSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &LidoContractVersionSetIterator{contract: _Lido.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Lido *LidoFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *LidoContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoContractVersionSet)
				if err := _Lido.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Lido *LidoFilterer) ParseContractVersionSet(log types.Log) (*LidoContractVersionSet, error) {
	event := new(LidoContractVersionSet)
	if err := _Lido.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoDepositedValidatorsChangedIterator is returned from FilterDepositedValidatorsChanged and is used to iterate over the raw logs and unpacked data for DepositedValidatorsChanged events raised by the Lido contract.
type LidoDepositedValidatorsChangedIterator struct {
	Event *LidoDepositedValidatorsChanged // Event containing the contract specifics and raw log

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
func (it *LidoDepositedValidatorsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoDepositedValidatorsChanged)
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
		it.Event = new(LidoDepositedValidatorsChanged)
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
func (it *LidoDepositedValidatorsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoDepositedValidatorsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoDepositedValidatorsChanged represents a DepositedValidatorsChanged event raised by the Lido contract.
type LidoDepositedValidatorsChanged struct {
	DepositedValidators *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositedValidatorsChanged is a free log retrieval operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_Lido *LidoFilterer) FilterDepositedValidatorsChanged(opts *bind.FilterOpts) (*LidoDepositedValidatorsChangedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return &LidoDepositedValidatorsChangedIterator{contract: _Lido.contract, event: "DepositedValidatorsChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedValidatorsChanged is a free log subscription operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_Lido *LidoFilterer) WatchDepositedValidatorsChanged(opts *bind.WatchOpts, sink chan<- *LidoDepositedValidatorsChanged) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoDepositedValidatorsChanged)
				if err := _Lido.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
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

// ParseDepositedValidatorsChanged is a log parse operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_Lido *LidoFilterer) ParseDepositedValidatorsChanged(log types.Log) (*LidoDepositedValidatorsChanged, error) {
	event := new(LidoDepositedValidatorsChanged)
	if err := _Lido.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoEIP712StETHInitializedIterator is returned from FilterEIP712StETHInitialized and is used to iterate over the raw logs and unpacked data for EIP712StETHInitialized events raised by the Lido contract.
type LidoEIP712StETHInitializedIterator struct {
	Event *LidoEIP712StETHInitialized // Event containing the contract specifics and raw log

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
func (it *LidoEIP712StETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoEIP712StETHInitialized)
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
		it.Event = new(LidoEIP712StETHInitialized)
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
func (it *LidoEIP712StETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoEIP712StETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoEIP712StETHInitialized represents a EIP712StETHInitialized event raised by the Lido contract.
type LidoEIP712StETHInitialized struct {
	Eip712StETH common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEIP712StETHInitialized is a free log retrieval operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_Lido *LidoFilterer) FilterEIP712StETHInitialized(opts *bind.FilterOpts) (*LidoEIP712StETHInitializedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return &LidoEIP712StETHInitializedIterator{contract: _Lido.contract, event: "EIP712StETHInitialized", logs: logs, sub: sub}, nil
}

// WatchEIP712StETHInitialized is a free log subscription operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_Lido *LidoFilterer) WatchEIP712StETHInitialized(opts *bind.WatchOpts, sink chan<- *LidoEIP712StETHInitialized) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoEIP712StETHInitialized)
				if err := _Lido.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
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

// ParseEIP712StETHInitialized is a log parse operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_Lido *LidoFilterer) ParseEIP712StETHInitialized(log types.Log) (*LidoEIP712StETHInitialized, error) {
	event := new(LidoEIP712StETHInitialized)
	if err := _Lido.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoELRewardsReceivedIterator is returned from FilterELRewardsReceived and is used to iterate over the raw logs and unpacked data for ELRewardsReceived events raised by the Lido contract.
type LidoELRewardsReceivedIterator struct {
	Event *LidoELRewardsReceived // Event containing the contract specifics and raw log

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
func (it *LidoELRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoELRewardsReceived)
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
		it.Event = new(LidoELRewardsReceived)
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
func (it *LidoELRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoELRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoELRewardsReceived represents a ELRewardsReceived event raised by the Lido contract.
type LidoELRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceived is a free log retrieval operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_Lido *LidoFilterer) FilterELRewardsReceived(opts *bind.FilterOpts) (*LidoELRewardsReceivedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &LidoELRewardsReceivedIterator{contract: _Lido.contract, event: "ELRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceived is a free log subscription operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_Lido *LidoFilterer) WatchELRewardsReceived(opts *bind.WatchOpts, sink chan<- *LidoELRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoELRewardsReceived)
				if err := _Lido.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
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

// ParseELRewardsReceived is a log parse operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_Lido *LidoFilterer) ParseELRewardsReceived(log types.Log) (*LidoELRewardsReceived, error) {
	event := new(LidoELRewardsReceived)
	if err := _Lido.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoETHDistributedIterator is returned from FilterETHDistributed and is used to iterate over the raw logs and unpacked data for ETHDistributed events raised by the Lido contract.
type LidoETHDistributedIterator struct {
	Event *LidoETHDistributed // Event containing the contract specifics and raw log

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
func (it *LidoETHDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoETHDistributed)
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
		it.Event = new(LidoETHDistributed)
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
func (it *LidoETHDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoETHDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoETHDistributed represents a ETHDistributed event raised by the Lido contract.
type LidoETHDistributed struct {
	ReportTimestamp                *big.Int
	PreCLBalance                   *big.Int
	PostCLBalance                  *big.Int
	WithdrawalsWithdrawn           *big.Int
	ExecutionLayerRewardsWithdrawn *big.Int
	PostBufferedEther              *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterETHDistributed is a free log retrieval operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_Lido *LidoFilterer) FilterETHDistributed(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*LidoETHDistributedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &LidoETHDistributedIterator{contract: _Lido.contract, event: "ETHDistributed", logs: logs, sub: sub}, nil
}

// WatchETHDistributed is a free log subscription operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_Lido *LidoFilterer) WatchETHDistributed(opts *bind.WatchOpts, sink chan<- *LidoETHDistributed, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoETHDistributed)
				if err := _Lido.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
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

// ParseETHDistributed is a log parse operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_Lido *LidoFilterer) ParseETHDistributed(log types.Log) (*LidoETHDistributed, error) {
	event := new(LidoETHDistributed)
	if err := _Lido.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoExternalBadDebtInternalizedIterator is returned from FilterExternalBadDebtInternalized and is used to iterate over the raw logs and unpacked data for ExternalBadDebtInternalized events raised by the Lido contract.
type LidoExternalBadDebtInternalizedIterator struct {
	Event *LidoExternalBadDebtInternalized // Event containing the contract specifics and raw log

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
func (it *LidoExternalBadDebtInternalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoExternalBadDebtInternalized)
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
		it.Event = new(LidoExternalBadDebtInternalized)
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
func (it *LidoExternalBadDebtInternalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoExternalBadDebtInternalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoExternalBadDebtInternalized represents a ExternalBadDebtInternalized event raised by the Lido contract.
type LidoExternalBadDebtInternalized struct {
	AmountOfShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExternalBadDebtInternalized is a free log retrieval operation binding the contract event 0x4e80196ef1285462b2c4ee20f88e18e58c59e405eec6fd51f5fd1d614bc98a7f.
//
// Solidity: event ExternalBadDebtInternalized(uint256 amountOfShares)
func (_Lido *LidoFilterer) FilterExternalBadDebtInternalized(opts *bind.FilterOpts) (*LidoExternalBadDebtInternalizedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ExternalBadDebtInternalized")
	if err != nil {
		return nil, err
	}
	return &LidoExternalBadDebtInternalizedIterator{contract: _Lido.contract, event: "ExternalBadDebtInternalized", logs: logs, sub: sub}, nil
}

// WatchExternalBadDebtInternalized is a free log subscription operation binding the contract event 0x4e80196ef1285462b2c4ee20f88e18e58c59e405eec6fd51f5fd1d614bc98a7f.
//
// Solidity: event ExternalBadDebtInternalized(uint256 amountOfShares)
func (_Lido *LidoFilterer) WatchExternalBadDebtInternalized(opts *bind.WatchOpts, sink chan<- *LidoExternalBadDebtInternalized) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ExternalBadDebtInternalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoExternalBadDebtInternalized)
				if err := _Lido.contract.UnpackLog(event, "ExternalBadDebtInternalized", log); err != nil {
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

// ParseExternalBadDebtInternalized is a log parse operation binding the contract event 0x4e80196ef1285462b2c4ee20f88e18e58c59e405eec6fd51f5fd1d614bc98a7f.
//
// Solidity: event ExternalBadDebtInternalized(uint256 amountOfShares)
func (_Lido *LidoFilterer) ParseExternalBadDebtInternalized(log types.Log) (*LidoExternalBadDebtInternalized, error) {
	event := new(LidoExternalBadDebtInternalized)
	if err := _Lido.contract.UnpackLog(event, "ExternalBadDebtInternalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoExternalEtherTransferredToBufferIterator is returned from FilterExternalEtherTransferredToBuffer and is used to iterate over the raw logs and unpacked data for ExternalEtherTransferredToBuffer events raised by the Lido contract.
type LidoExternalEtherTransferredToBufferIterator struct {
	Event *LidoExternalEtherTransferredToBuffer // Event containing the contract specifics and raw log

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
func (it *LidoExternalEtherTransferredToBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoExternalEtherTransferredToBuffer)
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
		it.Event = new(LidoExternalEtherTransferredToBuffer)
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
func (it *LidoExternalEtherTransferredToBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoExternalEtherTransferredToBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoExternalEtherTransferredToBuffer represents a ExternalEtherTransferredToBuffer event raised by the Lido contract.
type LidoExternalEtherTransferredToBuffer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExternalEtherTransferredToBuffer is a free log retrieval operation binding the contract event 0x4ee34277c93491eeca655ad5c42ae1c193a5719e1c8837df9058af7696817cce.
//
// Solidity: event ExternalEtherTransferredToBuffer(uint256 amount)
func (_Lido *LidoFilterer) FilterExternalEtherTransferredToBuffer(opts *bind.FilterOpts) (*LidoExternalEtherTransferredToBufferIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ExternalEtherTransferredToBuffer")
	if err != nil {
		return nil, err
	}
	return &LidoExternalEtherTransferredToBufferIterator{contract: _Lido.contract, event: "ExternalEtherTransferredToBuffer", logs: logs, sub: sub}, nil
}

// WatchExternalEtherTransferredToBuffer is a free log subscription operation binding the contract event 0x4ee34277c93491eeca655ad5c42ae1c193a5719e1c8837df9058af7696817cce.
//
// Solidity: event ExternalEtherTransferredToBuffer(uint256 amount)
func (_Lido *LidoFilterer) WatchExternalEtherTransferredToBuffer(opts *bind.WatchOpts, sink chan<- *LidoExternalEtherTransferredToBuffer) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ExternalEtherTransferredToBuffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoExternalEtherTransferredToBuffer)
				if err := _Lido.contract.UnpackLog(event, "ExternalEtherTransferredToBuffer", log); err != nil {
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

// ParseExternalEtherTransferredToBuffer is a log parse operation binding the contract event 0x4ee34277c93491eeca655ad5c42ae1c193a5719e1c8837df9058af7696817cce.
//
// Solidity: event ExternalEtherTransferredToBuffer(uint256 amount)
func (_Lido *LidoFilterer) ParseExternalEtherTransferredToBuffer(log types.Log) (*LidoExternalEtherTransferredToBuffer, error) {
	event := new(LidoExternalEtherTransferredToBuffer)
	if err := _Lido.contract.UnpackLog(event, "ExternalEtherTransferredToBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoExternalSharesBurntIterator is returned from FilterExternalSharesBurnt and is used to iterate over the raw logs and unpacked data for ExternalSharesBurnt events raised by the Lido contract.
type LidoExternalSharesBurntIterator struct {
	Event *LidoExternalSharesBurnt // Event containing the contract specifics and raw log

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
func (it *LidoExternalSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoExternalSharesBurnt)
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
		it.Event = new(LidoExternalSharesBurnt)
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
func (it *LidoExternalSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoExternalSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoExternalSharesBurnt represents a ExternalSharesBurnt event raised by the Lido contract.
type LidoExternalSharesBurnt struct {
	AmountOfShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExternalSharesBurnt is a free log retrieval operation binding the contract event 0xad21467656c56eb8c99f7916faa12f7a657a04d8802100acec62e92451ac5606.
//
// Solidity: event ExternalSharesBurnt(uint256 amountOfShares)
func (_Lido *LidoFilterer) FilterExternalSharesBurnt(opts *bind.FilterOpts) (*LidoExternalSharesBurntIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ExternalSharesBurnt")
	if err != nil {
		return nil, err
	}
	return &LidoExternalSharesBurntIterator{contract: _Lido.contract, event: "ExternalSharesBurnt", logs: logs, sub: sub}, nil
}

// WatchExternalSharesBurnt is a free log subscription operation binding the contract event 0xad21467656c56eb8c99f7916faa12f7a657a04d8802100acec62e92451ac5606.
//
// Solidity: event ExternalSharesBurnt(uint256 amountOfShares)
func (_Lido *LidoFilterer) WatchExternalSharesBurnt(opts *bind.WatchOpts, sink chan<- *LidoExternalSharesBurnt) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ExternalSharesBurnt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoExternalSharesBurnt)
				if err := _Lido.contract.UnpackLog(event, "ExternalSharesBurnt", log); err != nil {
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

// ParseExternalSharesBurnt is a log parse operation binding the contract event 0xad21467656c56eb8c99f7916faa12f7a657a04d8802100acec62e92451ac5606.
//
// Solidity: event ExternalSharesBurnt(uint256 amountOfShares)
func (_Lido *LidoFilterer) ParseExternalSharesBurnt(log types.Log) (*LidoExternalSharesBurnt, error) {
	event := new(LidoExternalSharesBurnt)
	if err := _Lido.contract.UnpackLog(event, "ExternalSharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoExternalSharesMintedIterator is returned from FilterExternalSharesMinted and is used to iterate over the raw logs and unpacked data for ExternalSharesMinted events raised by the Lido contract.
type LidoExternalSharesMintedIterator struct {
	Event *LidoExternalSharesMinted // Event containing the contract specifics and raw log

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
func (it *LidoExternalSharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoExternalSharesMinted)
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
		it.Event = new(LidoExternalSharesMinted)
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
func (it *LidoExternalSharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoExternalSharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoExternalSharesMinted represents a ExternalSharesMinted event raised by the Lido contract.
type LidoExternalSharesMinted struct {
	Receiver       common.Address
	AmountOfShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExternalSharesMinted is a free log retrieval operation binding the contract event 0xee473f96486a2f4b93ccb6729f121223e96975db6e0d6a5ef01f56477e3eab3b.
//
// Solidity: event ExternalSharesMinted(address indexed receiver, uint256 amountOfShares)
func (_Lido *LidoFilterer) FilterExternalSharesMinted(opts *bind.FilterOpts, receiver []common.Address) (*LidoExternalSharesMintedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ExternalSharesMinted", receiverRule)
	if err != nil {
		return nil, err
	}
	return &LidoExternalSharesMintedIterator{contract: _Lido.contract, event: "ExternalSharesMinted", logs: logs, sub: sub}, nil
}

// WatchExternalSharesMinted is a free log subscription operation binding the contract event 0xee473f96486a2f4b93ccb6729f121223e96975db6e0d6a5ef01f56477e3eab3b.
//
// Solidity: event ExternalSharesMinted(address indexed receiver, uint256 amountOfShares)
func (_Lido *LidoFilterer) WatchExternalSharesMinted(opts *bind.WatchOpts, sink chan<- *LidoExternalSharesMinted, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ExternalSharesMinted", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoExternalSharesMinted)
				if err := _Lido.contract.UnpackLog(event, "ExternalSharesMinted", log); err != nil {
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

// ParseExternalSharesMinted is a log parse operation binding the contract event 0xee473f96486a2f4b93ccb6729f121223e96975db6e0d6a5ef01f56477e3eab3b.
//
// Solidity: event ExternalSharesMinted(address indexed receiver, uint256 amountOfShares)
func (_Lido *LidoFilterer) ParseExternalSharesMinted(log types.Log) (*LidoExternalSharesMinted, error) {
	event := new(LidoExternalSharesMinted)
	if err := _Lido.contract.UnpackLog(event, "ExternalSharesMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoInternalShareRateUpdatedIterator is returned from FilterInternalShareRateUpdated and is used to iterate over the raw logs and unpacked data for InternalShareRateUpdated events raised by the Lido contract.
type LidoInternalShareRateUpdatedIterator struct {
	Event *LidoInternalShareRateUpdated // Event containing the contract specifics and raw log

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
func (it *LidoInternalShareRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoInternalShareRateUpdated)
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
		it.Event = new(LidoInternalShareRateUpdated)
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
func (it *LidoInternalShareRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoInternalShareRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoInternalShareRateUpdated represents a InternalShareRateUpdated event raised by the Lido contract.
type LidoInternalShareRateUpdated struct {
	ReportTimestamp    *big.Int
	PostInternalShares *big.Int
	PostInternalEther  *big.Int
	SharesMintedAsFees *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInternalShareRateUpdated is a free log retrieval operation binding the contract event 0xaf00d86be4cd299db16aa59803992e174fa88b67d81a0c7dd0148f9a75606a8d.
//
// Solidity: event InternalShareRateUpdated(uint256 indexed reportTimestamp, uint256 postInternalShares, uint256 postInternalEther, uint256 sharesMintedAsFees)
func (_Lido *LidoFilterer) FilterInternalShareRateUpdated(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*LidoInternalShareRateUpdatedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "InternalShareRateUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &LidoInternalShareRateUpdatedIterator{contract: _Lido.contract, event: "InternalShareRateUpdated", logs: logs, sub: sub}, nil
}

// WatchInternalShareRateUpdated is a free log subscription operation binding the contract event 0xaf00d86be4cd299db16aa59803992e174fa88b67d81a0c7dd0148f9a75606a8d.
//
// Solidity: event InternalShareRateUpdated(uint256 indexed reportTimestamp, uint256 postInternalShares, uint256 postInternalEther, uint256 sharesMintedAsFees)
func (_Lido *LidoFilterer) WatchInternalShareRateUpdated(opts *bind.WatchOpts, sink chan<- *LidoInternalShareRateUpdated, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "InternalShareRateUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoInternalShareRateUpdated)
				if err := _Lido.contract.UnpackLog(event, "InternalShareRateUpdated", log); err != nil {
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

// ParseInternalShareRateUpdated is a log parse operation binding the contract event 0xaf00d86be4cd299db16aa59803992e174fa88b67d81a0c7dd0148f9a75606a8d.
//
// Solidity: event InternalShareRateUpdated(uint256 indexed reportTimestamp, uint256 postInternalShares, uint256 postInternalEther, uint256 sharesMintedAsFees)
func (_Lido *LidoFilterer) ParseInternalShareRateUpdated(log types.Log) (*LidoInternalShareRateUpdated, error) {
	event := new(LidoInternalShareRateUpdated)
	if err := _Lido.contract.UnpackLog(event, "InternalShareRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoLidoLocatorSetIterator is returned from FilterLidoLocatorSet and is used to iterate over the raw logs and unpacked data for LidoLocatorSet events raised by the Lido contract.
type LidoLidoLocatorSetIterator struct {
	Event *LidoLidoLocatorSet // Event containing the contract specifics and raw log

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
func (it *LidoLidoLocatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoLidoLocatorSet)
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
		it.Event = new(LidoLidoLocatorSet)
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
func (it *LidoLidoLocatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoLidoLocatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoLidoLocatorSet represents a LidoLocatorSet event raised by the Lido contract.
type LidoLidoLocatorSet struct {
	LidoLocator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLidoLocatorSet is a free log retrieval operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_Lido *LidoFilterer) FilterLidoLocatorSet(opts *bind.FilterOpts) (*LidoLidoLocatorSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return &LidoLidoLocatorSetIterator{contract: _Lido.contract, event: "LidoLocatorSet", logs: logs, sub: sub}, nil
}

// WatchLidoLocatorSet is a free log subscription operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_Lido *LidoFilterer) WatchLidoLocatorSet(opts *bind.WatchOpts, sink chan<- *LidoLidoLocatorSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoLidoLocatorSet)
				if err := _Lido.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
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

// ParseLidoLocatorSet is a log parse operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_Lido *LidoFilterer) ParseLidoLocatorSet(log types.Log) (*LidoLidoLocatorSet, error) {
	event := new(LidoLidoLocatorSet)
	if err := _Lido.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoMaxExternalRatioBPSetIterator is returned from FilterMaxExternalRatioBPSet and is used to iterate over the raw logs and unpacked data for MaxExternalRatioBPSet events raised by the Lido contract.
type LidoMaxExternalRatioBPSetIterator struct {
	Event *LidoMaxExternalRatioBPSet // Event containing the contract specifics and raw log

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
func (it *LidoMaxExternalRatioBPSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoMaxExternalRatioBPSet)
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
		it.Event = new(LidoMaxExternalRatioBPSet)
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
func (it *LidoMaxExternalRatioBPSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoMaxExternalRatioBPSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoMaxExternalRatioBPSet represents a MaxExternalRatioBPSet event raised by the Lido contract.
type LidoMaxExternalRatioBPSet struct {
	MaxExternalRatioBP *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMaxExternalRatioBPSet is a free log retrieval operation binding the contract event 0x13c514ee70ee403f89bdf5ab83908edba92a77e3a61e19b774670c0a2cb2d7e6.
//
// Solidity: event MaxExternalRatioBPSet(uint256 maxExternalRatioBP)
func (_Lido *LidoFilterer) FilterMaxExternalRatioBPSet(opts *bind.FilterOpts) (*LidoMaxExternalRatioBPSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "MaxExternalRatioBPSet")
	if err != nil {
		return nil, err
	}
	return &LidoMaxExternalRatioBPSetIterator{contract: _Lido.contract, event: "MaxExternalRatioBPSet", logs: logs, sub: sub}, nil
}

// WatchMaxExternalRatioBPSet is a free log subscription operation binding the contract event 0x13c514ee70ee403f89bdf5ab83908edba92a77e3a61e19b774670c0a2cb2d7e6.
//
// Solidity: event MaxExternalRatioBPSet(uint256 maxExternalRatioBP)
func (_Lido *LidoFilterer) WatchMaxExternalRatioBPSet(opts *bind.WatchOpts, sink chan<- *LidoMaxExternalRatioBPSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "MaxExternalRatioBPSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoMaxExternalRatioBPSet)
				if err := _Lido.contract.UnpackLog(event, "MaxExternalRatioBPSet", log); err != nil {
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

// ParseMaxExternalRatioBPSet is a log parse operation binding the contract event 0x13c514ee70ee403f89bdf5ab83908edba92a77e3a61e19b774670c0a2cb2d7e6.
//
// Solidity: event MaxExternalRatioBPSet(uint256 maxExternalRatioBP)
func (_Lido *LidoFilterer) ParseMaxExternalRatioBPSet(log types.Log) (*LidoMaxExternalRatioBPSet, error) {
	event := new(LidoMaxExternalRatioBPSet)
	if err := _Lido.contract.UnpackLog(event, "MaxExternalRatioBPSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the Lido contract.
type LidoRecoverToVaultIterator struct {
	Event *LidoRecoverToVault // Event containing the contract specifics and raw log

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
func (it *LidoRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoRecoverToVault)
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
		it.Event = new(LidoRecoverToVault)
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
func (it *LidoRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoRecoverToVault represents a RecoverToVault event raised by the Lido contract.
type LidoRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Lido *LidoFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*LidoRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &LidoRecoverToVaultIterator{contract: _Lido.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Lido *LidoFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *LidoRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoRecoverToVault)
				if err := _Lido.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Lido *LidoFilterer) ParseRecoverToVault(log types.Log) (*LidoRecoverToVault, error) {
	event := new(LidoRecoverToVault)
	if err := _Lido.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Lido contract.
type LidoResumedIterator struct {
	Event *LidoResumed // Event containing the contract specifics and raw log

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
func (it *LidoResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoResumed)
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
		it.Event = new(LidoResumed)
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
func (it *LidoResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoResumed represents a Resumed event raised by the Lido contract.
type LidoResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Lido *LidoFilterer) FilterResumed(opts *bind.FilterOpts) (*LidoResumedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &LidoResumedIterator{contract: _Lido.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Lido *LidoFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *LidoResumed) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoResumed)
				if err := _Lido.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Lido *LidoFilterer) ParseResumed(log types.Log) (*LidoResumed, error) {
	event := new(LidoResumed)
	if err := _Lido.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the Lido contract.
type LidoScriptResultIterator struct {
	Event *LidoScriptResult // Event containing the contract specifics and raw log

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
func (it *LidoScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoScriptResult)
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
		it.Event = new(LidoScriptResult)
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
func (it *LidoScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoScriptResult represents a ScriptResult event raised by the Lido contract.
type LidoScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Lido *LidoFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*LidoScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &LidoScriptResultIterator{contract: _Lido.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Lido *LidoFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *LidoScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoScriptResult)
				if err := _Lido.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Lido *LidoFilterer) ParseScriptResult(log types.Log) (*LidoScriptResult, error) {
	event := new(LidoScriptResult)
	if err := _Lido.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoSharesBurntIterator is returned from FilterSharesBurnt and is used to iterate over the raw logs and unpacked data for SharesBurnt events raised by the Lido contract.
type LidoSharesBurntIterator struct {
	Event *LidoSharesBurnt // Event containing the contract specifics and raw log

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
func (it *LidoSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoSharesBurnt)
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
		it.Event = new(LidoSharesBurnt)
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
func (it *LidoSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoSharesBurnt represents a SharesBurnt event raised by the Lido contract.
type LidoSharesBurnt struct {
	Account               common.Address
	PreRebaseTokenAmount  *big.Int
	PostRebaseTokenAmount *big.Int
	SharesAmount          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSharesBurnt is a free log retrieval operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_Lido *LidoFilterer) FilterSharesBurnt(opts *bind.FilterOpts, account []common.Address) (*LidoSharesBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &LidoSharesBurntIterator{contract: _Lido.contract, event: "SharesBurnt", logs: logs, sub: sub}, nil
}

// WatchSharesBurnt is a free log subscription operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_Lido *LidoFilterer) WatchSharesBurnt(opts *bind.WatchOpts, sink chan<- *LidoSharesBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoSharesBurnt)
				if err := _Lido.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
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

// ParseSharesBurnt is a log parse operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_Lido *LidoFilterer) ParseSharesBurnt(log types.Log) (*LidoSharesBurnt, error) {
	event := new(LidoSharesBurnt)
	if err := _Lido.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingLimitRemovedIterator is returned from FilterStakingLimitRemoved and is used to iterate over the raw logs and unpacked data for StakingLimitRemoved events raised by the Lido contract.
type LidoStakingLimitRemovedIterator struct {
	Event *LidoStakingLimitRemoved // Event containing the contract specifics and raw log

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
func (it *LidoStakingLimitRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingLimitRemoved)
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
		it.Event = new(LidoStakingLimitRemoved)
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
func (it *LidoStakingLimitRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingLimitRemoved represents a StakingLimitRemoved event raised by the Lido contract.
type LidoStakingLimitRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitRemoved is a free log retrieval operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_Lido *LidoFilterer) FilterStakingLimitRemoved(opts *bind.FilterOpts) (*LidoStakingLimitRemovedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &LidoStakingLimitRemovedIterator{contract: _Lido.contract, event: "StakingLimitRemoved", logs: logs, sub: sub}, nil
}

// WatchStakingLimitRemoved is a free log subscription operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_Lido *LidoFilterer) WatchStakingLimitRemoved(opts *bind.WatchOpts, sink chan<- *LidoStakingLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingLimitRemoved)
				if err := _Lido.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
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

// ParseStakingLimitRemoved is a log parse operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_Lido *LidoFilterer) ParseStakingLimitRemoved(log types.Log) (*LidoStakingLimitRemoved, error) {
	event := new(LidoStakingLimitRemoved)
	if err := _Lido.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingLimitSetIterator is returned from FilterStakingLimitSet and is used to iterate over the raw logs and unpacked data for StakingLimitSet events raised by the Lido contract.
type LidoStakingLimitSetIterator struct {
	Event *LidoStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *LidoStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingLimitSet)
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
		it.Event = new(LidoStakingLimitSet)
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
func (it *LidoStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingLimitSet represents a StakingLimitSet event raised by the Lido contract.
type LidoStakingLimitSet struct {
	MaxStakeLimit              *big.Int
	StakeLimitIncreasePerBlock *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitSet is a free log retrieval operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_Lido *LidoFilterer) FilterStakingLimitSet(opts *bind.FilterOpts) (*LidoStakingLimitSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return &LidoStakingLimitSetIterator{contract: _Lido.contract, event: "StakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingLimitSet is a free log subscription operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_Lido *LidoFilterer) WatchStakingLimitSet(opts *bind.WatchOpts, sink chan<- *LidoStakingLimitSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingLimitSet)
				if err := _Lido.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
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

// ParseStakingLimitSet is a log parse operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_Lido *LidoFilterer) ParseStakingLimitSet(log types.Log) (*LidoStakingLimitSet, error) {
	event := new(LidoStakingLimitSet)
	if err := _Lido.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingPausedIterator is returned from FilterStakingPaused and is used to iterate over the raw logs and unpacked data for StakingPaused events raised by the Lido contract.
type LidoStakingPausedIterator struct {
	Event *LidoStakingPaused // Event containing the contract specifics and raw log

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
func (it *LidoStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingPaused)
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
		it.Event = new(LidoStakingPaused)
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
func (it *LidoStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingPaused represents a StakingPaused event raised by the Lido contract.
type LidoStakingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingPaused is a free log retrieval operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_Lido *LidoFilterer) FilterStakingPaused(opts *bind.FilterOpts) (*LidoStakingPausedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return &LidoStakingPausedIterator{contract: _Lido.contract, event: "StakingPaused", logs: logs, sub: sub}, nil
}

// WatchStakingPaused is a free log subscription operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_Lido *LidoFilterer) WatchStakingPaused(opts *bind.WatchOpts, sink chan<- *LidoStakingPaused) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingPaused)
				if err := _Lido.contract.UnpackLog(event, "StakingPaused", log); err != nil {
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

// ParseStakingPaused is a log parse operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_Lido *LidoFilterer) ParseStakingPaused(log types.Log) (*LidoStakingPaused, error) {
	event := new(LidoStakingPaused)
	if err := _Lido.contract.UnpackLog(event, "StakingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingResumedIterator is returned from FilterStakingResumed and is used to iterate over the raw logs and unpacked data for StakingResumed events raised by the Lido contract.
type LidoStakingResumedIterator struct {
	Event *LidoStakingResumed // Event containing the contract specifics and raw log

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
func (it *LidoStakingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingResumed)
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
		it.Event = new(LidoStakingResumed)
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
func (it *LidoStakingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingResumed represents a StakingResumed event raised by the Lido contract.
type LidoStakingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingResumed is a free log retrieval operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_Lido *LidoFilterer) FilterStakingResumed(opts *bind.FilterOpts) (*LidoStakingResumedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return &LidoStakingResumedIterator{contract: _Lido.contract, event: "StakingResumed", logs: logs, sub: sub}, nil
}

// WatchStakingResumed is a free log subscription operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_Lido *LidoFilterer) WatchStakingResumed(opts *bind.WatchOpts, sink chan<- *LidoStakingResumed) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingResumed)
				if err := _Lido.contract.UnpackLog(event, "StakingResumed", log); err != nil {
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

// ParseStakingResumed is a log parse operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_Lido *LidoFilterer) ParseStakingResumed(log types.Log) (*LidoStakingResumed, error) {
	event := new(LidoStakingResumed)
	if err := _Lido.contract.UnpackLog(event, "StakingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the Lido contract.
type LidoStoppedIterator struct {
	Event *LidoStopped // Event containing the contract specifics and raw log

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
func (it *LidoStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStopped)
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
		it.Event = new(LidoStopped)
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
func (it *LidoStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStopped represents a Stopped event raised by the Lido contract.
type LidoStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_Lido *LidoFilterer) FilterStopped(opts *bind.FilterOpts) (*LidoStoppedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &LidoStoppedIterator{contract: _Lido.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_Lido *LidoFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *LidoStopped) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStopped)
				if err := _Lido.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_Lido *LidoFilterer) ParseStopped(log types.Log) (*LidoStopped, error) {
	event := new(LidoStopped)
	if err := _Lido.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the Lido contract.
type LidoSubmittedIterator struct {
	Event *LidoSubmitted // Event containing the contract specifics and raw log

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
func (it *LidoSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoSubmitted)
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
		it.Event = new(LidoSubmitted)
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
func (it *LidoSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoSubmitted represents a Submitted event raised by the Lido contract.
type LidoSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*LidoSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &LidoSubmittedIterator{contract: _Lido.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *LidoSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoSubmitted)
				if err := _Lido.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) ParseSubmitted(log types.Log) (*LidoSubmitted, error) {
	event := new(LidoSubmitted)
	if err := _Lido.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoTokenRebasedIterator is returned from FilterTokenRebased and is used to iterate over the raw logs and unpacked data for TokenRebased events raised by the Lido contract.
type LidoTokenRebasedIterator struct {
	Event *LidoTokenRebased // Event containing the contract specifics and raw log

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
func (it *LidoTokenRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoTokenRebased)
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
		it.Event = new(LidoTokenRebased)
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
func (it *LidoTokenRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoTokenRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoTokenRebased represents a TokenRebased event raised by the Lido contract.
type LidoTokenRebased struct {
	ReportTimestamp    *big.Int
	TimeElapsed        *big.Int
	PreTotalShares     *big.Int
	PreTotalEther      *big.Int
	PostTotalShares    *big.Int
	PostTotalEther     *big.Int
	SharesMintedAsFees *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenRebased is a free log retrieval operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_Lido *LidoFilterer) FilterTokenRebased(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*LidoTokenRebasedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &LidoTokenRebasedIterator{contract: _Lido.contract, event: "TokenRebased", logs: logs, sub: sub}, nil
}

// WatchTokenRebased is a free log subscription operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_Lido *LidoFilterer) WatchTokenRebased(opts *bind.WatchOpts, sink chan<- *LidoTokenRebased, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoTokenRebased)
				if err := _Lido.contract.UnpackLog(event, "TokenRebased", log); err != nil {
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

// ParseTokenRebased is a log parse operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_Lido *LidoFilterer) ParseTokenRebased(log types.Log) (*LidoTokenRebased, error) {
	event := new(LidoTokenRebased)
	if err := _Lido.contract.UnpackLog(event, "TokenRebased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lido contract.
type LidoTransferIterator struct {
	Event *LidoTransfer // Event containing the contract specifics and raw log

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
func (it *LidoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoTransfer)
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
		it.Event = new(LidoTransfer)
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
func (it *LidoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoTransfer represents a Transfer event raised by the Lido contract.
type LidoTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lido *LidoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoTransferIterator{contract: _Lido.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lido *LidoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LidoTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoTransfer)
				if err := _Lido.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Lido *LidoFilterer) ParseTransfer(log types.Log) (*LidoTransfer, error) {
	event := new(LidoTransfer)
	if err := _Lido.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the Lido contract.
type LidoTransferSharesIterator struct {
	Event *LidoTransferShares // Event containing the contract specifics and raw log

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
func (it *LidoTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoTransferShares)
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
		it.Event = new(LidoTransferShares)
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
func (it *LidoTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoTransferShares represents a TransferShares event raised by the Lido contract.
type LidoTransferShares struct {
	From        common.Address
	To          common.Address
	SharesValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_Lido *LidoFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoTransferSharesIterator{contract: _Lido.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_Lido *LidoFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *LidoTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoTransferShares)
				if err := _Lido.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_Lido *LidoFilterer) ParseTransferShares(log types.Log) (*LidoTransferShares, error) {
	event := new(LidoTransferShares)
	if err := _Lido.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoUnbufferedIterator is returned from FilterUnbuffered and is used to iterate over the raw logs and unpacked data for Unbuffered events raised by the Lido contract.
type LidoUnbufferedIterator struct {
	Event *LidoUnbuffered // Event containing the contract specifics and raw log

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
func (it *LidoUnbufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoUnbuffered)
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
		it.Event = new(LidoUnbuffered)
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
func (it *LidoUnbufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoUnbufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoUnbuffered represents a Unbuffered event raised by the Lido contract.
type LidoUnbuffered struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbuffered is a free log retrieval operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_Lido *LidoFilterer) FilterUnbuffered(opts *bind.FilterOpts) (*LidoUnbufferedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return &LidoUnbufferedIterator{contract: _Lido.contract, event: "Unbuffered", logs: logs, sub: sub}, nil
}

// WatchUnbuffered is a free log subscription operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_Lido *LidoFilterer) WatchUnbuffered(opts *bind.WatchOpts, sink chan<- *LidoUnbuffered) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoUnbuffered)
				if err := _Lido.contract.UnpackLog(event, "Unbuffered", log); err != nil {
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

// ParseUnbuffered is a log parse operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_Lido *LidoFilterer) ParseUnbuffered(log types.Log) (*LidoUnbuffered, error) {
	event := new(LidoUnbuffered)
	if err := _Lido.contract.UnpackLog(event, "Unbuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoWithdrawalsReceivedIterator is returned from FilterWithdrawalsReceived and is used to iterate over the raw logs and unpacked data for WithdrawalsReceived events raised by the Lido contract.
type LidoWithdrawalsReceivedIterator struct {
	Event *LidoWithdrawalsReceived // Event containing the contract specifics and raw log

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
func (it *LidoWithdrawalsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoWithdrawalsReceived)
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
		it.Event = new(LidoWithdrawalsReceived)
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
func (it *LidoWithdrawalsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoWithdrawalsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoWithdrawalsReceived represents a WithdrawalsReceived event raised by the Lido contract.
type LidoWithdrawalsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsReceived is a free log retrieval operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_Lido *LidoFilterer) FilterWithdrawalsReceived(opts *bind.FilterOpts) (*LidoWithdrawalsReceivedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return &LidoWithdrawalsReceivedIterator{contract: _Lido.contract, event: "WithdrawalsReceived", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsReceived is a free log subscription operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_Lido *LidoFilterer) WatchWithdrawalsReceived(opts *bind.WatchOpts, sink chan<- *LidoWithdrawalsReceived) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoWithdrawalsReceived)
				if err := _Lido.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
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

// ParseWithdrawalsReceived is a log parse operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_Lido *LidoFilterer) ParseWithdrawalsReceived(log types.Log) (*LidoWithdrawalsReceived, error) {
	event := new(LidoWithdrawalsReceived)
	if err := _Lido.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
