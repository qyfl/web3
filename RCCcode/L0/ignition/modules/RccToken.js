// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { ethers, upgrades } = require("hardhat");

module.exports = buildModule("RccTokenModule", (m) => {
  const rccToken = m.contract("RccToken");
  return { rccToken };
});
