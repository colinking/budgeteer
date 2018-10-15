import * as React from 'react'

import { User } from '../../lib/auth'
import { Provider } from './context'
import withFetchedUser from './withFetchedUser'

export declare interface ProviderProps {
  user?: User
}

class UserProvider extends React.Component<ProviderProps> {
  public render() {
    console.log('provider rendering with value:')
    console.log(this.props)
    console.log(Provider)
    return <Provider value={this.props}>{this.props.children}</Provider>
  }
}

export default withFetchedUser(UserProvider)
