import * as React from 'react'
import UserProvider from '../UserContext'

class AuthOptional<P> extends React.Component<P> {
  public render() {
    return <UserProvider>{this.props.children}</UserProvider>
  }
}

export default AuthOptional
