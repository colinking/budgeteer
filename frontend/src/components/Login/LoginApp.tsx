import { Location } from 'history'
import * as React from 'react'
import { RouteComponentProps, withRouter } from 'react-router-dom'
import {
  getLoggedInUser,
  handleAuthenticationCallback,
  isAuthenticated,
  login
} from '../../lib/auth'
import { REDIRECT_KEY } from '../Auth'
import Login from './Login'

class LoginApp extends React.Component<RouteComponentProps<any>> {
  public async componentDidMount() {
    const hash = this.props.location.hash
    if (/access_token|id_token|error/.test(hash)) {
      await handleAuthenticationCallback(hash)
      await this.registerNewUser()
    }

    if (isAuthenticated()) {
      await this.registerNewUser()
      const redirect = localStorage.getItem(REDIRECT_KEY)
      localStorage.removeItem(REDIRECT_KEY)

      if (redirect) {
        const location: Location = JSON.parse(redirect)
        this.props.history.replace(location)
      } else {
        this.props.history.replace('/dashboard')
      }
    } else {
      login()
    }
  }

  public render() {
    return <Login />
  }

  protected async registerNewUser() {
    const user = await getLoggedInUser()
    console.log(user)

    // TODO: upload to back-end to register the user
  }
}

export default withRouter(LoginApp)
