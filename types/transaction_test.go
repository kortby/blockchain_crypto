package types

import (
	"testing"

	"github.com/kortby/blockchaincrypto/crypto"
	"github.com/kortby/blockchaincrypto/proto"
	"github.com/kortby/blockchaincrypto/util"
	"github.com/stretchr/testify/assert"
)


func TestNewTransaction(t *testing.T) {
	var (
		fromPrivKey = crypto.GeneratePrivateKey()
		fromAddress = fromPrivKey.Public().Address().Bytes()
		toPrivKey = crypto.GeneratePrivateKey()
		toAddress = toPrivKey.Public().Address().Bytes()
	)
	input := &proto.TxInput{
		PervTxHash: util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey: fromPrivKey.Public().Bytes(),
	}
	outputA := &proto.TxOutput{
		Amount: 5,
		Address: toAddress,
	}
	outputB := &proto.TxOutput{
		Amount: 95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{outputA, outputB},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()

	assert.True(t, VerifyTransaction(tx))
}