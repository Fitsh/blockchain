pragma solidity ^0.4.24;

contract Test {
    address public addr1 = 0x00dd870fa1b7c4700f2bd7f44238821c26f7392148;

  function add() public view returns(uint160) {
      return uint160(addr1)+ uint160(10);
  }
    function getbalance() public view returns(uint256) {
        return addr1.balance;
    }
}