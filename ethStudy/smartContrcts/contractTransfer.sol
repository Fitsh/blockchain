pragma  solidity ^0.4.24;
contract InfoFeed {
	function info() public payable returns(uint ret) {
		return 42;
	}
}
contract Consumer {
	InfoFeed feed;
	constructor()  public payable{
	    
	}
	function setFeed(address addr) public {
		feed = InfoFeed(addr);
	}
	function callFeed() public {
		feed.info.value(10).gas(800)();
	}
	function pay() public payable {
	    
	}
	function getBalance() public view returns(uint256) {
	    return address(this).balance;
	}
}