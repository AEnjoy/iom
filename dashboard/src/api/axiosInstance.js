import axios from 'axios'

axios.defaults.withCredentials = true;
const API = axios.create({
	baseURL:'http://localhost:8088', //
	timeout: 2000,                   //ms
	withCredentials:true
})

export default API