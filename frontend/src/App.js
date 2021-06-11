import React from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import HomePage from './pages/Homepage';
import LoginPage from './pages/LoginPage';

function App() {
  return (
      <>
        <Router>
          <Switch>
            <Route path="/register" exact component={RegisterPage} />
            <Route path="/login" exact component={LoginPage} />
            <Route path="/" exact component={HomePage} />
          </Switch>
        </Router>
      </>
  );
}

export default App;
