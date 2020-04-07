pragma solidity ^0.4.24;

contract Test {
    string public owner = "";
constructor () payable public  {
}

function setOwner(string o) public payable {
    owner=o;
}
    function getbalance() public view returns(uint256) {
        return address(this).balance;
    }
}