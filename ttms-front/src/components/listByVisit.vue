<template>
  <div class="main">
        <ul class="list">
          <div class="statement">
            <p style="color: #ffb400">2022-05-30已更新</p>
            <p style="margin-top:6px">榜单规则: 将昨日国内热映的影片, 按照评分从高到低排列取前10名, 每天上午10点更新。相关数据来源于“刘一琛统计”及“王宇飞统计”。</p>
          </div>
          <li 
          @click="toMovie(item.id)"
          class="item" v-for="(item,index) in movie" :key="item.id">
            <div class="aside">
              <img v-if="index==0" src="../assets/img/1.png" alt="" class="ranking">
              <div v-else-if="index!==0" class="ranking rankingimg">
                <p>{{index+1}}</p>
              </div>
              <img :src="item.avatar" alt="cover" class="cover">
            </div>
            <div class="describe">  
                <a>{{item.name}}</a>
                <p>主演：{{item.actors.slice(0,3).join(',')}}</p>
                <p style="margin-top:8px;  opacity:0.7 ;">上映时间: {{item.showTime.slice(0,10)}}</p>
            </div>
            <div class="score">
                  <i class="integer">{{parseInt(item.score)}}.</i>
                  <i class="fraction">{{parseInt((item.score-parseInt(item.score))*10)}}</i>
            </div>
          </li>
        </ul>
        <div class="page">
          <el-pagination
          background
          layout="prev, pager, next"
          @current-change="handlePage"
          :page-size="10"
          :total="50">
          </el-pagination>
      </div>
      </div>
</template>

<script setup>
import { onMounted, reactive } from "vue";
import {useRouter} from 'vue-router'
import { sortByRead } from "../apis/movie";
const router = useRouter()
const movie = reactive([])
const getMovieByRead = async () => {
  let res = await sortByRead()
  if(res.data.status_code == 0){
    console.log(res.data.data.movieInfo);
    movie.splice(0,movie.length,...res.data.data.movieInfo)
  }
}
const toMovie = (movieID) => {
  router.push({name:'thefilm',query:{movieID}})
}
onMounted(()=>{
  getMovieByRead()
})
//调接口返回列表
</script>

<style scoped>
.page {
  position: relative;
  top: 50px;
  left: 38%;
  flex-basis: 100%;
  text-align: center;
}
.main{
  width: 100%;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}
.describe a:hover{
      white-space: normal;

}
.statement{
  text-align: center;
  font-size: 12px;
  color: #999;
  margin:70px 0 40px 0;
}
.item{
  list-style: none;
  width:950px;
  height:220px;
  background-color:#fff;
  display:flex;
  align-items: center;
  border-bottom: 0.5px solid rgb(181, 176, 176);
  cursor: pointer;
}
.list{
  margin-bottom:100px;
}
.aside{
  display:flex;
  align-items: center;
  position: relative;
}
.ranking{
  display: inline-block;
    width: 50px;
    height: 50px;
    background: #f7f7f7;
    font-size: 18px;
    color: #999;
    font-weight: 700;
    position: absolute;
    left: 0;
    top: 85px;
    display:flex;
    align-items: center;
    justify-content: center;
}
.cover{
  width:160px;
  margin-left:80px;
  height:220px;
}
.describe a{
    font-size: 26px;
    color: #333;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    width: 420px;
    display: block;
    text-decoration: none;
    margin-bottom:18px;
}
.describe{
  margin:0 30px;
}
.score{
  margin-left:60px
}
.integer{
  color: #ffb400;
  font-size: 56px;
  font-weight: 700;
  font-style: italic;
}
.fraction{
  color: #ffb400;
  font-size: 26px;
    font-weight: 700;
  font-style: italic;
}

</style>