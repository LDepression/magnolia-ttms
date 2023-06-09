
import { createRouter,createWebHistory } from "vue-router";
import Cookies from "js-cookie";
import Login from '../views/Login/index.vue';
import register from '../views/Login/register.vue'
import home from '../views/home/index.vue'
import layout from '../views/layout/index.vue'
import character from '../views/layout/authManger/character.vue'
import resource from '../views/layout/authManger/resource.vue'
import check from '../views/layout/cinemaManger/check.vue'
import play from '../views/layout/cinemaManger/play.vue'
import playSchedule from '../views/layout/cinemaManger/playSchedule.vue'
import studio from '../views/layout/cinemaManger/studio/index.vue'
import refund from '../views/layout/ticketManger/refund.vue'
import salesStatistics from '../views/layout/ticketManger/salesStatistics.vue'
import sellTicket from '../views/layout/ticketManger/sellTicket.vue'
import audio from '../views/LayOut/userManger/audio.vue'
import employee from '../views/LayOut/userManger/employee.vue'
import fare from '../views/LayOut/moneyManger/fare.vue'
import sales from '../views/LayOut/moneyManger/sales.vue'
import saleCollection from '../views/LayOut/moneyManger/saleCollection.vue'
import addCinema from '../views/LayOut/cinemaManger/studio/control/addCinema.vue'
import userCenter from '../views/userCenter/index.vue'
import user from '../views/LayOut/userManger/user.vue'
import myOrder from '../views/userCenter/child/myOrder.vue'
import userInfo from '../views/userCenter/child/userInfo.vue'
import addMovie from '../views/LayOut/cinemaManger/addMovie.vue'
import movieDetail from '../views/LayOut/cinemaManger/movieDetail.vue'
import start from '../components/start.vue'
import film from '../views/films/index.vue'
import list from '../views/list/index.vue'
import backHome from '../views/backHome/index.vue'
import chooseSeat from '../views/chooseSeat.vue'
import chooseTicket from '../views/chooseTicket.vue'
import comment from '../views/comment.vue'
import editCinema from '../views/LayOut/cinemaManger/studio/control/editCinema.vue'
import theFilm from '../views/theFilm.vue'
import order from '../views/order/index.vue'
import orderList from '../views/order/orderList/orderList.vue'
import orderBack from '../views/order/orderBack/orderBack.vue'
import search from '../views/search.vue'
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
            path:'/',
            redirect:'/home',
            name:'start',
            component:start,
           
            children:[
                {
                    path:'/home',
                    component:home,
                    name:'home',
                    meta:{title:'首页'},
                },
                {
                    path: "/film",
                    name: "film",
                    component: film,
                  meta: { title: "电影" }
  
                },
                {
                    path: "/list",
                    name: "list",
                    component: list,
                    meta: { title: "榜单" }
    
                  },
                  {
                    path: "/chooseSeat",
                    name: "chooseSeat",
                    component: chooseSeat
                  },
                  {
                    path:'/search',
                    name:'search',
                    component:search
                  },
                  {
                    path:'/theFilm',
                    name:'theFilm',
                    // redirect:'/comment',
                    component:theFilm,
                    children:[
                        // {
                        //   path: "/chooseticket",
                        //   name: "chooseticket",
                        //   component:chooseTicket,
                        // },
                        {
                          path: "/comment",
                          name: "comment",
                          component:Comment,
                        },
                      ]
                  },
                  {
                    path:'/chooseTicket',
                    name:'chooseTicket',
                    component: chooseTicket
                  },
                  {
                    path:'/comment',
                    component: comment
                  }
            ]
        },
        {
            path:'/backHome',
            component:backHome
        },
        {
            path:'/userCenter',
            component:userCenter,
            children:[
                {
                    path:'/userInfo',
                    component:userInfo
                },
                {
                    path:'/myOrder',
                    component:myOrder
                }
            ]
        },
        {
            path:'/layout',
            component:layout,
            redirect:'/studio',
            children:[
                {
                    path:'/character',
                    component:character,
                },
                {
                    path:'/editCinema/:id/:cinemaInfo',
                    component:editCinema,
                    props: true
                },
                {
                    path:'/movieDetail',
                    name:'movieDetail',
                    component:movieDetail,
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
                    path:'/user',
                    component:user,
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
                {
                    path:'/audio',
                    component:audio,
                },
                {
                    path:'/employee',
                    component:employee,
                },
                {
                    path:'/addMovie',
                    component:addMovie
                }
            ]
        }
    ]
})
// router.beforeEach((to,from,next) => {
//     //判断是否已经登录

//         if(!Cookies.get('AccessToken')){
//             next({
//                 path:'/login',
//                 query:{redirect:to.fullPath}// 将跳转的路由path作为参数，登录成功后跳转到该路由 
//             })
//         }
//         else{
//             next()
//         }
    
 
// })
export default router;
