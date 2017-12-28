package blockchain

type BlockChain struct {
	Blocks []*Block
}

// AddBlock appends the block to end of blockchain
func (b *BlockChain) AddBlock(data string) {

	prevHash := b.Blocks[len(b.Blocks)-1].Hash
	block := NewBlock(data, prevHash)
	b.Blocks = append(b.Blocks, block)
}

func GenerateGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenerateGenesisBlock()}}
}
