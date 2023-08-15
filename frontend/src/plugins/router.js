import {
    createRouter,
    createWebHistory
} from 'vue-router'

const Home = () => import('../views/home.vue')

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            name: 'Home',
            path: '/home',
            component: Home
        }
    ],
})

export default router