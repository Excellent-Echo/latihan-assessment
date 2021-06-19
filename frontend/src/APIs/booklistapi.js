import axios from "axios";

const bookListAPI = axios.create({
  baseURL: "https://booklistappbymarwan.herokuapp.com",
});

export default bookListAPI;