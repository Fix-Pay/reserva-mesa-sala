import axios from 'axios'

export const http = axios.create({
    baseURL: 'localhos:8000/'
})