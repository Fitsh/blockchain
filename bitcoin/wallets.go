package main

// 定义一个wallets结构，他保存所有的wallet和它的地址
type Wallets struct {
	// map[地址]钱包
	WalletMap map[string]*Wallet
}

// 创建方法
func NewWallets() *Wallets {
	wallet := NewWallet()
	address := wallet.NewAddress()

	var wallets Wallets
	wallets.WalletMap = make(map[string]*Wallet)
	wallets.WalletMap[address] = wallet
	return &wallets
}

// 保存方法，把它新建的wallet添加进去

// 读取文件方法， 把所有的wallet都读出来
