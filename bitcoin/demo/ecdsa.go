package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

// ecdsa生成私钥公钥
// 签名校验

func main() {
	// 创建曲线
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := privateKey.PublicKey

	data := "hello world~"
	hash := sha256.Sum256([]byte(data))

	//签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%x\n", pubKey)
	fmt.Printf("%x\n", r.Bytes())
	fmt.Printf("%x\n", s.Bytes())
	// 把r, s 进行序列化进行传输
	signature := append(r.Bytes(), s.Bytes()...)

	//校验需要三样东西，数据，签名，公钥

	r1 := big.Int{}
	s1 := big.Int{}

	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])
	//func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
	res := ecdsa.Verify(&pubKey, hash[:], &r1, &s1)
	fmt.Printf("校验结果%v\n", res)
}
