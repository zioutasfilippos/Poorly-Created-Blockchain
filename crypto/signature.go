package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"strconv"
	"strings"
)

func SignTransaction(tx Transaction, privateKey *ecdsa.PrivateKey) string {

	TransactionData := tx.SenderID + tx.ReceiverID + strconv.Itoa(tx.Amount)

	hash := sha256.Sum256([]byte(TransactionData))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return ""
	}

	Signature := r.String() + ":" + s.String()
	return Signature
}

func VerifyTransaction(tx Transaction, Signature string, publicKey *ecdsa.PublicKey) bool {

	TransactionData := tx.SenderID + tx.ReceiverID + strconv.Itoa(tx.Amount)

	hash := sha256.Sum256([]byte(TransactionData))

	signature := strings.Split(tx.Signature, ":")

	r, ok := new(big.Int).SetString(signature[0], 10)
	if !ok {
		return false
	}

	s, ok := new(big.Int).SetString(signature[1], 10)

	return ecdsa.Verify(publicKey, hash[:], r, s)
}
