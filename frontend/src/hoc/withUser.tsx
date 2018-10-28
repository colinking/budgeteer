import React from 'react'
import { Consumer } from '../context/User'
import { ProviderProps } from '../context/User/Provider'

export { ProviderProps as UserProps }

export default function withUser<P extends ProviderProps>(
  Component: React.ComponentType<P>
) {
  return class WithUser extends React.Component {
    public static displayName = 'WithUser()'

    public render() {
      return (
        <Consumer>
          {({ user }: ProviderProps) => {
            return <Component {...this.props} user={user} />
          }}
        </Consumer>
      )
    }
  }
}
