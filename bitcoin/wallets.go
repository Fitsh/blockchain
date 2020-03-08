package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallets.dat"

// 定义一个wallets结构，他保存所有的wallet和它的地址
type Wallets struct {
	// map[地址]钱包
	WalletMap map[string]*Wallet
}

// 创建方法
func NewWallets() *Wallets {
	//	wallet := NewWallet()
	//	address := wallet.NewAddress()
	//
	var wallets Wallets
	wallets.WalletMap = make(map[string]*Wallet)
	//	wallets.WalletMap[address] = wallet
	wallets.loadFile()
	return &wallets
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	ws.WalletMap[address] = wallet
	ws.saveToFile()

	return address
}

// 保存方法，把它新建的wallet添加进去
func (ws *Wallets) saveToFile() {
	var buffer bytes.Buffer

	// gob: type not registered for interface: elliptic.p256Curve
	//panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256()) // 注册一个interface()对象

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	ioutil.WriteFile(walletFile, buffer.Bytes(), 0600)
}

// 读取文件方法， 把所有的wallet都读出来
func (ws *Wallets) loadFile() {
	// 在读取之前先确定文件是否存在， 如果不存在，直接退出
	_, err := os.Stat(walletFile)
	if os.IsNotExist(err) {
		return
	}
	content, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	// gob: type not registered for interface: elliptic.p256Curve
	//panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256()) // 注册一个interface()对象

	decoder := gob.NewDecoder(bytes.NewReader(content))
	var wsLocal Wallets
	err = decoder.Decode(&wsLocal)
	if err != nil {
		log.Panic(err)
	}
	// 对于结构体来说，里面有map的，要指定来赋值, 不要在最外层直接赋值
	ws.WalletMap = wsLocal.WalletMap
}

func (ws *Wallets) GetAllAddress() []string {
	var addresses []string
	for address := range ws.WalletMap {
		addresses = append(addresses, address)
	}
	return addresses
}
