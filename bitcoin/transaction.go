package main

// 1. 定义交易结构
type Transaction struct {
	TXID      []byte     // 交易ID
	TxInPuts  []TxInput  //  交易输入数组
	TxOutputs []TxOutput // 交易输出数组
}

// 定义交易输入
type TxInput struct {
	Txid  []byte // 引用的交易ID
	Index int64  // 引用的output的索引值
	Sig   string // 解锁脚本，我们用地址来模拟
}

// 定义交易输出
type TxOutput struct {
	value      float64 // 转账金额
	PubKeyHash string  // 锁定脚本，我们用地址模拟
}

// 2. 提供创建交易的方法
// 3. 创建挖矿交易
// 4. 根据交易调整程序
