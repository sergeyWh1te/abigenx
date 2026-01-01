// Code generated - DO NOT EDIT.

package lido

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ApprovalEvent                         = "Approval"
	CLValidatorsUpdatedEvent              = "CLValidatorsUpdated"
	ContractVersionSetEvent               = "ContractVersionSet"
	DepositedValidatorsChangedEvent       = "DepositedValidatorsChanged"
	EIP712StETHInitializedEvent           = "EIP712StETHInitialized"
	ELRewardsReceivedEvent                = "ELRewardsReceived"
	ETHDistributedEvent                   = "ETHDistributed"
	ExternalBadDebtInternalizedEvent      = "ExternalBadDebtInternalized"
	ExternalEtherTransferredToBufferEvent = "ExternalEtherTransferredToBuffer"
	ExternalSharesBurntEvent              = "ExternalSharesBurnt"
	ExternalSharesMintedEvent             = "ExternalSharesMinted"
	InternalShareRateUpdatedEvent         = "InternalShareRateUpdated"
	LidoLocatorSetEvent                   = "LidoLocatorSet"
	MaxExternalRatioBPSetEvent            = "MaxExternalRatioBPSet"
	RecoverToVaultEvent                   = "RecoverToVault"
	ResumedEvent                          = "Resumed"
	ScriptResultEvent                     = "ScriptResult"
	SharesBurntEvent                      = "SharesBurnt"
	StakingLimitRemovedEvent              = "StakingLimitRemoved"
	StakingLimitSetEvent                  = "StakingLimitSet"
	StakingPausedEvent                    = "StakingPaused"
	StakingResumedEvent                   = "StakingResumed"
	StoppedEvent                          = "Stopped"
	SubmittedEvent                        = "Submitted"
	TokenRebasedEvent                     = "TokenRebased"
	TransferEvent                         = "Transfer"
	TransferSharesEvent                   = "TransferShares"
	UnbufferedEvent                       = "Unbuffered"
	WithdrawalsReceivedEvent              = "WithdrawalsReceived"
)

var (
	ApprovalEventSignature                         = crypto.Keccak256Hash([]byte("Approval(address,address,uint256)"))
	CLValidatorsUpdatedEventSignature              = crypto.Keccak256Hash([]byte("CLValidatorsUpdated(uint256,uint256,uint256)"))
	ContractVersionSetEventSignature               = crypto.Keccak256Hash([]byte("ContractVersionSet(uint256)"))
	DepositedValidatorsChangedEventSignature       = crypto.Keccak256Hash([]byte("DepositedValidatorsChanged(uint256)"))
	EIP712StETHInitializedEventSignature           = crypto.Keccak256Hash([]byte("EIP712StETHInitialized(address)"))
	ELRewardsReceivedEventSignature                = crypto.Keccak256Hash([]byte("ELRewardsReceived(uint256)"))
	ETHDistributedEventSignature                   = crypto.Keccak256Hash([]byte("ETHDistributed(uint256,uint256,uint256,uint256,uint256,uint256)"))
	ExternalBadDebtInternalizedEventSignature      = crypto.Keccak256Hash([]byte("ExternalBadDebtInternalized(uint256)"))
	ExternalEtherTransferredToBufferEventSignature = crypto.Keccak256Hash([]byte("ExternalEtherTransferredToBuffer(uint256)"))
	ExternalSharesBurntEventSignature              = crypto.Keccak256Hash([]byte("ExternalSharesBurnt(uint256)"))
	ExternalSharesMintedEventSignature             = crypto.Keccak256Hash([]byte("ExternalSharesMinted(address,uint256)"))
	InternalShareRateUpdatedEventSignature         = crypto.Keccak256Hash([]byte("InternalShareRateUpdated(uint256,uint256,uint256,uint256)"))
	LidoLocatorSetEventSignature                   = crypto.Keccak256Hash([]byte("LidoLocatorSet(address)"))
	MaxExternalRatioBPSetEventSignature            = crypto.Keccak256Hash([]byte("MaxExternalRatioBPSet(uint256)"))
	RecoverToVaultEventSignature                   = crypto.Keccak256Hash([]byte("RecoverToVault(address,address,uint256)"))
	ResumedEventSignature                          = crypto.Keccak256Hash([]byte("Resumed()"))
	ScriptResultEventSignature                     = crypto.Keccak256Hash([]byte("ScriptResult(address,bytes,bytes,bytes)"))
	SharesBurntEventSignature                      = crypto.Keccak256Hash([]byte("SharesBurnt(address,uint256,uint256,uint256)"))
	StakingLimitRemovedEventSignature              = crypto.Keccak256Hash([]byte("StakingLimitRemoved()"))
	StakingLimitSetEventSignature                  = crypto.Keccak256Hash([]byte("StakingLimitSet(uint256,uint256)"))
	StakingPausedEventSignature                    = crypto.Keccak256Hash([]byte("StakingPaused()"))
	StakingResumedEventSignature                   = crypto.Keccak256Hash([]byte("StakingResumed()"))
	StoppedEventSignature                          = crypto.Keccak256Hash([]byte("Stopped()"))
	SubmittedEventSignature                        = crypto.Keccak256Hash([]byte("Submitted(address,uint256,address)"))
	TokenRebasedEventSignature                     = crypto.Keccak256Hash([]byte("TokenRebased(uint256,uint256,uint256,uint256,uint256,uint256,uint256)"))
	TransferEventSignature                         = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	TransferSharesEventSignature                   = crypto.Keccak256Hash([]byte("TransferShares(address,address,uint256)"))
	UnbufferedEventSignature                       = crypto.Keccak256Hash([]byte("Unbuffered(uint256)"))
	WithdrawalsReceivedEventSignature              = crypto.Keccak256Hash([]byte("WithdrawalsReceived(uint256)"))
)

