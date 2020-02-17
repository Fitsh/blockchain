package main

import (
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
func NewBlockChain() *BlockChain {
	// 创建一个创始块，并作为第一个区块添加到区块链中
	//  genesisBlock := GenesisBlock()
    //	return &BlockChain{
    //		blocks: []*Block{genesisBlock},
    //	}
    db, err := bolt.Open(blockChainDb, 0600, nil)
    if err != nil {
        log.Panic("打开数据库失败!")
    }
    defer db.Close() 
    
    var lastHash []byte

    db.Update(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte(blockBucket))
        if bucket == nil {
            bucket, err = tx.CreateBucket([]byte(blockBucket))
            if err != nil {
                log.Panic("创建bucket失败")
            }

            genesisBlock := GenesisBlock()

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

    return &BlockChain {
        db: db,
        tail: lastHash,
    }
}

// 定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("创世块", []byte{})
}

// 5.添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 获取最后一个区块
//	lastBlock := bc.blocks[len(bc.blocks)-1]
//	prevHash := lastBlock.Hash
    db:= bc.db
    lastHash := bc.tail
//
//	// a 创建区块
//	// b 添加到区块链数组中
//	block := NewBlock(data, prevHash)
//	bc.blocks = append(bc.blocks, block)
    
    block := NewBlock(data, lastHash)

    db.Update(func(tx *bolt.Tx) error {

        bucket := tx.Bucket([]byte(blockBucket))
        if bucket == nil {
            log.Panic("bucket 不应该不存在")
        }
        bucket.Put(block.Hash, block.Serialize())
        bucket.Put([]byte("LastHashKey"), block.Hash)

        bc.tail = block.Hash
        return nil
    })
}
