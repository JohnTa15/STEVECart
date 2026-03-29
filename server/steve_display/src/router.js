import { createRouter, createWebHistory } from 'vue-router'
import Login from './pages/Login.vue'
import Signup from './pages/Signup.vue'
import Index from './pages/Index.vue'

const routes = [
    { path: '/login', component: Login },
    { path: '/signup', component: Signup },
    { path: '/index', component: Index },
    { path: '/', redirect: '/login' },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router