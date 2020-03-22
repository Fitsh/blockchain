package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"math/big"
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
	privateKey := wallet.PrivateKey

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

	bc.SignTransaction(tx, privateKey)

	return tx
}

// 3. 创建挖矿交易
// 4. 根据交易调整程序

// 签名的具体实现
// 参数：私钥，inputs里面所有引用的交易的结构map[string]Transaction
func (tx *Transaction) Sign(privateKey *ecdsa.PrivateKey, prevTxs map[string]Transaction) {
	//TODO
	// 1. 创建一个当前交易的副本：txCopy 使用函数：trimmedCopy , 要把Signature和pubkey 设置为nil
	// 2. 循环遍历txCopy 的inputs,得到这个input索引的output的公钥哈希
	// 3. 生成要签名的数据，要签名的数据一定是哈希值
	//		a. 我们要对每一个input都要签名一次，签名的数据是当前input引用的output的哈希+当前的outputs（都承载在当前这个txCopy里面）
	//		b.要对这个拼好的txCopy进行哈希处理，SetHash 得到TXID， 这个TXID就是我们要签名的最终数据
	// 4. 执行签名动作得到r,s 字节流
	// 5. 放到我们所签名的input的Signature中

	if tx.IsCoinBaseTx() {
		return
	}

	txCopy := tx.TrimmedCopy()

	for i, input := range txCopy.TxInputs {
		prevTx := prevTxs[string(input.Txid)]
		if len(prevTx.TXID) == 0 {
			log.Panic("引用的交易无效")
		}

		// 不要对input赋值，input只是一个副本
		txCopy.TxInputs[i].PubKey = prevTx.TxOutputs[input.Index].PubKeyHash

		// 所需要三个数据都具备了，开始哈希
		// 生成要签名的数据，要签名的数据一定是哈希
		txCopy.SetHash()

		// 还原数据，以免影响input的签名
		txCopy.TxInputs[i].PubKey = nil

		signData := txCopy.TXID

		r, s, err := ecdsa.Sign(rand.Reader, privateKey, signData)
		if err != nil {
			log.Panic(err)
		}

		signature := append(r.Bytes(), s.Bytes()...)
		tx.TxInputs[i].Signature = signature
	}
}

func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	for _, input := range tx.TxInputs {
		inputs = append(inputs, TxInput{input.Txid, input.Index, nil, nil})
	}

	for _, output := range tx.TxOutputs {
		outputs = append(outputs, output)
	}

	return Transaction{tx.TXID, inputs, outputs}
}

// 分析校验
// 所需要的数据：公钥，数据(TxCopy, 生成哈希), 签名
// 要对每一个签名过的input进行校验

func (tx *Transaction) Verify(prevTxs map[string]Transaction) bool {
	if tx.IsCoinBaseTx() {
		return true
	}

	// 得到签名数据
	// 得到签名，反推r,s
	// 拆解Pubkey，x,y得到原生公钥
	// Verify

	txCopy := tx.TrimmedCopy()

	for i, input := range tx.TxInputs {
		prevTx := prevTxs[string(input.Txid)]
		if len(prevTx.TXID) == 0 {
			log.Panic("引用的交易无效")
		}

		txCopy.TxInputs[i].PubKey = prevTx.TxOutputs[input.Index].PubKeyHash

		txCopy.SetHash()
		dataHash := txCopy.TXID

		signature := input.Signature // 拆解r,s
		pubKey := input.PubKey       // 拆解x,y

		r := big.Int{}
		s := big.Int{}
		r.SetBytes(signature[:len(signature)/2])
		s.SetBytes(signature[len(signature)/2:])

		x := big.Int{}
		y := big.Int{}
		x.SetBytes(pubKey[:len(pubKey)/2])
		y.SetBytes(pubKey[len(pubKey)/2:])

		pubKeyOrigin := ecdsa.PublicKey{elliptic.P256(), &x, &y}

		if !ecdsa.Verify(&pubKeyOrigin, dataHash, &r, &s) {
			return false
		}
	}
	return true
}
