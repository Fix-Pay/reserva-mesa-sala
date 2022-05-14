import Vue from "vue";
import Vuex from 'vuex';

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {fotos: [
        {
        id: '1',
        titulo: 'mesa 1',
        status: false,
        },
        {
        id: '2',
        titulo: 'mesa 2',
        status: false,
        },
         {
        id: '3',
        titulo: 'mesa 3',
        status: false,
        },
      ] },
    getters: {},
    mutation: {},
    actions: {},
});

export {store};
