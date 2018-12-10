import React from 'react'
import { Pane, majorScale, Heading, Text, Strong, Badge, Avatar, Tooltip, Position } from 'evergreen-ui'
import { Institution, Account as PlaidAccount } from '../../../clients'
import { formatAmount, formatName } from './lib/formatters'

export interface AccountProps {
  account: PlaidAccount
  institution?: Institution
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

    const institutionAvatarProps: {name?: string, src?: string} = {}
    if (props.institution) {
      institutionAvatarProps.name = props.institution.brandName
      if (props.institution.logo) {
        institutionAvatarProps.src = `data:image/png;base64,${props.institution.logo}`
      }
    }

    const truncateProps = {
      width: majorScale(35),
      whiteSpace: "nowrap",
      overflow: "hidden",
      textOverflow: "ellipsis"
    }

    return (
      <Pane display="flex" padding={majorScale(2)} background="tint2" borderRadius={3} marginBottom={majorScale(1)} marginTop={majorScale(1)}>
        <Pane flex={1} alignItems="center" display="flex">
          <Tooltip content={institutionAvatarProps.name} position={Position.LEFT}>
            <Avatar size={majorScale(4)} {...institutionAvatarProps} border="1px #333 solid" marginRight={majorScale(2)} />
          </Tooltip>
          <Tooltip content={formatName(props.account)} position={Position.RIGHT}>
            <Heading size={500} {...truncateProps}>{formatName(props.account)}</Heading>
          </Tooltip>
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
