package test

import (
	"encoding/hex"
	"errors"
	"fmt"
	. "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/gaia/app"
	"github.com/tendermint/tendermint/crypto"
	rpchttp "github.com/tendermint/tendermint/rpc/client"
	"io/ioutil"
	"testing"
	"time"
)

func sendMsg(client *rpchttp.HTTP, fromAddr sdk.AccAddress, signerPriv crypto.PrivKey, appCdc *codec.Codec, msg sdk.Msg) error {
	// make sure the account exist in auth module
	bs, err := appCdc.MarshalJSON(auth.NewQueryAccountParams(fromAddr))
	if err != nil {
		return fmt.Errorf("marshaljson , auth.NewQueryAccountParams error:%v", err)
	}
	res, err := client.ABCIQueryWithOptions(fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount), bs, rpchttp.ABCIQueryOptions{Prove: true})
	if err != nil {
		return fmt.Errorf("ABCIQuery , error:%v", err)
	}
	if !res.Response.IsOK() {
		return fmt.Errorf("not resp.IsOK")
	}
	//get exported account
	var expAcct exported.Account
	if err := appCdc.UnmarshalJSON(res.Response.Value, &expAcct); err != nil {
		return fmt.Errorf("Get exported account error:%v", err)
	}
	sequenceNumber := expAcct.GetSequence()

	gasPrice, err := sdk.ParseDecCoins("0.0000000000001stake")
	if err != nil {
		return fmt.Errorf("ParseDecCoins to get gasprice error:%v", err)
	}
	var gas uint64 = 200000
	fee, err := CalcCosmosFees(gasPrice, gas)
	if err != nil {
		return fmt.Errorf("CalcCosmosFees to get gasprice error:%v", err)
	}
	toSign := auth.StdSignMsg{
		Sequence:      sequenceNumber,
		AccountNumber: expAcct.GetAccountNumber(),
		//ChainID:       "testing",
		ChainID: ChainID,
		Msgs:    []sdk.Msg{msg},
		Fee:     auth.NewStdFee(200000, fee),
	}
	sig, err := signerPriv.Sign(toSign.Bytes())
	if err != nil {
		return fmt.Errorf("failed to sign raw tx: (error: %v, raw tx: %x)", err, toSign.Bytes())
	}
	tx := auth.NewStdTx([]sdk.Msg{msg}, toSign.Fee, []auth.StdSignature{{signerPriv.PubKey(), sig}}, toSign.Memo)

	txEncoder := auth.DefaultTxEncoder(appCdc)
	rawTx, err := txEncoder(tx)
	if err != nil {
		return fmt.Errorf("failed to encode signed tx: %v", err)
	}

	broadRes, err := client.BroadcastTxSync(rawTx)
	if err != nil {
		return fmt.Errorf("failed to broadcast tx: (error: %v, raw tx: %x)", err, rawTx)

	}
	fmt.Printf("ResultBroadCastTxSync is %v\n", *broadRes)
	if broadRes.Code == 0 {
		fmt.Printf("hash is %x\n", broadRes.Hash)
		time.Sleep(6 * time.Second)
		CheckTxSuccessful(client, hex.EncodeToString(broadRes.Hash))
	}

	return nil
}

func Test_CheckTxSuccess(t *testing.T) {
	client := rpchttp.NewHTTP(ip, "/websocket")
	txHash := "E41D0D509DD7E2FE928471E60A0731AE55CD90BFAE52412F3D590F25C10D2B8B"
	CheckTxSuccessful(client, txHash)
}

func CheckTxSuccessful(client rpchttp.Client, txHash string) {
	cliCtx := NewCLIContext().WithCodec(app.MakeCodec()).WithClient(client).WithFrom("cosmos1vnnpptmw2vlm5h06ej3t23vsx6jaqgtcwexesm").WithTrustNode(true)
	output, err := utils.QueryTx(cliCtx, txHash)
	if err != nil {
		fmt.Printf("QueryTx err:%v", err)
	}
	if output.Code != 0 {
		fmt.Printf("Tx:%s failed \n result:%s\n", txHash, output.String())

	} else {
		fmt.Printf("Tx:%s Success \n result:%s\n", txHash, output.String())
	}
}

func GetCosmosPrivateKey(path string, pwd []byte) (crypto.PrivKey, types.AccAddress, error) {
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, types.AccAddress{}, err
	}

	privKey, err := mintkey.UnarmorDecryptPrivKey(string(bz), string(pwd))
	if err != nil {
		return nil, types.AccAddress{}, fmt.Errorf("failed to decrypt private key: %v", err)
	}

	return privKey, types.AccAddress(privKey.PubKey().Address().Bytes()), nil
}

func CalcCosmosFees(gasPrice types.DecCoins, gas uint64) (types.Coins, error) {
	if gasPrice.IsZero() {
		return types.Coins{}, errors.New("gas price is zero")
	}
	if gas == 0 {
		return types.Coins{}, errors.New("gas is zero")
	}
	glDec := types.NewDec(int64(gas))
	fees := make(types.Coins, len(gasPrice))
	for i, gp := range gasPrice {
		fee := gp.Amount.Mul(glDec)
		fees[i] = types.NewCoin(gp.Denom, fee.Ceil().RoundInt())
	}
	return fees, nil
}
