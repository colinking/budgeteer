import { Location } from 'history'
import * as React from 'react'
import { RouteComponentProps, withRouter } from 'react-router-dom'
import {
  getLoggedInUser,
  handleAuthenticationCallback,
  isAuthenticated,
  login
} from '../../lib/auth'
import { getHost } from '../../lib/host'
import { User, UserLoginRequest } from '../../proto/user/user_service_pb'
import { UserServiceClient } from '../../proto/user/user_service_pb_service'
import { REDIRECT_KEY } from '../Auth'
import Login from './Login'

class LoginApp extends React.Component<RouteComponentProps<any>> {
  public async componentDidMount() {
    const hash = this.props.location.hash
    if (/access_token|id_token|error/.test(hash)) {
      await handleAuthenticationCallback(hash)
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
    const u = await getLoggedInUser()

    const user = new User()
    user.setAuthId(u.auth_id)
    user.setFirstname(u.firstName)
    user.setLastname(u.lastName)
    user.setEmail(u.email)
    user.setPicture(u.picture)
    const req = new UserLoginRequest()
    req.setUser(user)

    new UserServiceClient(getHost()).userLogin(req, (err, res) => {
      if (err) {
        console.error(err)
        throw err
      }

      console.log(`Registered user. Are they new? ${res!.getNew()}`)
    })
  }
}

export default withRouter(LoginApp)
