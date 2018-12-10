import React from 'react'
import { Consumer, UserProviderProps } from './UserContext'

export { UserProviderProps }

export default function withUser<P extends UserProviderProps>(
  Component: React.ComponentType<P>
) {
  return class WithUser extends React.Component {
    public static displayName = 'WithUser()'

    public render() {
      return (
        <Consumer>
          {({ user, refetchUser }: UserProviderProps) => {
            return <Component {...this.props} user={user} refetchUser={refetchUser} />
          }}
        </Consumer>
      )
    }
  }
}
