// 1 引入web3
let Web3 = require("web3")

// 2 new 一个web3实例
let web3 = new Web3()

// 3 设置网络
web3.setProvider("http://127.0.0.1:7545")

let abi = [{"constant":false,"inputs":[{"name":"message","type":"string"}],"name":"setMessage","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getMessage","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"src","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]
let address = "0x559f9a90aC6eDEACB18295282aE7eD565623fBA9"

// 此处abi已经是json对象不需要进行parse动作
let contractInstance = new web3.eth.Contract(abi, address)

console.log("address: ", contractInstance.options.address)

module.exports = contractInstance
