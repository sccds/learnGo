pragma solidity >=0.4.21 <0.6.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/MathContract.sol";

contract TestMathContract {
    function testMulAToB() public {
        MathContract meta = MathContract(DeployedAddresses.MathContract());
        Assert.equal(meta.mulAToB(3, 4), 12, "3 * 4 should be 12");
    }
}
