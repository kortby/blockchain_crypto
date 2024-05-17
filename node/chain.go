package node

import (
	"encoding/hex"
	"fmt"

	"github.com/kortby/blockchaincrypto/proto"
	"github.com/kortby/blockchaincrypto/types"
)

type HeaderList struct {
	headers []*proto.Header
}

type Chain struct {
	bloackStore BlockStorer
	headers *HeaderList
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		headers: []*proto.Header{},
	}
}

func (list *HeaderList) Add(h *proto.Header) {
	list.headers = append(list.headers, h)
}

func (list *HeaderList) Height() int {
	return list.Len() - 1
}

func (list *HeaderList) Get(index int) *proto.Header {
	if index > list.Height() {
		panic("index is too high")
	}
	return list.headers[index]
}

func (list *HeaderList) Len() int {
	return len(list.headers)
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{
		bloackStore: bs,
		headers: NewHeaderList(),
	}
}

func (c *Chain) Height() int {
	return c.headers.Height()
}

func (c *Chain) AddBlock(b *proto.Block) error {
	c.headers.Add(b.Header)
	// valid
	return c.bloackStore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.bloackStore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if c.Height() < height {
		return nil, fmt.Errorf("given height (%d) too high (%d)", height, c.Height())
	}
	header := c.headers.Get(height)
	hash := types.HashHeader(header)

	return c.GetBlockByHash(hash)
}
