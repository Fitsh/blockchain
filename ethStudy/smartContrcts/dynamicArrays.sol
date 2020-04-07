pragma solidity ^0.4.24;

contract Test {
    // 创建方式之一：字面量形式
    uint[] public c  = [1, 2, 3];
    function h() public {
        c.push(4);
    }
    // 创建方式二： new关键字
    // a storage 类型数组，状态变量，最初为空，下标访问时越界
    uint[] public b;

    // 复杂类型在局部是引用
    function g() {
        b = new uint[](7);  // 7 为长度
        // 可以修改数组的长度
        b.length = 10;
        b[9] = 199;
        b.push(101);
    }

    // b memory类型数组
    function f() public returns (uint[]) {
        uint[] memory a = new uint[](7); // 7为长度
        // ERROR 不能修改长度
        // a.length = 100;
        // a.push(10);
        for (uint i = 0; i < a.length; i++) {
            a[i] = i;
        }
    }


    function getNumber() public view returns(uint[]) {
        return c;
    }
}
