// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC777/ERC777.sol";

contract MyERC777 is ERC777 {
    constructor(address[] memory defaultOperators_)  ERC777("MyERC777","GGB" ,defaultOperators_){

    }
}