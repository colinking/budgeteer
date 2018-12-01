import { Location } from 'history'
import * as React from 'react'
import { RouteComponentProps, withRouter } from 'react-router-dom'
import { LoginRequest, User } from '../../gen/userpb/user_service_pb'
import {
  getLoggedInUser,
  handleAuthenticationCallback,
  isAuthenticated,
  login
} from '../../lib/auth'
import clients from '../../lib/clients'
import { getMetadata } from '../../lib/requests'
import { REDIRECT_KEY } from '../Auth'
import Error, { WithError } from '../Error'
import Login from './Login'

interface LoginState {
  error?: WithError
}

class LoginApp extends React.Component<RouteComponentProps<any>, LoginState> {
  public state = {
    error: undefined
  }

  public async componentDidMount() {
    const hash = this.props.location.hash
    if (/access_token|id_token|error/.test(hash)) {
      try {
        await handleAuthenticationCallback(hash).catch(err => {
          throw err
        })
      } catch (err) {
        this.setState({
          error: {
            error: err.error,
            description: err.errorDescription
          }
        })
        return
      }
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
    return this.state.error ? <Error {...this.state.error!} /> : <Login />
  }

  protected async registerNewUser() {
    const u = await getLoggedInUser()

    const user = new User()
    user.setFirstname(u.firstName)
    user.setLastname(u.lastName)
    user.setEmail(u.email)
    user.setPictureurl(u.picture)
    const req = new LoginRequest()
    req.setUser(user)

    clients.users.login(req, getMetadata(), (err, res) => {
      if (err) {
        console.error(err)
        throw err
      }

      console.log(`Registered user. Are they new? ${res!.getNew()}`)
    })
  }
}

export default withRouter(LoginApp)
