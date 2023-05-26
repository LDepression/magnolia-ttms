import { createRouter,createWebHistory } from "vue-router";
import Login from '../views/Login/index.vue';
const router = createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:'/Login',
            component:Login,
        }
    ]
})
export default router;
