// 1 引入web3
let Web3 = require("web3")

let HDWalletProvider = require("truffle-hdwallet-provider")

// 2 new 一个web3实例
let web3 = new Web3()

let terms = "soon cross else dentist video end rare cause only sail exist convince" // 助记词
let netIp = "https://ropsten.infura.io/v3/4a48e25f901c4807b8af85cef781c194"

let provider = new HDWalletProvider(terms, netIp)

// 3 设置网络
//web3.setProvider("http://127.0.0.1:7545")
web3.setProvider(provider)

let abi = [{"constant":false,"inputs":[{"name":"message","type":"string"}],"name":"setMessage","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getMessage","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"src","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]
//let address = "0x559f9a90aC6eDEACB18295282aE7eD565623fBA9"
let address = "0x0DAb718F6C0c1F9268B0b6051DB1aeceAfF5b7C5"

// 此处abi已经是json对象不需要进行parse动作
let contractInstance = new web3.eth.Contract(abi, address)

console.log("address: ", contractInstance.options.address)

module.exports = contractInstance
