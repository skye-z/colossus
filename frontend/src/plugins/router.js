import {
    createRouter,
    createWebHistory
} from 'vue-router'

const Home = () => import('../views/home.vue')
const Host = () => import('../views/host.vue')
const Cert = () => import('../views/cert.vue')
const Code = () => import('../views/code.vue')
const Setting = () => import('../views/setting.vue')

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
            path: '/host/:id',
            component: Host
        },
        {
            name: 'Cert',
            path: '/cert',
            component: Cert
        },
        {
            name: 'Code',
            path: '/code',
            component: Code
        },
        {
            name: 'Setting',
            path: '/setting',
            component: Setting
        }
    ],
})

export default router