import React from 'react';

import { Pane } from 'evergreen-ui';
import { RouteProps } from 'react-router-dom';

import NavBarTab from './Tab';

export interface TabItem {
	name: string;
	href: string;
	visible: boolean;
}

export interface TabbedSubNavigationProps {
	items: TabItem[];
	location: RouteProps['location'];
}

const TabbedSubNav: React.SFC<TabbedSubNavigationProps> = ({ items, location }) => {
	return (
		<Pane flexDirection="column">
			{items
				.filter((i: TabItem) => i.visible)
				.map((item: TabItem) => (
					<NavBarTab key={item.href} label={item.name} to={item.href} location={location} />
				))}
		</Pane>
	);
};

export default TabbedSubNav;
