package flow

import (
	"errors"
	"fmt"

	"github.com/onflow/flow-go-sdk/access/http"
)

type AccountProofSpec struct {
	Ftype      string           `json:"f_type"`
	Fvsn       string           `json:"f_vsn"`
	Address    string           `json:"address"`
	Nonce      string           `json:"nonce"`
	Signatures []SignaturesSpec `json:"signatures"`
}

type SignaturesSpec struct {
	Ftype     string `json:"f_type"`
	Fvsn      string `json:"f_vsn"`
	Addr      string `json:"addr"`
	KeyId     int    `json:"keyId"`
	Signature string `json:"signature"`
}

func getVerifyAccountProofScript(network string) (string, error) {
	var fclCryptoContract string
	switch network {
	case http.EmulatorHost:
		fclCryptoContract = "0xf8d6e0586b0a20c7"
	case http.TestnetHost:
		fclCryptoContract = "0x74daa6f9c7ef24b1"
	case http.MainnetHost:
		fclCryptoContract = "0xb4b82a1c9d21d284"
	default:
		return "", errors.New("Network is not supported")
	}

	script := fmt.Sprintf(`
		import FCLCrypto from %s

		pub fun main(
			address: Address,
			message: String,
			keyIndices: [Int],
			signatures: [String]
		): Bool {
			return FCLCrypto.%s(address: address, message: message, keyIndices: keyIndices, signatures: signatures)
		}
	`, fclCryptoContract, "verifyAccountProofSignatures")

	return script, nil
}
