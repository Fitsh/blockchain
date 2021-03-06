package main

// 这是一个用来接收命令行参数并且控制区块链操作的文件

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
	printChain                  "正向打印区块链"
	printChainR                 "反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余饿"
	send FROM TO AMOUNT MINER DATA 由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
	newWallet                   "创建一个新的钱包(私钥公钥对)"
	listAddresses               "列出所有的地址"
`

// 接收参数的动作放到一个函数中
func (cli *CLI) Run() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}

	cmd := args[1]
	switch cmd {
	case "printChain":
		fmt.Printf("printChain\n")
		cli.PrintBlockChain()
	case "printChainR":
		fmt.Printf("printChainR\n")
		cli.PrintBlockChainReverse()
	case "getBalance":
		fmt.Printf("getBalance\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		} else {
			fmt.Printf("获取balance参数使用不当\n")
			fmt.Printf(Usage)
		}
	case "send":
		fmt.Printf("转账开始...\n")
		if len(args) != 7 {
			fmt.Printf("参数个数错误，清检查~")
			fmt.Printf(Usage)
			return
		}
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64)
		miner := args[5]
		data := args[6]
		cli.Send(from, to, amount, miner, data)
	case "newWallet":
		fmt.Printf("创建新的钱包...\n")
		cli.NewWallet()
	case "listAddresses":
		fmt.Printf("列出所有的地址...\n")
		cli.ListAddresses()
	default:
		fmt.Printf(Usage)
	}
}
