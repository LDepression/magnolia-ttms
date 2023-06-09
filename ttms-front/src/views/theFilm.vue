<template>
    <div>
        <div class="back">
            <div class="main">
                <div class="cover">
                    <img :src="moviedetail.avatar" alt="封面">
                </div>
                <div class="content">
                    <div class="info">
                        <h1 class="name">{{moviedetail.name}}</h1>
                      
                        <ul class="label">
                         
                            <li>
                              评分：
                              <el-rate
                                  style="display:inline"
                                  :value="(moviedetail.score/2)"
                                  disabled
                                  text-color="#ff9900">
                              </el-rate>
                              <span style="color:#ffc600">{{moviedetail.score}}</span>
                            </li>
                            <li> {{moviedetail.duration}}分钟</li>
                            <li> {{moviedetail.showTime}}中国大陆上映</li>
                        </ul>
                    </div>
                    <div style="margin-top:30px">
                        <el-button type="warning" class="box" icon="el-icon-video-play" >{{moviedetail.IsFollow?'取消':'想看'}}</el-button>
                      <el-button type="warning"
                      
                       icon="el-icon-star-off" class="box" >评分</el-button>
                      <br>
                      <el-button type="warning" 
                      @click="route.push('/chooseTicket')"
                      style="width: 188px; margin-top:10px" >特惠购票</el-button>
                    </div>
                    <div class="evaluate">
                        <div class="movieindex">
                           
                              
                        </div>
                    </div>
                </div>
            </div>
              <div class="plot">
              <h1 class="labelname">剧情简介</h1>
              <hr/>
              <!-- <p class="synopsis"><span style="color: red;font-weight: bold;">主要演员：</span> {{moviedetail.actors.join(' ')}}<br/></p> -->
              <p class="synopsis">{{moviedetail.content}}</p>
              </div>
        </div>
      <!-- <rouetr-view></rouetr-view> -->
    </div>
  </template>
  
  <script setup>
  import comment from '../views/comment.vue'
  import chooseTicket from '../views/chooseTicket.vue'
  import { useRoute,useRouter } from 'vue-router'
  import Cookies from 'js-cookie'
  import { onMounted, reactive,ref } from 'vue'
  import {getMovieDetail} from '../apis/movie'
  const router = useRoute()
  const moviedetail = reactive({})
  const movie_id = ref()
  const route = useRoute()
  movie_id.value = router.query.movie_id
  console.log(movie_id.value);
  const getInfo = async (movie_id) => {
      console.log(1);
      let res = await getMovieDetail(movie_id);
     console.log(res.data);
      if(res.data.status_code == 0){
          Object.assign(moviedetail,res.data.data.movieInfo)
      }
      console.log(moviedetail);
  }
  const giveLike = () => {
  //   let token = Cookies.get('AccessToken')
  //   if(router.path=='/chooseTicket'){
  //     return ;
  //   }
  //   if(!token){
  //     ElMessage({
  //       message: "请前往登录",
  //       center: true,
  //       type: "error",
  //     });
  //   }
  
  }
  onMounted(() => { getInfo(Number(movie_id.value)) })
  </script>
  
  <style scoped>
  .synopsis{
      margin-left:20px;
      text-indent: 2em;
  }
  .plot{
      width:1000px;
      margin:50px auto;
      padding:10px;
  }
  .labelname::before{
      content: "";
      display: inline-block;
      width: 8px;
      height: 28px;
      margin-right: 6px;
      background-color: #11a8cd;
  }
  .title{
      font-size: 12px;
      margin-bottom: 8px;
  }
  .number{
      font-size: 30px;
      color: #ffc600;
      height: 30px;
      line-height: 30px;
      margin-bottom:16px;
  }
  .evaluate{
      position: absolute;
      width: 116.7px;
      height: 136.8px;
      top:158px;
      left:342px;
      z-index: 100;
  }
.box{
    text-indent: -15px;
}
  .label {
      width: 250px;
      list-style: none;
      padding-left: 0;
      margin-bottom: 20px;
  }
  .label li{
      margin: 12px 0;
      line-height: 100%;
  }
  .name{  
      white-space:nowrap;
      width: 400px;
      margin-top: 0;
      font-size: 26px;
      line-height: 32px;
      font-weight: 700;
      margin-bottom: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      max-height: 64px;
  }
  .ename:hover{
      width: 700px;
  }
  .name:hover{
      width: 700px;
  }
  .ename{
      text-overflow: ellipsis;
      overflow: hidden;
      white-space: nowrap;
      width: 340px;
      font-size: 18px;
      line-height: 1.3;
      margin-bottom: 14px;
  }
  .content {
      color:#fff;
      position: relative;
      margin:70px 30px 0 300px;
      width: 870px;
      height:300px;
  }
  .back{
      height: 876px;
      width: 100vw;
      position: relative;
      
      display: inline-block;
      background: #392f59 url(../assets/img/theFilmBack.png) no-repeat 50%;
  }
  .cover{
      float: left;
      overflow: hidden;
      z-index: 9;
      width: 240px;
      height: 330px;
      margin: 0 30px;
  }
  .cover img{
      border: 4px solid #fff;
      height: 322px;
      width: 232px;
  }
  .main{
      width:1200px;
      margin: 0 auto;
  }
  </style>