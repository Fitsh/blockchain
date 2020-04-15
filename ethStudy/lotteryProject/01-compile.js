// 导入solc编译器
let solc = require("solc")
let fs = require("fs")

// 读取合约
var sourceCode = fs.readFileSync("./contracts/lottery.sol", "UTF-8")
var output = solc.compile(sourceCode, 1)

//console.log(output)
//console.log("abi: ", output['contracts'][":Lottery"]["interface"])
module.exports = output['contracts'][":Lottery"]
