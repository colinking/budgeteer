import { Location } from 'history'
import * as React from 'react'
import { RouteComponentProps, withRouter } from 'react-router-dom'
import {
  getAuth0UserData,
  handleAuthenticationCallback,
  isAuthenticated,
  login
} from '../../lib/auth'
import { users } from '../../clients'
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
      await this.registerUser()
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

  protected registerUser = async () => {
    const u = await getAuth0UserData()

    const resp = await users.login(req => {
      req.setName(u.name)
      if (u.email) {
        req.setEmail(u.email)
      }
      req.setPictureurl(u.picture)
    })

    console.log(`Registered user. Are they new? ${resp.getNew()}`)
  }
}

export default withRouter(LoginApp)
