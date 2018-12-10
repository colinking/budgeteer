import React from 'react'
import { Item, Account as PlaidAccount } from '../../../clients'
import Account from './Account'
import Empty from './Empty'
import { formatName } from './lib/formatters'

export interface AccountListProps {
	items: Item[]
}

class AccountList extends React.Component<AccountListProps> {
	public render() {
		const props = this.props

		if (props.items.length === 0) {
			return <Empty />
		}

		const accounts = ([] as Array<{ account: PlaidAccount, item: Item }>).concat(
			...props.items.map((item) => item.accountsList.map((account) => ({ account, item })))
		)

		// Sort accoutns by formatted name.
		accounts.sort(({account: a}, {account: b}) => {
			const nameA = formatName(a)
			const nameB = formatName(b)
			if (nameA === nameB) {
				return 0
			}
			return nameA > nameB ? 1 : -1
		})

		return accounts.map(({ account, item }) => (
			<Account account={account} key={account.id} institution={item.institution} />
		))
	}
}

export default AccountList
