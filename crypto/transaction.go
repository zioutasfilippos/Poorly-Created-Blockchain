package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

type Transaction struct {
	SenderID         string          `json:"senderID"`
	ReceiverID       string          `json:"receiverID"`
	Amount           int             `json:"amount"`
	Signature        string          `json:"signature"`
	SendersPublicKey ecdsa.PublicKey `json:"senderPublicKey"`
}

func CreateWallet() Wallet {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	return Wallet{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PublicKey,
	}
}

func CreateTransaction(wallet Wallet, senderID string, receiverID string, amount int, signature string, senderPublicKey ecdsa.PublicKey) Transaction {

	return Transaction{
		SenderID:         senderID,
		ReceiverID:       receiverID,
		Amount:           amount,
		Signature:        signature,
		SendersPublicKey: wallet.PrivateKey.PublicKey,
	}

}
