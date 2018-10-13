import React from 'react';

import Dashboard from '../Dashboard';
import NavBar from '../NavBar';
import Settings from '../Settings';

import { Pane } from 'evergreen-ui';
import { Route, Switch } from 'react-router-dom';

export default class App extends React.Component<{}> {
  public render() {
    return (
      <Pane background="tint1" height="100vh">
        <NavBar />

        <Switch>
          <Route path="/dashboard" component={Dashboard} />
          <Route path="/settings" component={Settings} />
        </Switch>
      </Pane>
    );
  }
}
