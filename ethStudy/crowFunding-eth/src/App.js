import React, {Component} from 'react';
import CardExampleCard from './display/ui'
import 'semantic-ui-css/semantic.min.css'

import web3 from "./utils/initWeb3"
import lotteryInstance from "./eth/lotteryInstance"

//function App() {
//  return (
//      <p> Hello, World </p>
//  );
//}

console.log(lotteryInstance.methods.manager().call())

class App extends Component {
    constructor() {
        super()
        this.state = {
            manager: '',
            round: 0,
            winner: '',
            playersCount: 0,
            balance: 0,
            players: [],
            currentAccount: '',
            isClicked: false,
            isShowButton: '', // 控制管理员按钮是否显示
        }
    }

    // 内置钩子函数，在页面渲染之后调用
    componentDidMount() {

    }

    // 内置钩子函数，在页面渲染之前调用
    async componentWillMount() {
        // 获取当前的所有地址
        let accounts = await web3.eth.getAccounts()
        console.log(accounts)
        let _manager = await lotteryInstance.methods.manager().call()
        let _round = await lotteryInstance.methods.round().call()
        let _winner = await lotteryInstance.methods.winner().call()
        let _playersCount = await lotteryInstance.methods.getPlayersCount().call()

        //  单位wei，需要转换为ether
        let _balanceWei = await lotteryInstance.methods.getBalance().call()
        // 从wei单位转为ether单位
        let _balance = web3.utils.fromWei(_balanceWei, "ether")

        let _players = await lotteryInstance.methods.getPlayers().call()

        let isShowButton = accounts[0] === _manager ? 'inline':'none'

        this.setState({
            manager: _manager,
            round: _round,
            winner: _winner,
            playersCount: _playersCount,
            balance: _balance,
            players: _players,
            currentAccount: accounts[0],
            isClicked: false,
            isShowButton,
        })
    }

    // 卸载钩子函数
    

    play = async () => {
        // 处理真正的业务逻辑
        // 1. 调用play方法
        // 2. 转1 ether
        //
        this.setState({isClicked: true})
        try {
            await lotteryInstance.methods.play().send({
                from: this.state.currentAccount,
                //value: 1 * 10 ** 18,
                value: web3.utils.toWei('1', "ether"),
                gas: '3000000',
            })
            this.setState({isClicked: false})
            // 重新加载页面
            window.location.reload(true)
            alert("投注成功~")
        } catch(e) {
            this.setState({isClicked: false})
            alert("投注失败~")
            console.log(e)
        }
    }

    kaijiang = async () => {
        // 处理真正的业务逻辑
        //
        this.setState({isClicked: true})
        try {
            await lotteryInstance.methods.kaijiang().send({
                from: this.state.currentAccount,
                //value: 1 * 10 ** 18,
                //value: web3.utils.toWei('1', "ether"),
                gas: '3000000',
            })
            this.setState({isClicked: false})
            // 显示中奖人
            let winner = await lotteryInstance.methods.winner().call()
            window.location.reload(true)
            alert(`开奖成功~\n中奖人：${winner}`)
        } catch(e) {
            this.setState({isClicked: false})
            alert("开奖失败~")
            console.log(e)
        }
    }

    tuijiang = async () => {
        // 处理真正的业务逻辑
        //
        this.setState({isClicked: true})
        try {
            await lotteryInstance.methods.tuijiang().send({
                from: this.state.currentAccount,
                //value: 1 * 10 ** 18,
                //value: web3.utils.toWei('1', "ether"),
                gas: '3000000',
            })
            this.setState({isClicked: false})
            window.location.reload(true)
            alert("退奖成功~")
        } catch(e) {
            this.setState({isClicked: false})
            alert("退奖失败~")
            console.log(e)
        }
    }

    render() {
      return (
          <div>
              <CardExampleCard
               manager={this.state.manager} 
               round={this.state.round} 
               winner={this.state.winner} 
               playersCount={this.state.playersCount} 
               balance={this.state.balance} 
               players={this.state.players} 
               currentAccount={this.state.currentAccount} 
               play={this.play}
               kaijiang={this.kaijiang}
               tuijiang={this.tuijiang}
               isClicked={this.state.isClicked}
               isShowButton={this.state.isShowButton}
              />
          </div>
      );
    }
}

export default App;
