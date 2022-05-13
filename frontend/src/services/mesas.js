import { http } from './config'

export default {

    listar:() => {
        return http.get('tables')
    },
    salvar:(id) => {
        return http.post('locais', id,horario1,horario2)
    }
}