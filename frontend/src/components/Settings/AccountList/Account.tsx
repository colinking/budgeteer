import React from 'react'
import { Pane, majorScale, Heading, Text, Strong, Badge } from 'evergreen-ui'
import { Account as PlaidAccount } from '../../../clients'
import { startCase, toLower, toUpper } from 'lodash'

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
          <Strong color="danger">${this.formatAmount(props.account.currentbalance)}</Strong>
          {" / "}
          <Text color="muted">${this.formatAmount(props.account.limit)}</Text>
        </Text>
      )
    } else {
      let amount = props.account.availablebalance
      if (amount === null) {
        amount = props.account.currentbalance
      }

      balance = <Strong color="green">${this.formatAmount(amount)}</Strong>
    }

    return (
      <Pane display="flex" padding={majorScale(2)} background="tint2" borderRadius={3} marginBottom={majorScale(1)} marginTop={majorScale(1)}>
        <Pane flex={1} alignItems="center" display="flex">
          <Heading size={500}>{this.formatName(props.account.officialname)}</Heading>
          <Badge marginLeft={majorScale(1)} color="purple">{props.account.subtype}</Badge>
          <Badge marginLeft={majorScale(1)} color="neutral">{props.account.mask}</Badge>
        </Pane>
        <Pane>
          {balance}
        </Pane>
      </Pane>
    )
  }

  protected formatAmount(amount: number): string {
    return amount.toFixed(2).replace(/\d(?=(\d{3})+\.)/g, '$&,')
  }

  protected formatName(name: string): string {
    if (toUpper(name) === name) {
      return startCase(toLower(name))
    }

    return name
  }
}

export default Account
