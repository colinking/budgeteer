import * as React from 'react'

import { users, User, ServiceError } from '../../clients'
import SettingsComponent from './Settings'
import withUser, { UserProviderProps } from '../withUser'
import Loader from '../Loader'

const plaidEnv = process.env.REACT_APP_PLAID_ENV as string
const plaidPublicKey = process.env.REACT_APP_PLAID_PUBLIC_KEY as string

interface SettingsProps extends UserProviderProps {}

class Settings extends React.Component<SettingsProps> {
  public handleOnLinkExit = async (err: Error | undefined) => {
    if (err) {
      this.handleLinkError(err)
    }
  }

  public handleLinkError = async (err: Error | ServiceError) => {
    console.error(err)
  }

  public handleOnLinkSuccess = async (token: string, metadata: any) => {
    await users.addItem(req => {
      req.setToken(token)  
    })

    this.props.refetchUser()
  }

  public render() {
    const props = this.props
    console.log(this.props.user)

    if (!props.user) {
      return <Loader/>
    }

    return (
      <SettingsComponent
        handleOnLinkExit={this.handleOnLinkExit}
        handleOnLinkSuccess={this.handleOnLinkSuccess}
        plaidEnv={plaidEnv}
        plaidPublicKey={plaidPublicKey}
        user={props.user}
      />
    )
  }
}

export default withUser(Settings)
