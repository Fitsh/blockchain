let {bytecode, interface} = require("./01-compile")

let Web3 = require("web3")

let HDWalletProvider = require("truffle-hdwallet-provider")

//console.log(bytecode)
//
//console.log(interface)
//
// 1 引入web3
// 2 new 一个web3实例
// 3 设置网络
//
let web3 = new Web3()

let terms = "soon cross else dentist video end rare cause only sail exist convince" // 助记词
let netIp = "https://ropsten.infura.io/v3/4a48e25f901c4807b8af85cef781c194"

let provider = new HDWalletProvider(terms, netIp)

//web3.setProvider("http://127.0.0.1:7545")
web3.setProvider(provider)

console.log("web version: ",web3.version)
//console.log(web3.currentProvider)

//const account = "0xe6EaF095Ab5fcc761032adc9746A0EC693E73950"

// 拼接合约interface
let contract  = new web3.eth.Contract(JSON.parse(interface))

//// 拼接bytecode


// 合约部署
//contract.deploy({
//    data: bytecode,  // 合约的bytecode
//    arguments: ["helloworld"],  // 给构造函数传递参数，使用数组
//}).send({
//    from: account,
//    gas: 1500000,
//    gasPrice: '30000000000000',
//})
//.then(instance => {
//    console.log('address: ', instance.options.address)
//})

let deploy =  async () => {
    // 先获取所有账户

    let accounts = await web3.eth.getAccounts()
    console.log("accounts: ", accounts)
    //
    // 执行部署
    let instance = await contract.deploy({
        data: bytecode,
        arguments: ["helloworld"],
    }).send({
        from: accounts[0],
        gas: 3000000,
//        gasPrice: '30000000000000000',
    })

    console.log("instance address: ", instance.options.address)
}

deploy()
