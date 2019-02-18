pragma solidity >=0.4.21 <0.6.0;

contract Greetings {
    string private message = "欢迎来到智能合约世界";
    function sayHello() public view returns (string memory) {
        return message;
    }
}
