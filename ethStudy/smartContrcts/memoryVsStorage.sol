pragma solidity ^0.4.24;

contract Test {
    string public name="lily";
    uint256 public num = 10;
    function call1() public {
        setName(name);
    }
    function setName(string input) private {
        bytes(input)[0]='H';
        num=1;
    }
    function call2() {
        setName2(name);
    }
    function setName2(string storage input) private {
                bytes(input)[0]='H';
        num=2;
    }
    function localTest() public {
    // string storage tmp = name;
        string tmp = name;
        num = 3;
        bytes(tmp)[0]="3";
    }
        function localTest1() public {
        string memory tmp = name;
        num = 3;
        bytes(tmp)[0]="3";
    }
}