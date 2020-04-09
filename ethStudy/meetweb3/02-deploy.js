let {bytecode, interface} = require("./01-compile")

//console.log(bytecode)
//
//console.log(interface)
//
// 1 引入web3
// 2 new 一个web3实例
// 3 设置网络
//
let Web3 = require("web3")

let web3 = new Web3()

web3.setProvider("http://127.0.0.1:7545")

console.log("web version: ",web3.version)
//console.log(web3.currentProvider)

const account = "0xe6EaF095Ab5fcc761032adc9746A0EC693E73950"

// 拼接合约interface
let contract  = new web3.eth.Contract(JSON.parse(interface))

// 拼接bytecode
contract.deploy({
    data: bytecode,  // 合约的bytecode
    arguments: ["helloworld"],  // 给构造函数传递参数，使用数组
}).send({
    from: account,
    gas: 1500000,
    gasPrice: '30000000000000',
})
.then(instance => {
    console.log('address: ', instance.options.address)
})


// 合约部署
