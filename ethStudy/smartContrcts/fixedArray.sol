pragma solidity ^0.4.24;

contract Test {
    bytes1 b1="h";
    bytes20 b10 ="hello, world";
    function getlength() public view returns(uint256) {
        return b10.length;
    }
    function setValue() public pure {
        
    }
    function getValue(uint256 i) public view returns(byte) {
        return b10[i];
    }
}