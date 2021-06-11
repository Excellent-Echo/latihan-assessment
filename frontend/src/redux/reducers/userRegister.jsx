import { SET_EMAIL, SET_PASSWORD, SET_ADDRESS, SET_BIRTHDATE } from '../actionTypes/userRegister'

const initialState = {
	email: '',
	password: '',
	address: '',
	birthDate: ''
}

const userRegister = (state = initialState, action) => {
	switch (action.type) {
		case SET_EMAIL:
			return {
				...state,
				email: action.payload.email
			}
		case SET_PASSWORD:
			return {
				...state,
				password: action.payload.password
			}
		case SET_PASSWORD:
			return {
				...state,
				address: action.payload.address
			}
		case SET_BIRTHDATE:
			return {
				...state,
				birthDate: action.payload.birthDate
			}
		default:
			return state
	}
}

export default userRegister
