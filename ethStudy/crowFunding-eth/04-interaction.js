//
//导入合约实例
let instance = require("./03-instance")

let Web3 = require("web3")

let HDWalletProvider = require("truffle-hdwallet-provider")

//let from = "0x795860597eC6D5a8100cF80Ee232984C523173eD"

//instance.methods.getMessage().call()
//.then(data => {
//    console.log("data: ", data)
//
//    instance.methods.setMessage("Hello, Hangzhou").send({
//        from: from,
//        value: 0,
//    }).then(res => {
//        console.log("res: ", res)
//    instance.methods.getMessage().call()
//    .then(data => {
//        console.log("data: ", data)
//        })
//    })
//})
//
let web3 = new Web3()

let terms = "soon cross else dentist video end rare cause only sail exist convince" // 助记词
let netIp = "https://ropsten.infura.io/v3/4a48e25f901c4807b8af85cef781c194"

let provider = new HDWalletProvider(terms, netIp)

web3.setProvider(provider)

// web3与区块链交互的返回值都是promise, 可以直接使用 async/await
//
let test = async () => {
    try {
        // 获取当前区块链对应助记词的账户地址
        let accounts = await web3.eth.getAccounts()
        let from = accounts[0]

        let v1 = await instance.methods.getMessage().call()
        console.log("data: ", v1)

        let v2 = await instance.methods.setMessage("Hello, ---------Hangzhou").send({
            from: from,
            gas: '3000000', //不要默认值，一定要写大些，使用单引号
            value: 0,
        })
        console.log(v2)
        let v3 = await instance.methods.getMessage().call()
            console.log("data: ", v3)
    }catch (e) {
        console.log(e)
    }
}


test()
