import { createRouter, createWebHistory } from 'vue-router'
import Login from './pages/Login.vue'
import Signup from './pages/Signup.vue'
import Index from './pages/Index.vue'
import Management from './pages/Management.vue'
import { transformVNodeArgs } from 'vue'

const routes = [
    { path: '/login', component: Login },
    { path: '/signup', component: Signup },
    { path: '/index', component: Index },
    { path: '/management', component: Management },
    { path: '/', redirect: '/login' },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

//blocked manual typing pages like /management or /index when is not logged in.. 
router.beforeEach((to, from, next) => {
    const publicPGs = ['/login', '/signup']
    const isloggedIn = !!localStorage.getItem('username')

    if (!publicPGs.includes(to.path) && !isloggedIn) {
        return next('/login')
    }
    next()
})
export default router