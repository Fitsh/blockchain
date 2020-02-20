package main

//import "fmt"

// 8.重构代码

func main() {
	bc := NewBlockChain()
	cli := CLI{bc}
	cli.Run()
	//	bc.AddBlock("bc add a block")
	//	bc.AddBlock("bc add 2th block")
	//
	//	//	for i, block := range bc.blocks {
	//	//		fmt.Printf("=============当前区块高度 %d ============\n", i)
	//	//		fmt.Printf("前区块哈希: %x\n", block.PrevHash)
	//	//		fmt.Printf("当前区块哈希: %x\n", block.Hash)
	//	//		fmt.Printf("数据: %s\n", block.Data)
	//	//	}
	//	it := bc.NewBlockChainIterator()
	//	for {
	//		block := it.Next()
	//		fmt.Printf("============= ============\n")
	//		fmt.Printf("前区块哈希: %x\n", block.PrevHash)
	//		fmt.Printf("当前区块哈希: %x\n", block.Hash)
	//		fmt.Printf("数据: %s\n\n", block.Data)
	//
	//		if len(block.PrevHash) == 0 {
	//			break
	//		}
	//	}
}
