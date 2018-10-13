import React from 'react';

import { majorScale, Pane } from 'evergreen-ui';
import { RouteComponentProps, withRouter } from 'react-router-dom';

import TabbedSubNav, { TabItem } from '../TabbedSubNav/TabbedSubNav';

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
						moss (pre-alpha ya)
					</Pane>
				</Pane>
				<Pane borderTop="muted" display="flex" justifyContent="center" paddingY={majorScale(1)}>
					<TabbedSubNav items={items} location={location} />
				</Pane>
			</Pane>
		);
	}
}

export default withRouter(NavBar);
