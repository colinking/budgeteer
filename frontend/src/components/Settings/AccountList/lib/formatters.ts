import { Account } from '../../../../clients'
import { startCase, toLower, toUpper } from 'lodash'

export function formatAmount(amount: number): string {
  return amount.toFixed(2).replace(/\d(?=(\d{3})+\.)/g, '$&,')
}

export function formatName(account: Account): string {
  const name = account.name.length > account.officialname.length ? account.name : account.officialname
  if (toUpper(name) === name) {
    return startCase(toLower(name))
  }

  return name
}
