pragma solidity ^0.4.24;

contract Test {
    bytes public names;
    function getLen() public  view returns(uint256) {
        return names.length;
    }
    function setValue(bytes input) public {
        names = input;
    }
    function getByIndex(uint256 i) public view returns (byte) {
        return names[i];
   }
   function setlen(uint256 len) public {
       names.length = len;
   }
   function setValue2(uint256 i) public {
       names[i]='h';
   }
   function PushData() public {
       names.push("a");
   }
}