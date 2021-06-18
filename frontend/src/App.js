import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import logo from './logo.svg';
import NotFound from './pages/NotFound';
import Register from './pages/Register';
import styled from "styled-components";

function App() {
  const App = styled.div`
  position: relative;
  min-height: 100vh;
  `;

  return (
    <App className="App">
      <Router>
        <Switch>
          <Route path="/register" exact component={Register} />
          <Route path="/" exact />
          <Route component={NotFound} />
        </Switch>
      </Router>
    </App>
  );
}

export default App;
