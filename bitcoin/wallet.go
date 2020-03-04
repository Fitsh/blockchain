package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	//	"github.com/mr-tron/base58"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// 每个钱包保存了公钥和私钥对

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	// pubKey  *ecdsa.PublicKey
	// 约定，这里的pubkey不存储原始的公钥，而是存储x和y拼接的字符串，在校验时拆分
	PubKey []byte
}

// 创建钱包
func NewWallet() *Wallet {
	// 创建曲线
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKeyOrig := privateKey.PublicKey

	pubKey := append(pubKeyOrig.X.Bytes(), pubKeyOrig.Y.Bytes()...)
	return &Wallet{
		privateKey,
		pubKey,
	}
}

// 生成地址
func (w *Wallet) NewAddress() string {
	pubKey := w.PubKey

	rip160HashValue := HashPubKey(pubKey)

	version := byte(00)
	payload := append([]byte{version}, rip160HashValue...)

	//checkSum
	checkCode := CheckSum(payload)

	//25字节数据
	payload = append(payload, checkCode...)
	//go语言有一个库，btcd, 这个是go语言实现的比特币全节点源码

	address := base58.Encode(payload)
	return address
}

func HashPubKey(pubKey []byte) []byte {
	hash := sha256.Sum256(pubKey)

	// 理解为编码器
	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}

	rip160HashValue := rip160hasher.Sum(nil)
	return rip160HashValue
}

func CheckSum(data []byte) []byte {
	//checkSum
	// 两次sha256
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])

	// 前四字节校验码
	checkCode := hash2[:4]

	return checkCode
}
