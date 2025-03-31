// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721Enumerable.sol";

contract Collectible is ERC721, ERC721Metadata, ERC721Enumerable {
    constructor() ERC721Metadata("My Collectible", "MCL") public {}
    // 创建新的 NFT
    function mintCollectible(address to, uint256 tokenId, string memory uri) public {
        _mint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }
}