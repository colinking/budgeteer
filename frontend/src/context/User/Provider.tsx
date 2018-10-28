import * as React from 'react'

import { User } from '../../lib/auth'
import { Provider } from './context'
import withFetchedUser from './withFetchedUser'

export declare interface ProviderProps {
  user?: User
}

class UserProvider extends React.Component<ProviderProps> {
  public render() {
    return <Provider value={this.props}>{this.props.children}</Provider>
  }
}

export default withFetchedUser(UserProvider)
