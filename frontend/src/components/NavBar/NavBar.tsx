import * as React from 'react'

import { majorScale, minorScale, Pane, Text } from 'evergreen-ui'
import { RouteComponentProps, withRouter } from 'react-router-dom'

import withUser, { UserProviderProps } from '../withUser'
import AvatarMenu from './AvatarMenu'
import TabbedSubNav, { TabItem } from './TabbedSubNav/TabbedSubNav'

export declare interface NavBarProps extends RouteComponentProps<any>, UserProviderProps {}

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

    return (
      <Pane background="white" elevation={0}>
        <Pane display="flex" height={majorScale(7)}>
          <Pane
            flexBasis={majorScale(30)}
            display="flex"
            alignItems="center"
            paddingLeft={minorScale(3)}
          />
          <Pane
            className="fade-in-simple"
            flex={1}
            display="flex"
            justifyContent="center"
            alignItems="center"
          >
            <Text>~ moss ~</Text>
          </Pane>
          <Pane
            paddingRight={minorScale(3)}
            flexBasis={majorScale(30)}
            display="flex"
            alignItems="center"
            justifyContent="flex-end"
          >
            {this.props.user && <AvatarMenu user={this.props.user} />}
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

export default withUser(withRouter(NavBar))
