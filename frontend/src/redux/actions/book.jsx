import BookListAPI from '../../api/BookListAPI'

import { SET_BOOKS } from '../actionTypes/book'

const setBooks = () => async (dispatch) => {
	try {
		const data = await BookListAPI({
			method: 'GET',
			url: '/books'
		})

		console.log(data)

		dispatch({
			type: SET_BOOKS,
			payload: {
				data: data.data
			}
		})
	} catch (error) {
		console.log(error)
	}
}

const book = {
	setBooks
}

export default book
