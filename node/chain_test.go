package node

import (
	"testing"

	"github.com/kortby/blockchaincrypto/types"
	"github.com/kortby/blockchaincrypto/util"
	"github.com/stretchr/testify/assert"
)

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	blcok := util.RandomBlock()
	blockHash := types.HashBlock(blcok)

	assert.Nil(t, chain.AddBlock(blcok))
	
	fetchedBlock, err := chain.GetBlockByHash(blockHash)
	assert.Nil(t, err)
	assert.Equal(t, blcok, fetchedBlock)

	// fetchedBlockByHeight
	assert.Nil(t, err)
	// assert.Equal(t, blcok, GetBlockByHeight())
}

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	for i := 0; i < 100; i++ {
		b := util.RandomBlock()
		assert.Nil(t, chain.AddBlock(b))
		assert.Equal(t, chain.Height(), i)
	}
}