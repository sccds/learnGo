pragma solidity >=0.4.21 <0.6.0;

contract Voting {

    // ["zhangsan", "lisi", "wangwu", "zhaoliu"]
    // ["0x7a68616e6773616e", "0x6c697369", "0x77616e677775", "0x7a68616f6c6975"]

    bytes32[] public candidateList;
    mapping (bytes32 => uint8) votesReceived;

    // 构造函数 初始化候选人名单
    constructor(bytes32[] memory candidateNames) public {
        for(uint i=0; i<candidateNames.length; i++) {
            candidateList = candidateNames;
        }
    }

    // 查询某某选人的总票数
    function totalVotesFor(bytes32 candidate) view public returns (uint) {
        require(validCandidate(candidate) == true);
        return votesReceived[candidate];
    }

    // 为某候选人投票
    function votingForCandidate(bytes32 candidate) public {
        require(validCandidate(candidate) == true);
        votesReceived[candidate] += 1;
    }

    // 检查候选人名字是否有效
    function validCandidate(bytes32 candidate) view public returns (bool) {
        for (uint i=0; i<candidateList.length; i++) {
            if (candidateList[i] == candidate) {
                return true;
            }
        }
        return false;
    }
}
