pragma solidity ^0.4.24;

contract Test {
    address public owner;
    uint256 a;
    address public caller;
    constructor() {
        owner = msg.sender;
    }
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
    function setValue(uint256 tmp ) onlyOwner {
        a =tmp;
        caller = msg.sender;
    }
}