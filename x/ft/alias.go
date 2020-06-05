// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/types
package ft

import (
	"github.com/cosmos/gaia/x/ft/internal/keeper"
	"github.com/cosmos/gaia/x/ft/internal/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	QueryParameters   = types.QueryParameters
	RouterKey         = types.RouterKey

	AttributeValueCategory = types.AttributeValueCategory

	AttributeKeyToChainId        = types.AttributeKeyToChainId
	AttributeKeyToChainProxyHash = types.AttributeKeyToChainProxyHash

	EventTypeBindAsset           = types.EventTypeBindAsset
	AttributeKeySourceAssetDenom = types.AttributeKeySourceAssetDenom
	AttributeKeyFromAssetHash    = types.AttributeKeyFromAssetHash
	AttributeKeyToChainAssetHash = types.AttributeKeyToChainAssetHash

	EventTypeLock           = types.EventTypeLock
	AttributeKeyFromAddress = types.AttributeKeyFromAddress
	AttributeKeyToAddress   = types.AttributeKeyToAddress
	AttributeKeyAmount      = types.AttributeKeyAmount

	EventTypeCreateCrossChainTx = types.EventTypeCreateCrossChainTx
	AttributeCrossChainId       = types.AttributeCrossChainId
	AttributeKeyTxParamHash     = types.AttributeKeyTxParamHash
	AttributeKeyMakeTxParam     = types.AttributeKeyMakeTxParam

	EventTypeVerifyToCosmosProof                        = types.EventTypeVerifyToCosmosProof
	AttributeKeyMerkleValueTxHash                       = types.AttributeKeyMerkleValueTxHash
	AttributeKeyMerkleValueMakeTxParamTxHash            = types.AttributeKeyMerkleValueMakeTxParamTxHash
	AttributeKeyMerkleValueMakeTxParamToContractAddress = types.AttributeKeyMerkleValueMakeTxParamToContractAddress
	AttributeKeyFromChainId                             = types.AttributeKeyFromChainId

	EventTypeUnlock              = types.EventTypeUnlock
	AttributeKeyFromContractHash = types.AttributeKeyFromContractHash
	AttributeKeyToAssetDenom     = types.AttributeKeyToAssetDenom

	EventTypeSetRedeemScript = types.EventTypeSetRedeemScript
	AttributeKeyRedeemKey    = types.AttributeKeyRedeemKey
	AttributeKeyRedeemScript = types.AttributeKeyRedeemScript
)

var (
	// functions aliases
	RegisterCodec = types.RegisterCodec
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier

	NewMsgLock                         = types.NewMsgLock
	NewMsgCreateDenom                  = types.NewMsgCreateDenom
	NewMsgCreateCoinAndDelegateToProxy = types.NewMsgCreateCoinAndDelegateToProxy
	NewMsgBindAssetHash                = types.NewMsgBindAssetHash
	NewMsgCreateCoins                  = types.NewMsgCreateCoins

	// key function
	GetBlockHashKey      = keeper.GetBlockHashKey
	GetBlockCurHeightKey = keeper.GetBlockCurHeightKey
	GetConsensusPeerKey  = keeper.GetConsensusPeerKey
	GetKeyHeightsKey     = keeper.GetKeyHeightsKey
	GetBindProxyKey      = keeper.GetBindProxyKey
	GetBindAssetKey      = keeper.GetBindAssetKey
	GetCrossedLimitKey   = keeper.GetCrossedLimitKey
	GetCrossedAmountKey  = keeper.GetCrossedAmountKey

	ParamKeyTable    = types.ParamKeyTable
	DefaultCoins     = types.DefaultCoins
	ValidateOperator = types.ValidateOperator
	HashToDenom      = types.HashToDenom
	DenomToHash      = types.DenomToHash
	// variable aliases
	ModuleCdc   = types.ModuleCdc
	OperatorKey = types.OperatorKey

	ErrInvalidChainId            = types.ErrInvalidChainId
	ErrEmptyTargetHash           = types.ErrEmptyTargetHash
	ErrBelowCrossedLimit         = types.ErrBelowCrossedLimit
	ErrCrossedAmountOverLimit    = types.ErrCrossedAmountOverLimit
	ErrCrossedAmountOverflow     = types.ErrCrossedAmountOverflow
	ErrSupplyKeeperMintCoinsFail = types.ErrSupplyKeeperMintCoinsFail

	BindProxyPrefix          = keeper.BindProxyPrefix
	BindAssetPrefix          = keeper.BindAssetPrefix
	CrossedLimitPrefix       = keeper.CrossedLimitPrefix
	CrossedAmountPrefix      = keeper.CrossedAmountPrefix
	CrossChainIdKey          = keeper.CrossChainIdKey
	CrossChainTxDetailPrefix = keeper.CrossChainTxDetailPrefix
	CrossChainDoneTxPrefix   = keeper.CrossChainDoneTxPrefix
	RedeemKeyScriptPrefix    = keeper.RedeemKeyScriptPrefix

	RedeemToHashPrefix     = keeper.RedeemToHashPrefix
	ContractToRedeemPrefix = keeper.ContractToRedeemPrefix
	BlockHashPrefix        = keeper.BlockHashPrefix
	ConsensusPeerPrefix    = keeper.ConsensusPeerPrefix
	KeyHeightPrefix        = keeper.KeyHeightPrefix

	BlockCurrentHeightKey = keeper.BlockCurrentHeightKey

	// query balance path

)

type (
	Keeper     = keeper.Keeper
	Operator   = types.Operator
	CoinsParam = types.CoinsParam

	MsgBindAssetHash                = types.MsgBindAssetHash
	MsgLock                         = types.MsgLock
	MsgCreateDenom                  = types.MsgCreateDenom
	MsgCreateCoinAndDelegateToProxy = types.MsgCreateCoinAndDelegateToProxy

	TxArgs = types.TxArgs
)
