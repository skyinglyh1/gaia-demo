package keeper

import (
	"testing"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tm-db"
	abci "github.com/tendermint/tendermint/abci/types"
	"encoding/hex"
	mctype "github.com/ontio/multi-chain/core/types"
	mcc "github.com/ontio/multi-chain/common"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/cosmos/gaia/x/headersync/internal/types"

)

type testInput struct {
	cdc *codec.Codec
	ctx sdk.Context
	hsKeeper  BaseKeeper
	pk  params.Keeper
}
var header0 = "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c033644e70a2b4f8de4a15c4a0cd79315673b8346d033804807058f3ff4252900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c8365b000000001dac2b7c00000000fd1a057b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a372c2263223a322c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303364313031383338383037656334303739613436666539386436626439613036393061626362643863653136653066626334353230633763376566373838356462227d2c7b22696e646578223a322c226964223a2231323035303361366231623065316336393737663434663336323332613566336236316236613835396234636535313437633439616363666139613432663438336631323034227d2c7b22696e646578223a332c226964223a2231323035303266663764666337303562623561633638643265383932333063363632393939616562313832383431333165396663653934656639666166356239393137353364227d2c7b22696e646578223a342c226964223a2231323035303334343031376363636138323064393066306562623436316466343633333762303932336230616532626365353833636565316132363234633932303865323038227d2c7b22696e646578223a352c226964223a2231323035303331326631303233393531333134336330323938346263346561396438353438383366636466343937333264633732376466613734373438326663383037653634227d2c7b22696e646578223a362c226964223a2231323035303333336334343833376464623934616435666130656234363062306634393135346639303530333631396434643263386565303833333066623831353834316432227d2c7b22696e646578223a372c226964223a2231323035303363366536383165353135346566626136346337356230616131636135343438396261653736353330373764313664646439373236663336356265333036323264227d5d2c22706f735f7461626c65223a5b362c352c342c332c372c322c372c372c352c352c322c322c322c322c362c352c322c342c312c332c342c312c342c332c332c322c342c352c372c312c342c332c342c352c332c352c352c342c322c312c342c332c312c352c352c352c322c362c342c332c312c362c322c322c312c332c332c322c332c372c372c362c342c342c362c372c372c362c322c362c372c372c312c332c342c312c352c362c322c372c342c342c362c352c312c332c352c372c352c332c312c362c312c322c362c362c312c372c362c362c372c332c372c312c315d2c226d61785f626c6f636b5f6368616e67655f76696577223a31303030307d7dfd9e5473b163f591a8829d83288809d97c20ab2a0000"
var header1 = "000000000000000000000000e48232a8468647e98bf2af215912a02d81bae8f94f0eca4e01de1a86ec6331110000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020b31282c6e546fa19c94726399748489950401f2c381f8232550cece1789ea05b98685e010000003705355644d2d063fd0c017b226c6561646572223a332c227672665f76616c7565223a224244714d70463966716a7754596d4e4c6c50795162464e505066546f374261786e677543775563697844713733636856386c6f6e525774474558596d77387671355279372f41505778434d573737433773594c33542b733d222c227672665f70726f6f66223a226d7272714956696352442b534963354c636b6a42426e427577355047466d2b4c5759414a6e4b486e4437537471636f4434376343506c6f6c366473716a42364f446150445a7a774430417332586330335937724c75513d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d00000000000000000000000000000000000000000623120502ff7dfc705bb5ac68d2e89230c662999aeb18284131e9fce94ef9faf5b991753d23120503d101838807ec4079a46fe98d6bd9a0690abcbd8ce16e0fbc4520c7c7ef7885db2312050344017ccca820d90f0ebb461df46337b0923b0ae2bce583cee1a2624c9208e2082312050333c44837ddb94ad5fa0eb460b0f49154f90503619d4d2c8ee08330fb815841d22312050312f10239513143c02984bc4ea9d854883fcdf49732dc727dfa747482fc807e6423120503c6e681e5154efba64c75b0aa1ca54489bae7653077d16ddd9726f365be30622d0642051c0265da7ed153880b8c78e594314fba05624a72ff06457d82f8db51ce35782b1e20b82154e8ad55113f155a47e1ccb1b56c773567e6678682e8f0308942f8329542051c0d0099185ff11de774e55b8ef4601c4b8672a5136debcd8ed50edb25b240866b23fc4bf4602e3d331219296cab8089ee6f6e0513b00ddee0f30468f58fcd56c142051ccdfc928b266182ba0fbed9741099b67517b14594e0c7220b97021c3b7b2efa240ccd293cdb4a7e1472bfd96e1ace841a67ac6b221a78d49f80a26dc81940ec2542051ca73d1477fffff931ea1846515ee62dc486dca1d1029e39fd9db3ce9821d8c7c004eaa8c38221ce4bcc267361297c2d62e38d68ac456e0d38f2f11ed57ceff05b42051b505034287c85586950d395af3a307976005ddb9f15eaef9af0e29d2ed187dc322e5a69123a3e964a91ad66c3519218182df9a3056aa50d68bc2201471ff4580942051b7e1151398a8700ced185d7e5428d15d118cc0bd450671f0a4548bab8e6aac13432e6ee56b1daa9395e6d9b39b71fb4e5252d65119291ba942d914cebc62eb787"
var header100 = "0000000000000000000000005905e989bcc6d6bc9e07e70403244d140d2c7b7cc728813bae0ac41ad70910fc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004e51e28cfe3e5fce4862ff758343eb4a8de8169bb1ebadb5f15a68f9e909b1643ba0685e64000000f9fca3bf095522b4fd0c017b226c6561646572223a352c227672665f76616c7565223a2242502b4a356265477956366a66332f4c525a344256554b55496a78337a54755268426450686a7763496b5239794f796a34456b5445624a7971352b506a686e374f2b383179563348353167753273644f787848457977773d222c227672665f70726f6f66223a22616f66594c465457494c79514a54795868327533334d5a7934764e4e7035567341733752412b445030514d6b53756e693662736265307671737444493052554d6b6f545545612f716c45737232563069584439514c773d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d0000000000000000000000000000000000000000052312050312f10239513143c02984bc4ea9d854883fcdf49732dc727dfa747482fc807e6423120503d101838807ec4079a46fe98d6bd9a0690abcbd8ce16e0fbc4520c7c7ef7885db2312050344017ccca820d90f0ebb461df46337b0923b0ae2bce583cee1a2624c9208e20823120503a6b1b0e1c6977f44f36232a5f3b61b6a859b4ce5147c49accfa9a42f483f12042312050333c44837ddb94ad5fa0eb460b0f49154f90503619d4d2c8ee08330fb815841d20542051cb3282836d80eb2d0e90b689b7ee7c4a728892821ad253a124967fcc42c98640424bc2e0aad81e6f524196b120740538dd4c58c070342b16820a7e2006a9640c942051cdc90a420b3fd2dbcdfc05f3f929847562baa4f64278bbe7f28314a012645eea00c46bad250bb990d3c04b61d27f0d8bf04294d11048dfc0795c5dab23ad3b1c842051ba7ee9c6d0995e02272171dd9cc1d5e3eed20a4a41977a7c416968323c3ffbf935fa19adc107247bcbee2ca7eb37d6f21eb8d7964ed8b6e714f34a26e3fa76f1142051cb6e1d5cf25897220ea40fc81b1f390a2b9bd0c5085a9d08889eb0abbe94ad40c0278802d73ea8458be483fd6962d7f7b360125d0b717c3ba040eafe70cac9e1842051ca96c8a95e1731945eefce36640b11235dd9960f2bdbe20177ceefb05a9447fbf4791d12c7aaf296950f81cbb934d3f221dce2c38515813a0e894072f1ceb9ab5"

