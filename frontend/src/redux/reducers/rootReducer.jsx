import bookReducer from './book'
import userRegister from './userRegister'

const rootReducer = {
	books: bookReducer,
	register: userRegister
}

export default rootReducer
