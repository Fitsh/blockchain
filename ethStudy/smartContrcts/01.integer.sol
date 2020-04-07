pragma solidity ^0.4.24;

contract Test {
    uint256 i256 = 100;
    
    int8    ui10  = 10;

    function  add()  public payable returns (uint256)  {  
        return i256 + uint256(ui10);
    }
    function isEqual() returns (bool) {
        return i256 == uint256(ui10);
    }
    function getbalance() public view returns(uint256) {
        return this.balance;
    }
}