import React from 'react'
import { Item } from '../../../clients'
import Account from './Account'
import Empty from './Empty'

export interface AccountListProps {
  items: Item[]
}

class AccountList extends React.Component<AccountListProps> {
  public render() {
    const props = this.props

    if (props.items.length === 0) {
      return <Empty/>
    }

    return props.items.map(item => item.accountsList.map(account => <Account account={account} key={account.id}/>))
  }
}

export default AccountList
