package main

import (
	"log"

	"./bolt"
)

type BlockChainIterator struct {
	db *bolt.DB
	// 游标， 用于不断索引
	currentHashPointer []byte
}

// func NewBlockChainIterator() {}

func (bc *BlockChain) NewBlockChainIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db: bc.db,
		// 最初指向区块链的最后一个区块，随着next调用，不断变化
		currentHashPointer: bc.tail,
	}
}

// 迭代器是区块链的
// Next方式是属于迭代器的
// 1. 返回当前的区块
// 2. 指针前移
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("Bucket 不应该为空， 清检查!")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		//解码
		block = Deserialize(blockTmp)
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}
