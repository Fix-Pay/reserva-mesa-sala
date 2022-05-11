import axios from 'axios'

export const http = axios.create({
    baseURL: 'http://localhost3000/api/'
})