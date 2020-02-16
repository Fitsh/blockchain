package main

import (
    "math/big"
	"bytes"
	"crypto/sha256"
    "fmt"
)

// 定义一个工作量证明的结构ProofOfWork
type ProofOfWork struct {
	// a. block
	block *Block

	// b. 目标值
	// 一个比较大的数，他有丰富的方法：比较，复制
	target *big.Int
}

// 2. 定义创建POW的函数
//
//- NewProofOfWork(参数)
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	// 指定的难度值，现在是string,需要进行转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"

	// 引入的辅助变量，目的是将上面的难度值转为big.Int
	tmpInt := big.Int{}
	//将难度值复制给big.int，指定16进制格式
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

//
//3. 通过不断计算哈希的函数
//
//- Run()
func (pow *ProofOfWork) Run() ([]byte, uint64) {
    // 1. 拼装数据（区块的数据， 还有不断变化的随机数）
    // 2. 做哈希运算
    // 3. 与pow中的target比较
           // a. 找到了退出返回
           // b. 没找到，继续找，随机数加1

    var nonce uint64
    block := pow.block 
    var hash [32]byte

    for {
        // 1. 拼装数据（区块的数据， 还有不断变化的随机数）
        tmp := [][]byte{
            Uint64ToByte(block.Version),
            block.PrevHash,
            block.MerkelRoot,
            Uint64ToByte(block.TimeStamp),
            Uint64ToByte(block.Difficulty),
            Uint64ToByte(nonce),
            block.Data,
        }
        // 将二维的切片连接起来，返回一个一维的切片
        blockInfo := bytes.Join(tmp, []byte{})

        // 2. 做哈希运算
        // func Sum256(data []byte) [size]byte
        hash = sha256.Sum256(blockInfo)

        // 3. 与pow中的target比较
        tmpInt := big.Int{}
        // 将我们得到的哈希转换为一个big.Int
        tmpInt.SetBytes(hash[:])
        // 比较当前的哈希值和目标哈希值，如果当前哈希值小于目标哈希值，就说明创建区块成功

        //    -1  x < y
        //    0   x = y
        //    1   x > y
        // func (x *Int) Cmp(y *Int) (r int) {}
        if tmpInt.Cmp(pow.target) == -1 {
               // a. 找到了退出返回
               fmt.Printf("挖矿成功， hash: %x, nonce %d\n", hash, nonce)
               break
           } else {
               // b. 没找到，继续找，随机数加1
               nonce++
           }

    }
    return hash[:], nonce
}
//
//4. 提供一个校验函数
//- IsValid()
