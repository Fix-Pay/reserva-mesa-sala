import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import {routes} from './routes'
import { store } from './index'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
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
