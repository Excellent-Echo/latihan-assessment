import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

import NavbarMenu from '../components/NavbarMenu'
import Footer from '../components/Footer'

import HomePage from '../pages/HomePage'
import Login from '../pages/Login'
import Register from '../pages/Register'
import NotFound from '../pages/NotFound'

const Routes = () => {
	return (
		<Router>
			<NavbarMenu />
			<Switch>
				<Router path="/register" exact component={Register} />
				<Router path="/login" exact component={Login} />
				<Route path="/" exact component={HomePage} />
				<Route component={NotFound} />
			</Switch>
			<Footer />
		</Router>
	)
}

export default Routes
