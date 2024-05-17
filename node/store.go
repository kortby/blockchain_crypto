package node

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/kortby/blockchaincrypto/proto"
	"github.com/kortby/blockchaincrypto/types"
)

type BlockStorer interface {
	Put(*proto.Block) error
	Get(string) (*proto.Block, error)
}

type MemoryBlockStore struct {
	lock sync.RWMutex
	blocks map[string]*proto.Block
}

func NewMemoryBlockStore() *MemoryBlockStore {
	return &MemoryBlockStore{
		blocks: make(map[string]*proto.Block),
	}
}

func (s *MemoryBlockStore) Put(b *proto.Block) error {
	hash := hex.EncodeToString(types.HashBlock(b))
	defer s.lock.Unlock()
	s.blocks[hash] = b
	return nil
}

func (s *MemoryBlockStore) Get(hash string) (*proto.Block, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	block, ok := s.blocks[hash]
	if !ok {
		return nil, fmt.Errorf("block with hash [%s] does not exist", hash)
	}
	return block, nil
}
