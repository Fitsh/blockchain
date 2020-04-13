pragma solidity ^0.4.24;

contract SimpleContract {
    string _message;
    
    constructor(string src) public {
        _message = src;
    }

    function setMessage(string message) public {
        _message = message;
    }

    function getMessage() public view returns(string) {
        return _message;
    }
}
