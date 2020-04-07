pragma solidity ^0.4.24;

contract Test {
    address public owner;
    uint256 a;
    address public caller;
    constructor() {
        owner = msg.sender;
    }
    function setValue(uint256 tmp ) {
        a =tmp;
        caller = msg.sender;
    }
}