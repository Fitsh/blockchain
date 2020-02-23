package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 1. 定义交易结构
type Transaction struct {
	TXID      []byte     // 交易ID
	TxInputs  []TxInput  //  交易输入数组
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
	Value      float64 // 转账金额
	PubKeyHash string  // 锁定脚本，我们用地址模拟
}

const reward = 12.5

// 设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	data := buffer.Bytes()

	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

// 提供一个方法，判断当前的交易是否是挖矿交易
func (tx *Transaction) IsCoinBaseTx() bool {
	// 交易input只有一个
	if len(tx.TxInputs) == 1 {
		input := tx.TxInputs[0]
		// 交易Id为空
		// 交易index为-1
		if !bytes.Equal(input.Txid, []byte{}) || input.Index != -1 {
			return false
		}
	} else {
		return false
	}
	return true
}

// 2. 提供创建交易的方法

func NewCoinBaseTx(address string, data string) *Transaction {
	//挖矿交易特点
	// 1. 只有一个input
	// 2. 无需引用交易id
	// 3. 无需引用index
	// 矿工由于挖矿时无需指定签名，所以这个 sig字段可以由矿工自由填写数据，一般是填写矿池的名字
	input := TxInput{[]byte{}, -1, data}
	output := TxOutput{reward, address}
	tx := &Transaction{[]byte{}, []TxInput{input}, []TxOutput{output}}
	tx.SetHash()

	return tx
}

// 3. 创建挖矿交易
// 4. 根据交易调整程序
