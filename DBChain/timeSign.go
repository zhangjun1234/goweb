package DBChain

import (
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"strconv"
	"time"
)

//////////////////
//              //
// helper funcs //
//              //
//////////////////

func MakeAccessCode() string {
	privKey := secp256k1.GenPrivKey()
	now := time.Now().UnixNano() / 1000000
	timeStamp := strconv.Itoa(int(now))

	signature, err := privKey.Sign([]byte(timeStamp))
	if err != nil {
		panic("failed to sign timestamp")
	}

	pubKey := privKey.PubKey()
	pubKeyArray := pubKey.(secp256k1.PubKeySecp256k1)

	encodedPubKey := base58.Encode(pubKeyArray[:])
	encodedSig := base58.Encode(signature)
	return fmt.Sprintf("%s:%s:%s", encodedPubKey, timeStamp, encodedSig)
}