var Events = map[string]common.Hash{
	ApprovalEvent:                         ApprovalEventSignature,
	CLValidatorsUpdatedEvent:              CLValidatorsUpdatedEventSignature,
	ContractVersionSetEvent:               ContractVersionSetEventSignature,
	DepositedValidatorsChangedEvent:       DepositedValidatorsChangedEventSignature,
	EIP712StETHInitializedEvent:           EIP712StETHInitializedEventSignature,
	ELRewardsReceivedEvent:                ELRewardsReceivedEventSignature,
	ETHDistributedEvent:                   ETHDistributedEventSignature,
	ExternalBadDebtInternalizedEvent:      ExternalBadDebtInternalizedEventSignature,
	ExternalEtherTransferredToBufferEvent: ExternalEtherTransferredToBufferEventSignature,
	ExternalSharesBurntEvent:              ExternalSharesBurntEventSignature,
	ExternalSharesMintedEvent:             ExternalSharesMintedEventSignature,
	InternalShareRateUpdatedEvent:         InternalShareRateUpdatedEventSignature,
	LidoLocatorSetEvent:                   LidoLocatorSetEventSignature,
	MaxExternalRatioBPSetEvent:            MaxExternalRatioBPSetEventSignature,
	RecoverToVaultEvent:                   RecoverToVaultEventSignature,
	ResumedEvent:                          ResumedEventSignature,
	ScriptResultEvent:                     ScriptResultEventSignature,
	SharesBurntEvent:                      SharesBurntEventSignature,
	StakingLimitRemovedEvent:              StakingLimitRemovedEventSignature,
	StakingLimitSetEvent:                  StakingLimitSetEventSignature,
	StakingPausedEvent:                    StakingPausedEventSignature,
	StakingResumedEvent:                   StakingResumedEventSignature,
	StoppedEvent:                          StoppedEventSignature,
	SubmittedEvent:                        SubmittedEventSignature,
	TokenRebasedEvent:                     TokenRebasedEventSignature,
	TransferEvent:                         TransferEventSignature,
	TransferSharesEvent:                   TransferSharesEventSignature,
	UnbufferedEvent:                       UnbufferedEventSignature,
	WithdrawalsReceivedEvent:              WithdrawalsReceivedEventSignature,
}
