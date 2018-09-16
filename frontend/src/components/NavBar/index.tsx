import React from 'react';

import { majorScale, Pane, Tab } from 'evergreen-ui';
import {
  Link,
  RouteComponentProps,
  RouteProps,
  withRouter
} from 'react-router-dom';

interface TabItem {
  name: string;
  href: string;
  visible: boolean;
}

interface NavigationBarTabProps {
  to: string;
  label: string;
  location: RouteProps['location'];
}

interface TabbedSubNavigationProps {
  items: TabItem[];
  location: RouteProps['location'];
}

const NavigationBarTab: React.SFC<NavigationBarTabProps> = ({
  to,
  label,
  location
}) => {
  const isActive =
    location && (location.pathname === to || location.pathname.startsWith(to));

  return (
    <Tab is={Link} to={to} isSelected={isActive}>
      {label}
    </Tab>
  );
};

const TabbedSubNavigation: React.SFC<TabbedSubNavigationProps> = ({
  items,
  location
}) => {
  return (
    <Pane flexDirection="column">
      {items.filter((i: TabItem) => i.visible).map((item: TabItem) => (
        <NavigationBarTab
          key={item.href}
          label={item.name}
          to={item.href}
          location={location}
        />
      ))}
    </Pane>
  );
};

class NavBar extends React.Component<RouteComponentProps<any>> {
  public render() {
    const { location } = this.props;
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
    ];

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
            moss
          </Pane>
        </Pane>
        <Pane
          borderTop="muted"
          display="flex"
          justifyContent="center"
          paddingY={majorScale(1)}
        >
          <TabbedSubNavigation items={items} location={location} />
        </Pane>
      </Pane>
    );
  }
}

export default withRouter(NavBar);
