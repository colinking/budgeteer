import * as React from 'react'

import Dashboard from '../Dashboard'
// import DefaultRoute from '../DefaultRoute'
import Login from '../Login'
import Logout from '../Logout'
import NavBar from '../NavBar'
import Settings from '../Settings'

import { Pane } from 'evergreen-ui'
import { Route, Switch } from 'react-router-dom'
import { AuthOptional, AuthRequired } from '../Auth'

export default class App extends React.Component<{}> {
  public render() {
    return (
      <AuthOptional>
        <Pane background="tint1" height="100vh">
          <NavBar />

          <Switch>
            {/* <DefaultRoute path="/dashboard"/> */}
            <Route path="/login" component={Login} />
            <AuthRequired>
              <Route path="/dashboard" component={Dashboard} />
              <Route path="/logout" component={Logout} />
              <Route path="/settings" component={Settings} />
            </AuthRequired>
          </Switch>
        </Pane>
      </AuthOptional>
    )
  }
}
