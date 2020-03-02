package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
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
