var Math = artifacts.require("./MathContract.sol");

module.exports = function(deployer) {
  deployer.deploy(Math);
};
