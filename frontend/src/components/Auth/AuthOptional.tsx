import * as React from 'react'
import { RouteComponentProps, withRouter } from 'react-router-dom'
import UserProvider from '../../context/User'

class AuthOptional extends React.Component<RouteComponentProps<any>> {
  public render() {
    return <UserProvider>{this.props.children}</UserProvider>
  }
}

export default withRouter(AuthOptional)