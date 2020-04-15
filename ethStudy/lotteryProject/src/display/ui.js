import React from 'react'
import { Card, Icon, Image, Segment, Statistic, Button } from 'semantic-ui-react'

//https://react.semantic-ui.com/views/statistic/#variations-inverted
const CardExampleCard = (props) => (
  <Card>
    <Image src='/images/logo.jpg' wrapped ui={false} />
    <Card.Content>
      <Card.Header>福利彩票</Card.Header>
      <Card.Meta>
        <p> 管理员地址：{props.manager} </p>
        <p> 当前地址：{props.currentAccount} </p>
        <p> 上期中奖地址：{props.winner} </p>
      </Card.Meta>
      <Card.Description>
        每晚八点开奖，不见不散
      </Card.Description>
    </Card.Content>
    <Card.Content extra>
      <a>
        <Icon name='user' />
    {props.playersCount} 人参与
      </a>
    </Card.Content>

    <Card.Content extra>
       <Statistic color='red' >
          <Statistic.Value>{props.balance} ETH</Statistic.Value>
          <Statistic.Label>奖金池</Statistic.Label>
        </Statistic>
    </Card.Content>

    <Card.Content extra>
       <Statistic color='yellow' >
          <Statistic.Value>第{props.round}期</Statistic.Value>
          <a href="https://ropsten.etherscan.io/address/0xb3fd6fc9bb1a236ccc4a2c086052a1236eb087f1"> 点击我查看历史交易 </a>
        </Statistic>
    </Card.Content>
    <Button animated='fade' color='orange' onClick={props.play} disabled={props.isClicked}>
      <Button.Content visible>投注产生希望</Button.Content>
      <Button.Content hidden>购买放飞梦想</Button.Content>
    </Button>
    <Button  inverted color='red' style={{display: props.isShowButton }} onClick={props.kaijiang} disabled={props.isClicked}>
        开奖
    </Button>
    <Button inverted color='orange' style={{display: props.isShowButton }} onClick={props.tuijiang} disabled={props.isClicked}>
        退奖
    </Button>
  </Card>
)

export default CardExampleCard
// import es6
