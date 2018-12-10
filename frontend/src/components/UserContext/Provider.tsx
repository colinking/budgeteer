import * as React from 'react'

import { Provider } from './context'
import withFetchedUser, { WithFetchedUserProps } from './withFetchedUser'

export declare interface UserProviderProps extends WithFetchedUserProps {}

class UserProvider extends React.Component<UserProviderProps> {
  public render() {
    return <Provider value={this.props}>{this.props.children}</Provider>
  }
}

export default withFetchedUser(UserProvider)
