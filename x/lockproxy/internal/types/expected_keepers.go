package types // noalias

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/supply/exported"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
	mctype "github.com/ontio/multi-chain/core/types"
)

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	StakingTokenSupply(ctx sdk.Context) sdk.Int
	BondedRatio(ctx sdk.Context) sdk.Dec
}
type AccountKeeper interface {
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authexported.Account

	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
	GetAllAccounts(ctx sdk.Context) []authexported.Account
	SetAccount(ctx sdk.Context, acc authexported.Account)

	IterateAccounts(ctx sdk.Context, process func(authexported.Account) bool)
}

// SupplyKeeper defines the expected supply keeper
type SupplyKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) supplyexported.ModuleAccountI
	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, exported.ModuleAccountI)

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) sdk.Error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) sdk.Error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) sdk.Error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) sdk.Error
	SetSupply(ctx sdk.Context, supply exported.SupplyI)
	GetSupply(ctx sdk.Context) (supply exported.SupplyI)
}

type SyncKeeper interface {
	ProcessHeader(ctx sdk.Context, header *mctype.Header) sdk.Error
	GetHeaderByHeight(ctx sdk.Context, chainId uint64, height uint32) (*mctype.Header, sdk.Error)
}

type SupplyI interface {
	SetTotal(total sdk.Coins) SupplyI
}