package main

// 4.引入区块链
type BlockChain struct {
	// 定一个区块链数组
	blocks []*Block
}

// 5. 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创始块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// 定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("创世块", []byte{})
}

// 5.添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	// a 创建区块
	// b 添加到区块链数组中
	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}
