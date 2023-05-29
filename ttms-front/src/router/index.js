import { createRouter,createWebHistory } from "vue-router";
import Login from '../views/Login/index.vue';
import register from '../views/Login/register.vue'
import home from '../views/home/index.vue'
import layout from '../views/layout/index.vue'
import character from '../views/layout/authManger/character.vue'
import resource from '../views/layout/authManger/resource.vue'
import check from '../views/layout/cinemaManger/check'
import play from '../views/layout/cinemaManger/play'
import playSchedule from '../views/layout/cinemaManger/playSchedule'
import studio from '../views/layout/cinemaManger/studio'
import refund from '../views/layout/ticketManger/refund'
import salesStatistics from '../views/layout/ticketManger/salesStatistics'
import sellTicket from '../views/layout/ticketManger/sellTicket'
import audio from '../views/LayOut/userManger/audio'
import employee from '../views/LayOut/userManger/employee'
import fare from '../views/LayOut/moneyManger/fare.vue'
import sales from '../views/LayOut/moneyManger/sales.vue'
import saleCollection from '../views/LayOut/moneyManger/saleCollection'
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
            ]
        }
    ]
})
export default router;
