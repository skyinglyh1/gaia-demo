package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply/exported"
	"github.com/cosmos/gaia/x/ft/internal/types"
	selfexported "github.com/cosmos/gaia/x/lockproxy/exported"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the mint store
type Keeper struct {
	cdc             *codec.Codec
	storeKey        sdk.StoreKey
	paramSpace      params.Subspace
	authKeeper      types.AccountKeeper
	bankKeeper      types.BankKeeper
	supplyKeeper    types.SupplyKeeper
	ccmKeeper       types.CrossChainManager
	lockProxyKeeper types.LockProxyKeeper
	selfexported.UnlockKeeper
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace, ak types.AccountKeeper, bankKeeper types.BankKeeper, supplyKeeper types.SupplyKeeper, lockProxyKeeper types.LockProxyKeeper, ccmKeeper types.CrossChainManager) Keeper {

	// ensure mint module account is set
	if addr := supplyKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("the %s module account has not been set", types.ModuleName))
	}

	return Keeper{
		cdc:             cdc,
		storeKey:        key,
		authKeeper:      ak,
		bankKeeper:      bankKeeper,
		supplyKeeper:    supplyKeeper,
		lockProxyKeeper: lockProxyKeeper,
		ccmKeeper:       ccmKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetModuleAccount(ctx sdk.Context) exported.ModuleAccountI {
	return k.supplyKeeper.GetModuleAccount(ctx, types.ModuleName)
}

func (k Keeper) EnsureAccountExist(ctx sdk.Context, addr sdk.AccAddress) error {
	acct := k.authKeeper.GetAccount(ctx, addr)
	if acct == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("lockproxy: account %s does not exist", addr.String()))
	}
	return nil
}

func (k Keeper) CreateCoins(ctx sdk.Context, creator sdk.AccAddress, coins sdk.Coins) error {
	for _, coin := range coins {
		if reason, exist := k.ExistDenom(ctx, coin.Denom); exist {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CreateCoins Error: denom:%s already exist, due to reason:%s", coin.Denom, reason))
		}
		k.ccmKeeper.SetDenomCreator(ctx, coin.Denom, creator)
	}
	if err := k.MintCoins(ctx, creator, sdk.NewCoins(coins...)); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("CreateCoins error: %v", err))
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateAndDelegateCoinToProxy,
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, coins.String()),
		),
	})
	return nil
}

func (k Keeper) MintCoins(ctx sdk.Context, toAcct sdk.AccAddress, amt sdk.Coins) error {

	_, err := k.bankKeeper.AddCoins(ctx, toAcct, amt)
	if err != nil {
		//panic(err)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("ft.keeper.bankKeeper.AddCoins(ctx, %s, %s), error is %+v", toAcct.String(), amt.String(), err))
	}

	// update total supply
	supply := k.supplyKeeper.GetSupply(ctx)
	supply = supply.Inflate(amt)

	k.supplyKeeper.SetSupply(ctx, supply)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("minted coin:%s to account:%s ", amt.String(), toAcct.String()))

	return nil
}

// BurnCoins burns coins deletes coins from the balance of the module account.
// Panics if the name maps to a non-burner module account or if the amount is invalid.
func (k Keeper) BurnCoins(ctx sdk.Context, fromAcct sdk.AccAddress, amt sdk.Coins) error {

	_, err := k.bankKeeper.SubtractCoins(ctx, fromAcct, amt)
	if err != nil {
		panic(err)
	}

	// update total supply
	supply := k.supplyKeeper.GetSupply(ctx)
	supply = supply.Deflate(amt)
	k.supplyKeeper.SetSupply(ctx, supply)

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("burned coin:%s from account:%s ", amt.String(), fromAcct.String()))
	return nil
}
