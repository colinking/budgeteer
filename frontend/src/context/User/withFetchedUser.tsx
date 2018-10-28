import * as React from 'react'
import { getLoggedInUser, isAuthenticated, User } from '../../lib/auth'

export declare interface WithFetchedUserState {
  user?: User
}

export default function withFetchedUser<P extends WithFetchedUserState>(
  Component: React.ComponentType<P>
) {
  return class WithFetchedUser extends React.Component<
    any,
    WithFetchedUserState
  > {
    public state = {
      user: undefined
    }

    public async componentDidMount() {
      // Manually fired in `lib/auth.tsx`
      window.addEventListener('storage', this.handleLogin)
      this.checkForLogin()
    }

    public async componentWillUnmount() {
      // Manually fired in `lib/auth.tsx`
      window.removeEventListener('storage', this.handleLogin)
    }

    public render() {
      return <Component {...this.props} user={this.state.user} />
    }

    protected handleLogin = async () => {
      this.checkForLogin()
    }

    protected async checkForLogin() {
      if (isAuthenticated()) {
        const user = await getLoggedInUser()
        console.log('checkForLogin: auth-ed', user)
        this.setState({ user })
      } else {
        console.log('checkForLogin: NOT auth-ed')
        this.setState({ user: undefined })
      }
    }
  }
}
