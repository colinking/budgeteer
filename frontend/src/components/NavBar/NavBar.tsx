import * as React from 'react'

import { majorScale, Pane } from 'evergreen-ui'
import { RouteComponentProps, withRouter } from 'react-router-dom'

import withUser from '../../hoc/withUser'
import { User } from '../../lib/auth'
import TabbedSubNav, { TabItem } from '../TabbedSubNav/TabbedSubNav'

export declare interface NavBarProps extends RouteComponentProps<any> {
  user?: User
}

class NavBar extends React.Component<NavBarProps> {
  public render() {
    const { location } = this.props
    const items: TabItem[] = [
      {
        href: '/dashboard',
        name: 'Dashboard',
        visible: true
      },
      {
        href: '/settings',
        name: 'Settings',
        visible: true
      }
    ]

    if (this.props.user) {
      items.push({
        href: '/logout',
        name: `Logout (${this.props.user.firstName})`,
        visible: true
      })
    }

    return (
      <Pane background="white" elevation={0}>
        <Pane display="flex" height={majorScale(7)}>
          <Pane
            className="fade-in-simple"
            flex={1}
            display="flex"
            justifyContent="center"
            alignItems="center"
          >
            moss (pre-alpha ya)
          </Pane>
        </Pane>
        <Pane
          borderTop="muted"
          display="flex"
          justifyContent="center"
          paddingY={majorScale(1)}
        >
          <TabbedSubNav items={items} location={location} />
        </Pane>
      </Pane>
    )
  }
}

// TODO: why does context not work here?
export default withUser(withRouter(NavBar))
