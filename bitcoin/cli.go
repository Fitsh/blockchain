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
	addBlock --data Data        add data to blockchain"
	printChain                  "正向打印区块链"
	printChainR                 "反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余饿"
	send FROM TO AMOUNT MINER DATA 由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
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
	case "addBlock":
		fmt.Printf("addBlock\n")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块参数使用不当")
			fmt.Printf(Usage)
		}
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
	default:
		fmt.Printf(Usage)
	}
}
