pragma solidity ^0.4.24;

contract Test {
        mapping(address => uint256) public money;
    
    function pay() public payable {
        if (msg.value != 100 )  {
            throw;
        }
        money[msg.sender]=msg.value;
    }
    function getBalance() public view returns(uint256) {
        return address(this).balance;
        }
}