<<<<<<< HEAD
import { createRouter,createWebHistory } from "vue-router";
import Login from '../views/Login/index.vue';
import register from '../views/Login/register.vue'
import home from '../views/home/index.vue'
import layout from '../views/layout/index.vue'
const router = createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:'/Login',
            component:Login,
        },
        {
            path:'/register',
            component:register,
        },
        {
            path:'/home',
            component:home,
        },
        {
            path:'/layout',
            component:layout,
            children:[]
        }
    ]
})
export default router;
=======
import { createRouter,createWebHistory } from "vue-router";
import Login from '../views/Login/index.vue';
import register from '../views/Login/register.vue'
import home from '../views/home/index.vue'
import layout from '../views/layout/index.vue'
const router = createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:'/Login',
            component:Login,
        },
        {
            path:'/register',
            component:register,
        },
        {
            path:'/home',
            component:home,
        },
        {
            path:'/layout',
            component:layout,
            children:[]
        }
    ]
})
export default router;
>>>>>>> 73d85f9 (添加注册功能)
