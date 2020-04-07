pragma solidity ^0.4.24;

contract Test {
    // id => name
    mapping(uint => string) public id_names;
    constructor()  public {
    id_names[1]="lily";
    id_names[2]="jim";
    id_names[3]="Tom";
    }
    function getNameById(uint id) public view returns(string) {
        string memory name = id_names[id];
        return name;
    }
    function setNameById(uint id) public  returns(string) {
        id_names[id]="Hello";
    }
}