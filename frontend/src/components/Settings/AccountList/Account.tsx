import React from 'react'
import { Pane, majorScale, Heading } from 'evergreen-ui'
import { Account as PlaidAccount } from '../../../clients'

export interface AccountProps {
  account: PlaidAccount
}

class Account extends React.Component<AccountProps> {
  public render() {
    const props = this.props

    return (
      <Pane display="flex" padding={majorScale(2)} background="tint2" borderRadius={3} marginBottom={majorScale(1)} marginTop={majorScale(1)}>
        <Pane flex={1} alignItems="center" display="flex">
          <Heading size={600}>{props.account.name}</Heading>
        </Pane>
        <Pane>
          <Heading size={600}>${props.account.currentbalance}</Heading>
        </Pane>
      </Pane>
    )
  }
}

export default Account
