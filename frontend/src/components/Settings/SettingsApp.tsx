import * as React from 'react'

import clients from '../../lib/clients'
import { getMetadata } from '../../lib/requests'
import SettingsComponent from './Settings'
import { AddItemRequest, AddItemResponse } from '../../gen/userpb/user_service_pb';
import { ServiceError } from '../../gen/userpb/user_service_pb_service'

const plaidEnv = process.env.REACT_APP_PLAID_ENV as string
const plaidPublicKey = process.env.REACT_APP_PLAID_PUBLIC_KEY as string

interface SettingsProps {}

export default class Settings extends React.Component<SettingsProps> {
  public handleOnLinkExit(err: Error | undefined) {
    console.log('User exited Link.')
    if (err) {
      this.handleLinkError(err)
    }
  }

  public handleLinkError(err: Error | ServiceError) {
    console.error(err)
  }

  public handleOnLinkSuccess(token: string, metadata: any) {
    console.log('Successfully authenticated user:')
    console.log(token)
    console.log(metadata)
    const req = new AddItemRequest()
    req.setToken(token)
    console.log(req)
    clients.users.addItem(
      req,
      getMetadata(),
      (
        error: ServiceError | null,
        resp: AddItemResponse | null
      ) => {
        if (error) {
          throw error
        }
        resp = resp as AddItemResponse
        console.log(resp.getUser())
      }
    )
  }

  public render() {
    return (
      <SettingsComponent
        handleOnLinkExit={this.handleOnLinkExit}
        handleOnLinkSuccess={this.handleOnLinkSuccess}
        plaidEnv={plaidEnv}
        plaidPublicKey={plaidPublicKey}
      />
    )
  }
}
