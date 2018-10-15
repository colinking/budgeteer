import * as React from 'react'
import { RouteComponentProps, withRouter } from 'react-router-dom'
import { logout } from '../../lib/auth'
import Logout from './Logout'

class LogoutApp extends React.Component<RouteComponentProps<any>> {
  public async componentDidMount() {
    logout()
    this.props.history.replace('/login')
  }

  public render() {
    return <Logout />
  }
}

export default withRouter(LogoutApp)
