pragma solidity ^0.4.24;

contract Test {
  string public name = "Lily";
  function nameBytes() public constant returns(bytes) {
      return bytes(name);
  }
  function nameLength() public constant returns(uint256) {
      return bytes(name).length;
  }
  function changeName() public {
      bytes(name)[0]='H';
  }
  function changeLength()  public {
      bytes(name).length=15;
      bytes(name)[14]='x';
  }
}