import Vue from 'vue';
import Router from 'vue-router';
import { isSignedIn } from './auth';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/home',
      component: () => import('components/home/Home.vue'),
      beforeEnter (_, __, next) { // Impede usuários não assinados
        if (isSignedIn()) {       // de acessar a página Home.
          next();
          return;
        }
        next('/sign-in')
      }
    },
    {
      path: '/sign-in',
      component: () => import('components/login/Login.vue'),
      beforeEnter (_, __, next) { // Impede usuários assinados de
        if (!isSignedIn()) {      // acessar a página de login.
          next();
          return;
        }
        next('/home')
      }
    }
  ]
})
