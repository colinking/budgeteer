import * as React from 'react'
import { Redirect, Route } from 'react-router-dom'

declare interface DefaultRouteProps {
  path: string
}

// FIXME: Seems like this doesn't work. ¯\_(ツ)_/¯
const DefaultRoute = ({ path }: DefaultRouteProps) => {
  return (
    <Route exact path="/">
      <Redirect to={path} />
    </Route>
  )
}

export default DefaultRoute
