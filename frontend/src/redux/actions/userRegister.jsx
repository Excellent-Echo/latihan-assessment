import BookListAPI from '../../api/BookListAPI'

import { SET_EMAIL, SET_PASSWORD, SET_ADDRESS, SET_BIRTHDATE } from '../actionTypes/userRegister'

const setEmail = (email) => {
	return {
		type: SET_EMAIL,
		payload: {
			email: email
		}
	}
}

const setPassword = (password) => {
	return {
		type: SET_PASSWORD,
		payload: {
			password: password
		}
	}
}

const setAddress = (address) => {
	return {
		type: SET_ADDRESS,
		payload: {
			address: address
		}
	}
}

const setBirthDate = (birthDate) => {
	return {
		type: SET_BIRTHDATE,
		payload: {
			birthDate: birthDate
		}
	}
}

const register = (email, password, address, birthDate) => async (dispatch) => {
	try {
		const submitData = {
			email: email,
			password: password,
			address: address,
			birth_date: birthDate
		}

		const register = await BookListAPI({
			method: 'POST',
			url: '/user',
			data: submitData
		})
	} catch (error) {
		console.log(error)
	}
}

const userRegister = {
	setEmail,
	setPassword,
	setAddress,
	setBirthDate,
	register
}

export default userRegister