func setupTestInput() testInput {
	newDb := dbm.NewMemDB()

	cdc := codec.New()
	auth.RegisterCodec(cdc)

	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)

	//cdc.RegisterInterface((*interface{})(nil), nil)

	hsKey := sdk.NewKVStoreKey("headersync")
	paKey := sdk.NewKVStoreKey(params.ModuleName)
	tKey := sdk.NewTransientStoreKey(params.TStoreKey)


	ms := store.NewCommitMultiStore(newDb)
	ms.MountStoreWithDB(hsKey, sdk.StoreTypeIAVL, newDb)
	ms.MountStoreWithDB(paKey, sdk.StoreTypeIAVL, newDb)
	ms.MountStoreWithDB(tKey, sdk.StoreTypeIAVL, newDb)

	ms.LoadLatestVersion()

	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[sdk.AccAddress([]byte("moduleAcc")).String()] = true


	pk := params.NewKeeper(cdc, paKey, tKey, params.DefaultCodespace)
	hsKeeper := NewBaseKeeper(cdc, hsKey, pk.Subspace("hssubspace"))



	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	return testInput{cdc: cdc, ctx: ctx, hsKeeper: hsKeeper, pk: pk}
}

func Test_BaseKeeper_SetGetHeaderByHeight(t *testing.T) {
	input := setupTestInput()
	ctx := input.ctx

	mcSerializedBs, _ := hex.DecodeString(header0)
	genesisHeader := new(mctype.Header)
	genesisHeader.Deserialization(mcc.NewZeroCopySource(mcSerializedBs))


	err := input.hsKeeper.SetBlockHeader(ctx, genesisHeader)
	if err != nil {
		t.Errorf("syncgenesisHeader error:%v", err)
	}
	getHeader, err := input.hsKeeper.GetHeaderByHeight(ctx, genesisHeader.ChainID, genesisHeader.Height)
	if err != nil {
		t.Errorf("getHeader error:%v", err)
	}
	fmt.Printf("getheader is %v\n", hex.EncodeToString(input.cdc.MustMarshalJSON(getHeader)))
	assert.Equal(t, getHeader.Version, genesisHeader.Version)
	assert.Equal(t, getHeader.ChainID, genesisHeader.ChainID)
	assert.Equal(t, getHeader.PrevBlockHash, genesisHeader.PrevBlockHash)
	assert.Equal(t, getHeader.TransactionsRoot, genesisHeader.TransactionsRoot)
	assert.Equal(t, getHeader.CrossStateRoot, genesisHeader.CrossStateRoot)
	assert.Equal(t, getHeader.BlockRoot, genesisHeader.BlockRoot)
	assert.Equal(t, getHeader.Timestamp, genesisHeader.Timestamp)
	assert.Equal(t, getHeader.Height, genesisHeader.Height)
	assert.Equal(t, getHeader.ConsensusData, genesisHeader.ConsensusData)
	assert.Equal(t, getHeader.ConsensusPayload, genesisHeader.ConsensusPayload)
	assert.Equal(t, getHeader.NextBookkeeper, genesisHeader.NextBookkeeper)
	assert.Equal(t, getHeader.Bookkeepers, genesisHeader.Bookkeepers)
	assert.Equal(t, getHeader.SigData, genesisHeader.SigData)
	assert.Equal(t, getHeader.Hash(), genesisHeader.Hash())

	err = input.hsKeeper.UpdateConsensusPeer(ctx, genesisHeader)
	if err != nil {
		t.Errorf("UpdateConsensusPeer error:%v", err)
	}


}


