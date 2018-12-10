import React from 'react'
import { Pane, majorScale, Heading, Text, Strong, Badge } from 'evergreen-ui'
import { Account as PlaidAccount } from '../../../clients'
import { formatAmount, formatName } from './lib/formatters'

export interface AccountProps {
  account: PlaidAccount
}

class Account extends React.Component<AccountProps> {
  public render() {
    const props = this.props

    let balance = null
    if (props.account.subtype === 'credit card') {
      balance = (
        <Text>
          <Strong color="danger">${formatAmount(props.account.currentbalance)}</Strong>
          {" / "}
          <Text color="muted">${formatAmount(props.account.limit)}</Text>
        </Text>
      )
    } else {
      balance = <Strong color="green">${formatAmount(props.account.currentbalance)}</Strong>
    }

    return (
      <Pane display="flex" padding={majorScale(2)} background="tint2" borderRadius={3} marginBottom={majorScale(1)} marginTop={majorScale(1)}>
        <Pane flex={1} alignItems="center" display="flex">
          <Heading size={500}>{formatName(props.account)}</Heading>
          <Badge marginLeft={majorScale(1)} color="purple">{props.account.subtype}</Badge>
          <Badge marginLeft={majorScale(1)} color="neutral">{props.account.mask}</Badge>
        </Pane>
        <Pane>
          {balance}
        </Pane>
      </Pane>
    )
  }
}

export default Account
