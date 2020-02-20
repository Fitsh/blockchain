package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
}

func (cli *CLI) PrintBlockChain() {
	it := cli.bc.NewBlockChainIterator()
	for {
		block := it.Next()
		fmt.Printf("============= ============\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("难度值: %d\n", block.Difficulty)
		fmt.Printf("随机数: %d\n", block.Nonce)
		fmt.Printf("当前区块哈希: %x\n", block.Hash)
		fmt.Printf("数据: %s\n\n", block.Data)

		if len(block.PrevHash) == 0 {
			break
		}
	}

}