func Test_BaseKeeper_SyncGenesisAndHeader(t *testing.T) {
	input := setupTestInput()
	ctx := input.ctx

	mcSerializedBs, _ := hex.DecodeString(header0)
	genesisHeader := new(mctype.Header)
	genesisHeader.Deserialization(mcc.NewZeroCopySource(mcSerializedBs))



	err := input.hsKeeper.SetBlockHeader(ctx, genesisHeader)
	if err != nil {
		t.Errorf("syncgenesisHeader error:%v", err)
	}
	getHeader, err := input.hsKeeper.GetHeaderByHeight(ctx, genesisHeader.ChainID, genesisHeader.Height)
	if err != nil {
		t.Errorf("getHeader error:%v", err)
	}
	fmt.Printf("getheader is %v\n", hex.EncodeToString(input.cdc.MustMarshalJSON(getHeader)))
	assert.Equal(t, getHeader.Version, genesisHeader.Version)
	assert.Equal(t, getHeader.ChainID, genesisHeader.ChainID)
	assert.Equal(t, getHeader.PrevBlockHash, genesisHeader.PrevBlockHash)
	assert.Equal(t, getHeader.TransactionsRoot, genesisHeader.TransactionsRoot)
	assert.Equal(t, getHeader.CrossStateRoot, genesisHeader.CrossStateRoot)
	assert.Equal(t, getHeader.BlockRoot, genesisHeader.BlockRoot)
	assert.Equal(t, getHeader.Timestamp, genesisHeader.Timestamp)
	assert.Equal(t, getHeader.Height, genesisHeader.Height)
	assert.Equal(t, getHeader.ConsensusData, genesisHeader.ConsensusData)
	assert.Equal(t, getHeader.ConsensusPayload, genesisHeader.ConsensusPayload)
	assert.Equal(t, getHeader.NextBookkeeper, genesisHeader.NextBookkeeper)
	assert.Equal(t, getHeader.Bookkeepers, genesisHeader.Bookkeepers)
	assert.Equal(t, getHeader.SigData, genesisHeader.SigData)
	assert.Equal(t, getHeader.Hash(), genesisHeader.Hash())

	err = input.hsKeeper.SyncGenesisHeader(ctx, mcSerializedBs)
	if err != nil {
		t.Errorf("SyncGenesisHeader error:%v", err)
	}


	mcSerializedHeaders := []string{
		header1, // header 1
		header100, // header 100
	}
	mcSerializedBss := make([][]byte, len(mcSerializedHeaders))
	headers := make([]*mctype.Header, len(mcSerializedHeaders))
	for i, mcSerializedHeader := range mcSerializedHeaders {
		mcSerializedBss[i], _ = hex.DecodeString(mcSerializedHeader)
		headerX := new(mctype.Header)
		s1 := mcc.NewZeroCopySource(mcSerializedBss[i])
		headerX.Deserialization(s1)
		headers[i] = headerX
		fmt.Printf("headerX is %v\n", headerX)
	}

	err = input.hsKeeper.SyncBlockHeaders(ctx, mcSerializedBss)
	if err != nil {
		t.Errorf("SyncBlockHeaders error:%v", err)
	}

}