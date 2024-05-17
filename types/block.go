package types

import (
	"crypto/sha256"

	"github.com/kortby/blockchaincrypto/crypto"
	"github.com/kortby/blockchaincrypto/proto"
	pb "google.golang.org/protobuf/proto"
)


func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(b))
}

// hashBlock creates a SHA256 of header
func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}