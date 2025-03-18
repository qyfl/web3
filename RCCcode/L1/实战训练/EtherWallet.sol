// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract EtherWallet {
    // 拥有着
    address payable public immutable owner;
    event Log(string funName, address from, uint256 value, bytes data);

    constructor() {
        owner = payable(msg.sender);
    }

    // 检查是否有权限
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    // 收到转账
    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }

    function withdraw1() external onlyOwner {
        // owner.transfer 相比 msg.sender 更消耗Gas
        // owner.transfer(address(this).balance);
        payable(msg.sender).transfer(100);
    }

    function withdraw2() external onlyOwner {
        bool success = payable(msg.sender).send(200);
        require(success, "Send Failed");
    }

    function withdraw3() external onlyOwner {
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Call Failed");
    }

    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
