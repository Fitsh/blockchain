// 导入solc编译器
let solc = require("solc")
let fs = require("fs")

// 读取合约
var sourceCode = fs.readFileSync("./contracts/SimpleStorage.sol", "UTF-8")
var output = solc.compile(sourceCode, 1)

//console.log(output)
//console.log("abi: ", output['contracts'][":SimpleContract"]["interface"])
module.exports = output['contracts'][":SimpleContract"]
