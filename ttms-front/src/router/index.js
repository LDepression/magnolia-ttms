<<<<<<< HEAD
import { createRouter,createWebHistory } from "vue-router";
import Login from '../views/Login/index.vue';
import register from '../views/Login/register.vue'
import home from '../views/home/index.vue'
import layout from '../views/layout/index.vue'
import character from '../views/layout/authManger/character.vue'
import resource from '../views/layout/authManger/resource.vue'
import check from '../views/layout/cinemaManger/check.vue'
import play from '../views/layout/cinemaManger/play.vue'
import playSchedule from '../views/layout/cinemaManger/playSchedule.vue'
import studio from '../views/layout/cinemaManger/studio.vue'
import refund from '../views/layout/ticketManger/refund.vue'
import salesStatistics from '../views/layout/ticketManger/salesStatistics.vue'
import sellTicket from '../views/layout/ticketManger/sellTicket.vue'
import audio from '../views/LayOut/userManger/audio.vue'
import employee from '../views/LayOut/userManger/employee.vue'
import fare from '../views/LayOut/moneyManger/fare.vue'
import sales from '../views/LayOut/moneyManger/sales.vue'
import saleCollection from '../views/LayOut/moneyManger/saleCollection.vue'
import addCinema from '../views/LayOut/cinemaManger/addCinema.vue'
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
            children:[
                {
                    path:'/character',
                    component:character,
                },
                {
                    path:'/resource',
                    component:resource,
                },
                {
                    path:'/check',
                    component:check,
                },
                {
                    path:'/play',
                    component:play,
                },
                {
                    path:'/playSchedule',
                    component:playSchedule,
                },
                {
                    path:'/studio',
                    component:studio,
                    
                        
                },
                {
                    path:'/addCinema',
                    component:addCinema,
                }
            
                ,

                {
                    path:'/fare',
                    component:fare,
                },
                {
                    path:'/saleCollection',
                    component:saleCollection,
                },
                {
                    path:'/sales',
                    component:sales,
                },
                {
                    path:'/refund',
                    component:refund,
                },
                {
                    path:'/salesStatistics',
                    component:salesStatistics,
                },
                {
                    path:'/sellTicket',
                    component:sellTicket,
                },
                {
                    path:'/audio',
                    component:audio,
                },
                {
                    path:'/employee',
                    component:employee,
                },
            ]
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
