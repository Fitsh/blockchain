pragma solidity ^0.4.24;

contract Test {
    enum Weekdays  {
        Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
    }
    
    Weekdays currentDay;
    Weekdays defaultDay = Weekdays.Sunday;
    
    function setDay(Weekdays _day)  public {
        currentDay = _day; 
    }
    function getDay() public view returns (Weekdays) {
        return currentDay;
    }
    function getDafultDay() public view returns(Weekdays) {
        return defaultDay;
    }
}