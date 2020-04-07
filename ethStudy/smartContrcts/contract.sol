pragma solidity ^0.4.24;

contract C1 {
    uint256 public value ;
    constructor (uint256 input){
        value = input;
    }
    function getValue() public returns(uint256) {
        return value;
    }
}
contract C2{
    C1 public c1; //0x0000000000000
   C1 public c2; //0x00000000000
      C1 public c3; //0x0000000000
    
    function getValue() public   returns(uint256) {
        // 创建一个合约，返回地址
        address addr1 = new C1(10);
        // 需要在显示的转换为特定类型才能使用
        c1 = C1(addr1);
        return c1.getValue();
    }
    function getValue2() public returns(uint256) {
    // 定义合约的同时进行类型转换
        c2 = new C1(20);
        return c2.getValue();
    }

    function getValue3(address addr) public  returns(uint256) {
        c3 = new C1(30);
        return c3.getValue();
    }
    
}