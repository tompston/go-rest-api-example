import { createWebHistory, createRouter } from "vue-router";
import { isAuthenticated } from "./assets/ts/auth";


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'Home',
            component: () => import('./pages/Home.vue'),
            async beforeEnter(to, from, next) {
                /** 
                 * Block visit to homepage and redirect to
                 * login page if user is not authenticated
                 */
                if (await isAuthenticated()) next();
                else next("/auth");
            }
        },
        {
            path: '/auth',
            name: 'Authenticate',
            component: () => import('./pages/Authenticate.vue'),
        }
    ],
})

export default router