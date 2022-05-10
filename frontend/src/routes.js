import Home from './components/home/Home.vue'
import Logs from './components/logs/Logs.vue'
import Login from './components/login/Login.vue'
import Register from './components/register/Register.vue'

export const routes = [
    { path: '', component: Home },
    { path: '/logs', component: Logs},
    { path: '/login', component: Login},
    { path: '/register', component: Register}
];
