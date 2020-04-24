package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateCoins{}, "gaia/MsgCreateCoins", nil)
	cdc.RegisterConcrete(MsgBindProxyParam{}, "gaia/MsgBindProxyParam", nil)
	cdc.RegisterConcrete(MsgBindAssetParam{}, "gaia/MsgBindAssetParam", nil)
	cdc.RegisterConcrete(MsgLock{}, "gaia/MsgLock", nil)

}

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}