package tron

import (
	"fmt"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/fbsobreira/gotron-sdk/pkg/keys/hd"
	"github.com/tyler-smith/go-bip39"
)

// FromMnemonicSeedAndPassphrase derive form mnemonic and passphrase at index
func FromMnemonicSeedAndPassphrase(mnemonic string, index int) (*secp256k1.PrivateKey, *secp256k1.PublicKey) {
	seed := bip39.NewSeed(mnemonic, "pool")
	master, ch := hd.ComputeMastersFromSeed(seed, []byte("Bitcoin seed"))
	private, _ := hd.DerivePrivateKeyForPath(
		secp256k1.S256(),
		master,
		ch,
		fmt.Sprintf("44'/195'/0'/0/%d", index),
	)

	return secp256k1.PrivKeyFromBytes(secp256k1.S256(), private[:])
}