import * as React from 'react'
import { getLoggedInUser, isAuthenticated } from '../../lib/auth'
import { User } from '../../clients'

export declare interface WithFetchedUserState {
  user?: User
}

export declare interface WithFetchedUserProps {
  user?: User
  refetchUser: () => void
}

export default function withFetchedUser<P extends WithFetchedUserProps>(
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
      return <Component {...this.props} user={this.state.user} refetchUser={this.checkForLogin} />
    }

    protected handleLogin = async () => {
      this.checkForLogin()
    }

    protected checkForLogin = async () => {
      if (isAuthenticated()) {
        const user = await getLoggedInUser()
        this.setState({ user })
      } else {
        this.setState({ user: undefined })
      }
    }
  }
}
