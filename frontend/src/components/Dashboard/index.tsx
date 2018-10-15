import * as React from 'react'

import { Heading, Pane } from 'evergreen-ui'
import withUser, { UserProps } from '../../hoc/withUser'

class Dashboard extends React.Component<UserProps> {
  public render() {
    return (
      <Pane>
        <Heading>Dashboard ({ this.props.user && this.props.user!.email })</Heading>
      </Pane>
    )
  }
}

export default withUser(Dashboard)