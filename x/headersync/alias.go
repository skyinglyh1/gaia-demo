// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/internal/types
package headersync

import (
	"github.com/cosmos/gaia/x/headersync/internal/keeper"
	"github.com/cosmos/gaia/x/headersync/internal/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	QueryParameters   = types.QueryParameters
	RouterKey         = types.RouterKey

	AttributeValueCategory        = types.AttributeValueCategory
	EventTypeSyncHeader           = types.EventTypeSyncHeader
	AttributeKeyChainId           = types.AttributeKeyChainId
	AttributeKeyHeight            = types.AttributeKeyHeight
	AttributeKeyBlockHash         = types.AttributeKeyBlockHash
	AttributeKeyNativeChainHeight = types.AttributeKeyNativeChainHeight
)

var (
	// functions aliases
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier

	NewKeeper              = keeper.NewKeeper
	NewMsgSyncGenesisParam = types.NewMsgSyncGenesisParam
	NewMsgSyncHeadersParam = types.NewMsgSyncHeadersParam

	NewQueryHeaderParams        = types.NewQueryHeaderParams
	NewQueryCurrentHeightParams = types.NewQueryCurrentHeightParams

	// key function
	GetBlockHeaderKey    = keeper.GetBlockHeaderKey
	GetBlockHashKey      = keeper.GetBlockHashKey
	GetBlockCurHeightKey = keeper.GetBlockCurHeightKey
	GetConsensusPeerKey  = keeper.GetConsensusPeerKey
	GetKeyHeightsKey     = keeper.GetKeyHeightsKey

	// variable aliases
	ModuleCdc = types.ModuleCdc

	ErrInvalidChainId            = types.ErrInvalidChainId
	ErrEmptyTargetHash           = types.ErrEmptyTargetHash
	ErrBelowCrossedLimit         = types.ErrBelowCrossedLimit
	ErrCrossedAmountOverLimit    = types.ErrCrossedAmountOverLimit
	ErrCrossedAmountOverflow     = types.ErrCrossedAmountOverflow
	ErrSupplyKeeperMintCoinsFail = types.ErrSupplyKeeperMintCoinsFail

	BlockHeaderPrefix   = keeper.BlockHeaderPrefix
	BlockHashPrefix     = keeper.BlockHashPrefix
	ConsensusPeerPrefix = keeper.ConsensusPeerPrefix
	KeyHeightPrefix     = keeper.KeyHeightPrefix

	BlockCurrentHeightKey = keeper.BlockCurrentHeightKey
	QueryHeader           = types.QueryHeader
	QueryCurrentHeight    = types.QueryCurrentHeight
	QueryKeyHeights       = types.QueryKeyHeights
	QueryKeyHeight        = types.QueryKeyHeight
)

type (
	Keeper = keeper.Keeper

	KeeperI = keeper.KeeperI

	MsgSyncGenesisParam = types.MsgSyncGenesisParam
	MsgSyncHeadersParam = types.MsgSyncHeadersParam
)