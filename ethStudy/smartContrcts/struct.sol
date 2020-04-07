pragma solidity ^0.4.24;

contract Test {
    struct Student  {
        string name;
        uint age;
        uint score;
        string sex;
    }
    Student[] public Students;
    Student  public stu1  = Student("lily", 18, 20, "girl");
    Student public stu2 = Student({name:"Ji米", age:10, score:123, sex: "boy"});
    function assign() public {
        Students.push(stu1);
        Students.push(stu2);
        stu1.name = "Lily";
    }
    function returnStudent() public view returns(string, uint, uint, string){
        return (stu1.name, stu1.age, stu1.score, stu1.sex);
    }
}