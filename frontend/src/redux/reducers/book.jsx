import { SET_BOOKS } from '../actionTypes/book'

const initialState = []

const book = (state = initialState, action) => {
	switch (action.type) {
		case SET_BOOKS:
			return action.payload.data

		default:
			return state
	}
}

export default book
