package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const reward = 50

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
	//	Sig   string // 解锁脚本，我们用地址来模拟
	Signature []byte // 真正的数字签名，由r,s拼成的[]byte
	// 约定，这里的pubkey不存储原始的公钥，而是存储x和y拼接的字符串，在校验时拆分
	// 是公钥，不是地址或哈希
	PubKey []byte
}

// 定义交易输出
type TxOutput struct {
	Value float64 // 转账金额
	//PubKeyHash string  // 锁定脚本，我们用地址模拟
	PubKeyHash []byte // 收款方的公钥哈希，是哈希，不是地址和公钥
}

func (output *TxOutput) Lock(address string) {
	// 解码
	// 取出公钥哈希：去除version(1字节),去除校验码(4字节)

	pubKeyHash := GetPubKeyHashFromAddress(address)
	// 真正的锁定动作
	output.PubKeyHash = pubKeyHash
}

// 给TxOutput提供一个创建的方法，否则无法调用Lock
func NewTxOutput(value float64, address string) *TxOutput {
	output := &TxOutput{
		Value: value,
	}

	output.Lock(address)
	return output
}

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
	// 交易Id为空
	// 交易index为-1
	if len(tx.TxInputs) == 1 {
		input := tx.TxInputs[0]
		if bytes.Equal(input.Txid, []byte{}) && input.Index == -1 {
			return true
		}
	}
	return false
}

// 2. 提供创建交易的方法

func NewCoinBaseTx(address string, data string) *Transaction {
	//挖矿交易特点
	// 1. 只有一个input
	// 2. 无需引用交易id
	// 3. 无需引用index
	// 矿工由于挖矿时无需指定签名，所以这个 pubkey字段可以由矿工自由填写数据，一般是填写矿池的名字
	// 签名先填写为空， 后面创建完整交易后， 最后作一次签名即可
	input := TxInput{[]byte{}, -1, nil, []byte(data)}
	//	output := TxOutput{reward, address}
	output := NewTxOutput(reward, address)

	tx := &Transaction{[]byte{}, []TxInput{input}, []TxOutput{*output}}
	tx.SetHash()

	return tx
}

// 创建普通的转账交易

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 创建交易后要进行数字签名 -> 所以需要私钥 -> 打开钱包(NewWallets())
	// 找到自己的钱包， 根据地址饭或自己的wallet
	// 得到对应的公钥,私钥

	wallets := NewWallets()
	wallet := wallets.WalletMap[from]
	if wallet == nil {
		log.Printf("没有找到该地址的钱包，交易创建失败\n")
		return nil
	}

	pubKey := wallet.PubKey
	// privateKey := wallet.PrivateKey

	// 1. 找到合适的UTXO集合 map[string][]uint64
	hash := HashPubKey(pubKey)
	utxos, resValue := bc.FindNeedUTXOs(hash[:], amount)

	if resValue < amount {
		log.Printf("余额不足~\n")
		return nil
	}

	var inputs []TxInput
	var outputs []TxOutput

	// 2. 将这些UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, index := range indexArray {
			input := TxInput{[]byte(id), index, nil, pubKey}
			inputs = append(inputs, input)
		}
	}

	// 3. 创建outputs
	//output := TxOutput{amount, to}
	output := NewTxOutput(amount, to)
	outputs = append(outputs, *output)

	// 4. 如果有零钱，要找零
	if resValue > amount {
		// 找零
		outputs = append(outputs, *NewTxOutput(resValue-amount, from))
	}

	tx := &Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()

	return tx
}

// 3. 创建挖矿交易
// 4. 根据交易调整程序
