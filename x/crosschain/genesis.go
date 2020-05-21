package crosschain

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gaia/x/crosschain/internal/types"
)

// GenesisState - minter state
type GenesisState struct {
	Operator Operator `json:"operator" yaml:"operator"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(operator Operator) GenesisState {
	return GenesisState{
		Operator: operator,
	}
}

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	if data.Operator.Operator != nil {
		ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName)).Info("operator:", data.Operator.Operator.String())
		keeper.SetOperator(ctx, data.Operator)
	}

	// check if the module account exists
	moduleAcc := keeper.GetModuleAccount(ctx)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	operator := keeper.GetOperator(ctx)
	return NewGenesisState(operator)
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	err := ValidateOperator(data.Operator)
	if err != nil {
		return err
	}
	return nil
}
