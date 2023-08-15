import {
    createRouter,
    createWebHistory
} from 'vue-router'

const Home = () => import('../views/home.vue')
const Host = () => import('../views/host.vue')

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            name: 'Home',
            path: '/home',
            component: Home
        },
        {
            name: 'Host',
            path: '/host',
            component: Host
        }
    ],
})

export default router