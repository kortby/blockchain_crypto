package types

import (
	"crypto/sha256"

	"github.com/kortby/blockchaincrypto/crypto"
	"github.com/kortby/blockchaincrypto/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		sig := crypto.SignatureFromBytes(input.Signature)
		pubKey := crypto.PublicKeyFromBytes(input.PublicKey)
		input.Signature = nil
		if !sig.Verify(pubKey, HashTransaction(tx)) {
			return false
		}
	}
	return true
}