// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/types
package lockproxy

import (
	"github.com/cosmos/gaia/x/lockproxy/internal/keeper"
	"github.com/cosmos/gaia/x/lockproxy/internal/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	QueryParameters   = types.QueryParameters

	RouterKey = types.RouterKey
)

var (
	// functions aliases
	RegisterCodec             = types.RegisterCodec
	NewKeeper                 = keeper.NewKeeper
	NewQuerier                = keeper.NewQuerier
	NewMsgBindProxyParam      = types.NewMsgBindProxyParam
	NewMsgBindAssetParam      = types.NewMsgBindAssetParam
	NewMsgLock                = types.NewMsgLock
	NewMsgProcessCrossChainTx = types.NewMsgProcessCrossChainTx
	NewMsgCreateCoins         = types.NewMsgCreateCoins

	NewQueryProxyHashParams    = types.NewQueryProxyHashParams
	NewQueryAssetHashParams    = types.NewQueryAssetHashParams
	NewQueryCrossedAmountParam = types.NewQueryCrossedAmountParam
	NewQueryCrossedLimitParam  = types.NewQueryCrossedLimitParam

	ParamKeyTable    = types.ParamKeyTable
	DefaultCoins     = types.DefaultCoins
	ValidateOperator = types.ValidateOperator

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	OperatorKey              = types.OperatorKey
	KeyMintDenom             = types.KeyMintDenom
	KeyInflationRateChange   = types.KeyInflationRateChange
	KeyInflationMax          = types.KeyInflationMax
	KeyInflationMin          = types.KeyInflationMin
	KeyGoalBonded            = types.KeyGoalBonded
	KeyBlocksPerYear         = types.KeyBlocksPerYear
	CurrentChainCrossChainId = types.CurrentChainCrossChainId

	BindProxyPrefix          = keeper.BindProxyPrefix
	BindAssetPrefix          = keeper.BindAssetPrefix
	CrossedLimitPrefix       = keeper.CrossedLimitPrefix
	CrossedAmountPrefix      = keeper.CrossedAmountPrefix
	CrossChainIdKey          = keeper.CrossChainIdKey
	CrossChainTxDetailPrefix = keeper.CrossChainTxDetailPrefix
	CrossChainDoneTxPrefix   = keeper.CrossChainDoneTxPrefix
)

type (
	Keeper     = keeper.Keeper
	Operator   = types.Operator
	CoinsParam = types.CoinsParam

	MsgBindProxyParam      = types.MsgBindProxyParam
	MsgBindAssetParam      = types.MsgBindAssetParam
	MsgLock                = types.MsgLock
	MsgProcessCrossChainTx = types.MsgProcessCrossChainTx
	MsgCreateCoins         = types.MsgCreateCoins
)