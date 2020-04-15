let Web3 = require("web3")
//let HDWalletProvider = require("truffle-hdwallet-provider")

let web3 = new Web3()

console.log("window 的web3版本:", window.web3.version)

// 设置网络
// 使用用户自己的Provider来填充web3
let init = async () => {
    if(window.ethereum) {
        let provider = window.ethereum
        web3.setProvider(provider)

        try {
            await provider.enable()  // 异步操作
        } catch(error) {
            console.log("授权失败")
        }
    } else if (window.web3) {
        web3.setProvider(window.web3.currentProvider)
    } else {
        console.log("Non-Ethereum browser detected. You should consider trying MetaMask!")
    }
}

init()

console.log("我们的web3版本:", web3.version)
module.exports = web3
