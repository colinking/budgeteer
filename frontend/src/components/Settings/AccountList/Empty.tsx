import React from 'react'
import { Pane, Text, majorScale } from 'evergreen-ui'
import withUser, { UserProviderProps } from '../../withUser'

export interface EmptyAccountsProps extends UserProviderProps {}

class EmptyAccounts extends React.Component<EmptyAccountsProps> {
  public render() {
    return (
      <Pane marginBottom={majorScale(2)}>
        <Text>No Associated Accounts</Text>
      </Pane>
    )
  }
}

export default withUser(EmptyAccounts)
