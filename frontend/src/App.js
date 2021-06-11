import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import logo from './logo.svg';
import Register from './pages/Register';

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/register" exact component={Register} />
          <Route path="/" exact />
        </Switch>
      </Router>
    </div>
  );
}

export default App;
