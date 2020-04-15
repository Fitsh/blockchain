pragma solidity ^0.4.24;

contract Lottery {
    // 1. 管理员，负责开奖和退奖
    // 2. 彩民池， address[] players
    // 3. 当前期数，每期结束后加1， round
    address public manager;
    address[] public players;
    uint256 public round;
    address public winner;
    
    constructor() public {
        manager = msg.sender;
    }

// 每个人可以投多次，但每次只能投1ether
    function play() public payable {
        require(msg.value == 1 ether);
        // 把参与者添加到彩民池中
        players.push(msg.sender);
    }
    
    function getBalance() public view returns(uint256) {
        return address(this).balance;
    }    
    
    function getPlayers() public view returns(address[] ) {
        return players;
    }
    // 开奖函数
    // 从彩民池（数组）中找一个随机彩民（一个随机数）
    // 找一个特别大的随机数，对我们的彩民数组长度求余数
    // 用哈希数值实现大的随机数
    // 哈希内容的随机：当前时间，区块的挖矿难度，彩民数量作为输入
    // bytes memory v1 = abi.encodePacked(block.timestamp,block.difficulty,players.length);
    // bytes32 v2 = keccak256(v1);
    // uint256 v3 = uint256(v2);
    function kaijiang() onlyManager  public {
        bytes memory v1 = abi.encodePacked(block.timestamp, block.difficulty, players.length);
        bytes32 v2 = keccak256(v1);
        uint256 v3 = uint256(v2);
        uint256 index = v3 % players.length;
        winner = players[index];
        
        uint256 money = address(this).balance * 90 /100;
        uint256 money1 = address(this).balance - money;
        
        winner.transfer(money);
        manager.transfer(money1);
        
        round++;
        
        delete players;
    }
    
    // 退奖逻辑：
    // 遍历players数组，逐一退款1ether
    // 期数加1
    // 彩民池清零
    // 调用者花费手续费
    function tuijiang() onlyManager public {
        for (uint256 i = 0; i < players.length; i++ ) {
            players[i].transfer(1 ether);
        }
        round++;
        delete players;
    }
    
    modifier onlyManager {
        require(msg.sender == manager);
        _;
    }
    
    function getPlayersCount() public view returns(uint256) {
        return players.length;
    }
}
