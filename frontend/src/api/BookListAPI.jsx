import axios from 'axios'

const BookListAPI = axios.create({
	baseURL: 'http://localhost:4444'
})

export default BookListAPI
