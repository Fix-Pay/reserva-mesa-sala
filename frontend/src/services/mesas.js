import { http } from './config'

export default {

    listar:() => {
        return http.get('locais')
    },
    salvar:(id) => {
        return http.post('locais', id)
    }
}