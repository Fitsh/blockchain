pragma solidity ^0.4.24;

contract Test {
    address public addr0 = 0x00dd870fa1b7c4700f2bd7f44238821c26f7392148;
    address public addr1 = 0x00dd870fa1b7c4700f2bd7f44238821c26f7392148;

constructor() payable public {
     
 }
  function add() public view returns(uint160) {
      return uint160(addr1)+ uint160(10);
  }
    function getbalance() public view returns(uint256) {
        return addr1.balance;
    }
    function getContractBalance() public view returns(uint256) {
        return address(this).balance;
    }
    function tranfer ()  public payable {
        addr1.transfer(10 * 10 ** 18);
    }
}