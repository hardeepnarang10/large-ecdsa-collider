package generator

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hardeepnarang10/large-ecdsa-collider/enum"
	"github.com/hardeepnarang10/large-ecdsa-collider/wallet"
)

func GenerateWallet() wallet.Wallet {
	// Generate private key
	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		enum.LOGGER.Error("Error occurred in generating private ECDSA key: " + err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]

	// Generate corresponding public key
	publicKeyECDSA := privateKeyECDSA.Public()
	publicKeyECDSAAssert, ok := publicKeyECDSA.(*ecdsa.PublicKey)
	if !ok {
		enum.LOGGER.Error("Error occurred while asserting type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSAAssert)
	publicKeyString := hexutil.Encode(publicKeyBytes)[4:]
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSAAssert)

	// Return
	return wallet.Wallet{
		PrivateKey:    privateKeyString,
		PublicKey:     publicKeyString,
		PublicAddress: publicAddress,
		Balance:       0,
	}
}
