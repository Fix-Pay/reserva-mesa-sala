import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import {routes} from './routes'
import { store } from './index'

Vue.use(VueRouter)

const router = new VueRouter({
  routes: routes,
  mode:'history'
})

export default store;

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
