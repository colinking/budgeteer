import * as React from 'react'

import { Heading, Pane } from 'evergreen-ui'
import withUser, { UserProviderProps } from '../withUser'

class Dashboard extends React.Component<UserProviderProps> {
  public render() {
    return (
      <Pane>
        <Heading>Dashboard ({ this.props.user && this.props.user!.email })</Heading>
      </Pane>
    )
  }
}

export default withUser(Dashboard)
