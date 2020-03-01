package main

import (
	"bytes"
	"fmt"
	"log"

	"./bolt"
)

// 4.引入区块链
//   BlockChain结构重写
//   使用数据库代替数组

type BlockChain struct {
	// 定一个区块链数组
	//	blocks []*Block
	db *bolt.DB

	tail []byte // 存储最后一块的哈希
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

// 5. 定义一个区块链
func NewBlockChain(address string) *BlockChain {
	// 创建一个创始块，并作为第一个区块添加到区块链中
	//  genesisBlock := GenesisBlock()
	//	return &BlockChain{
	//		blocks: []*Block{genesisBlock},
	//	}
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败!")
	}

	var lastHash []byte

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket失败")
			}

			genesisBlock := GenesisBlock(address)
			log.Printf("gensisBlock %s\n", genesisBlock)

			// 哈希作为key, block的字节流作为value
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			// // 测试
			// blockBytes := bucket.Get(genesisBlock.Hash)
			// block := Deserialize(blockBytes)
			// fmt.Printf("decode %s\n", block)
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})

	return &BlockChain{
		db:   db,
		tail: lastHash,
	}
}

// 定义一个创世块
func GenesisBlock(address string) *Block {
	//return NewBlock("创世块", []byte{})
	coinBase := NewCoinBaseTx(address, "创始块")
	return NewBlock([]*Transaction{coinBase}, []byte{})
}

// 5.添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	// 获取最后一个区块
	//	lastBlock := bc.blocks[len(bc.blocks)-1]
	//	prevHash := lastBlock.Hash
	db := bc.db
	lastHash := bc.tail
	//
	//	// a 创建区块
	//	// b 添加到区块链数组中
	//	block := NewBlock(data, prevHash)
	//	bc.blocks = append(bc.blocks, block)

	block := NewBlock(txs, lastHash)

	err := db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该不存在")
		}
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)

		bc.tail = block.Hash

		return nil
	})
	if err != nil {
		log.Printf("----------------------err %v\n", err)
		log.Panic(err)
	}
}

func (bc *BlockChain) PrintChain() {
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		// assunme bucket exist and has key
		b := tx.Bucket([]byte("blockBucket"))

		// 从第一个key-value 进行遍历，到最后一个key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := Deserialize(v)
			fmt.Printf("============= 区块高度: %d ============\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("前区块哈希: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值: %d\n", block.Difficulty)
			fmt.Printf("随机数: %d\n", block.Nonce)
			fmt.Printf("当前区块哈希: %x\n", block.Hash)
			fmt.Printf("数据: %s\n\n", block.Transactions[0].TxInputs[0].Sig)

			return nil
		})
		return nil
	})
}

// 找到指定地址的所有的utxo
func (bc *BlockChain) FindUTXOs(address string) []TxOutput {
	var UTXO []TxOutput
	// map[交易Id][]int64
	spentOutputs := make(map[string][]int64)
	// 遍历区块
	// 遍历交易
	// 遍历output,找到和自己相关的utxo（在添加output之前，检查一下是否已经消耗过）
	// 遍历input，找到自己花费过的input集合（把自己消耗过的标识出来）

	// 创建迭代器
	it := bc.NewBlockChainIterator()
	for {
		// 遍历区块
		block := it.Next()

		// 遍历交易
	OUTPUT:
		for _, tx := range block.Transactions {
			log.Printf("current txid :%x\n", tx.TXID)
			// 遍历output,找到和自己相关的utxo（在添加output之前，检查一下是否已经消耗过）
			for i, output := range tx.TxOutputs {
				log.Printf("current index :%d\n", i)
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}
				if output.PubKeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			if !tx.IsCoinBaseTx() {
				// 如果当前交易是挖矿交易，那么不做遍历，直接跳过
				for _, input := range tx.TxInputs {
					// 遍历input，找到自己花费过的input集合（把自己消耗过的标识出来）
					if input.Sig == address {
						indexArray := spentOutputs[string(tx.TXID)]
						indexArray = append(indexArray, input.Index)
					}
				}
			}
		}
		if len(block.PrevHash) == 0 {
			log.Printf("区块遍历完成退出\n")
			break
		}
	}
	return UTXO
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]int64, float64) {
	//找到的合理的utxo集合
	var utxos map[string][]int64
	// 找到的utxo里包含的总数
	var calc float64

	//TODO
	return utxos, calc
}
