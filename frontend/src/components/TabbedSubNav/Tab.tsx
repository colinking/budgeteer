import * as React from 'react'

import { Tab } from 'evergreen-ui'
import { Link, RouteProps } from 'react-router-dom'

export interface NavigationBarTabProps {
  to: string
  label: string
  location: RouteProps['location']
}

const NavigationBarTab: React.SFC<NavigationBarTabProps> = ({
  to,
  label,
  location
}) => {
  const isActive =
    location && (location.pathname === to || location.pathname.startsWith(to))

  return (
    <Tab is={Link} to={to} isSelected={isActive}>
      {label}
    </Tab>
  )
}

export default NavigationBarTab
