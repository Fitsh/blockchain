pragma solidity ^0.4.24;

contract Test {
  bytes10 b10 = 0x68656c6c6f776f726c64; // hello,world
  // bytes bs10 = b10; //无法直接转换
  bytes public bs10 = new bytes(b10.length);
  
  // 固定字节数组转动态字节数组
  function fixedBytesToBytes() public  {
      for(uint i = 0 ; i < b10.length; i++) {
      			bs10[i] = b10[i];
      }
  }
  function BytesToString() public view  returns (string) {
      return string(bs10);
  }
  function stringToBytes(string str) public  view returns (bytes){
      return bytes(str);
  }
}