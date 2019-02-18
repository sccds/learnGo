import React, { Component } from "react";
import VotingContract from "./contracts/Voting.json";
import getWeb3 from "./utils/getWeb3";


import "./css/oswald.css"
import "./css/open-sans.css"
import "./css/pure-min.css"
import "./App.css";

// 合约地址
const contractAddr = "0xf7693e8252cc43ef608150c7ad9011d6a677ff11";
var votingContractInstance;
var account;

var _modifyVotingCount = (candidates, i, votingCount) => {
    console.log("-----------");
    console.log(candidates);
    console.log(i);
    console.log(votingCount);
    let obj = candidates[i];
    obj.votingCount = votingCount;
    return candidates;
}

class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            candidates: [
                {
                    name: "zhangsan",
                    byte32Name: "0x7a68616e6773616e",
                    votingCount: 0,
                    id: 100
                },
                {
                    name: "lisi",
                    byte32Name: "0x6c697369",
                    votingCount: 0,
                    id: 101
                },
                {
                    name: "wangwu",
                    byte32Name: "0x77616e677775",
                    votingCount: 0,
                    id: 102
                },
                {
                    name: "zhaoliu",
                    byte32Name: "0x7a68616f6c6975",
                    votingCount: 0,
                    id: 103
                }
            ],
            candidatesVoteCount: ["0", "0", "0", "0"],
            web3: null
        }
    }

    componentWillMount() {
        getWeb3
        .then(results => {
            this.setState({
                web3: results.web3
            }, () => {
                this.instantiateContract()
            })
        })
        .catch(() => {
            console.log("Error finding web3")
        })
    }

    instantiateContract() {
        const contract = require('truffle-contract')
        const votingContract = contract(VotingContract)
        votingContract.setProvider(this.state.web3.currentProvider)

        // get accounts
        this.state.web3.eth.getAccounts((error, accounts) => {
            votingContract.at(contractAddr).then((instance) => {
                account = accounts[0];
                votingContractInstance = instance;
                for (let i = 0; i < this.state.candidates.length; i++) {
                    let object = this.state.candidates[i];
                    console.log(accounts[0]);
                    console.log(votingContractInstance);
                    console.log(votingContractInstance.totalVotesFor(object.byte32Name));
                    votingContractInstance.totalVotesFor(object.byte32Name).then(result => {
                        console.log("vote for ", i);
                        console.log(result["words"]["0"]);
                        this.setState({
                            candidates: _modifyVotingCount(this.state.candidates, i, result["words"]["0"])
                        });
                    });
                }
            })
        })
    }

  render() {
    return (
      <div className="App">
        Voting:
        <ul>
            { this.state.candidates.map((object) => {
                    console.log(object);
                    return (
                        <li key={object.id}>候选人: {object.name}  支持票数: {object.votingCount}</li>
                    )
                })
            }
        </ul>

        <input ref="candidateInput" style={{width:200,height:40,borderWidth:2,marginLeft:30}}/>
        <button onClick={ () => {
            let candidateName = this.refs.candidateInput.value;
            console.log("voting to person " + candidateName );
            console.log("account: ", account);
            votingContractInstance.votingForCandidate(this.state.web3.utils.fromAscii(candidateName), {from: account})
            .then((result) => {
                    console.log(result);
                    console.log(candidateName);
                    for (let i = 0; i < this.state.candidates.length; i++) {
                        let object = this.state.candidates[i];
                        if (object.byte32Name === this.state.web3.utils.fromAscii(candidateName)) {
                            votingContractInstance.totalVotesFor(object.byte32Name).then(result => {
                                this.setState({
                                    candidates: _modifyVotingCount(this.state.candidates, i, result["words"]["0"])
                                });
                            });
                            break;
                        }
                    }

                });
        }}>Vote</button>
      </div>
    )
  }
}

export default App;
