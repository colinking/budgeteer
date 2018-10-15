import { RouteComponentProps } from 'react-router-dom'
import { isAuthenticated } from '../../lib/auth'
import AuthOptional from './AuthOptional'

export const REDIRECT_KEY = 'auth_redirect'

class AuthRequired extends AuthOptional {
  public constructor(props: RouteComponentProps<any>) {
    super(props)

    if (!isAuthenticated()) {
      localStorage.setItem(REDIRECT_KEY, JSON.stringify(props.location))
      props.history.push('/login')
    }
  }
}

export default AuthRequired
