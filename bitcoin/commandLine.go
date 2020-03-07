package main

import (
	"fmt"
	"time"
)

// 正向打印
func (cli *CLI) PrintBlockChain() {
	cli.bc.PrintChain()
}

// 反向打印
func (cli *CLI) PrintBlockChainReverse() {
	it := cli.bc.NewBlockChainIterator()
	for {
		block := it.Next()
		fmt.Printf("============= ============\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间 : %s\n", timeFormat)
		fmt.Printf("难度值: %d\n", block.Difficulty)
		fmt.Printf("随机数: %d\n", block.Nonce)
		fmt.Printf("当前区块哈希: %x\n", block.Hash)
		//		fmt.Printf("数据: %s\n\n", block.Data)
		fmt.Printf("数据: %s\n\n", block.Transactions[0].TxInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			break
		}
	}

}

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)

	var total = 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}

	fmt.Printf("%s 的余额为 %f\n", address, total)
}

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {

	// 具体逻辑
	// 挖矿交易
	coinbase := NewCoinBaseTx(miner, data)
	// 普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		fmt.Printf("无效的交易\n")
		return
	}
	// 添加到区块
	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Printf("转账成功~")
}

func (cli *CLI) NewWallet() {
	//	wallet := NewWallet()
	//	address := wallet.NewAddress()
	wallets := NewWallets()
	address := wallets.CreateWallet()
	fmt.Printf("Address: %s\n", address)
	//	fmt.Printf("私钥: %v\n", wallet.PrivateKey)
	//	fmt.Printf("公钥: %x\n", wallet.PubKey)
	//	fmt.Printf("Address: %s\n", address)
}
