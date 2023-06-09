<template>
  <swiper class="swiper" />
  <div class="contain">
    <div class="content">
      <div class="main">
        <div class="movie-grid">
          <div class="panel">
            <div class="panel-header">
              <span class="more">
                <a href="#"
                  ><span class="all" @click="allfilm">全部</span
                  ><svg
                    t="1654044671881"
                    class="icon"
                    viewBox="0 0 1024 1024"
                    version="1.1"
                    xmlns="http://www.w3.org/2000/svg"
                    p-id="1237"
                    width="20"
                    height="20"
                  >
                    <path
                      d="M612.266667 514.133333c0-2.133333 0-2.133333 0 0L369.066667 755.2c-8.533333 8.533333-8.533333 21.333333 0 29.866667 8.533333 8.533333 21.333333 8.533333 29.866666 0l243.2-243.2c17.066667-17.066667 17.066667-42.666667 0-57.6L401.066667 241.066667c-8.533333-8.533333-21.333333-8.533333-29.866667 0-8.533333 8.533333-8.533333 21.333333 0 29.866666l241.066667 243.2z"
                      fill="#ef4238"
                      p-id="1238"
                    ></path></svg
                ></a>
              </span>
              <span class="title">
                <h2>正在热映 (50部)</h2>
              </span>
            </div>

            <div class="panel-content">
              <div class="movie-list">
                <dd
                  @click="toMovie(item.movieID)"
                  v-for="(item, index) in movies"
                  :key="item.movieID"
                >
                  <div class="movie-item" v-if="index < 6">
                    <img :src="item.avatar" alt="" />
                    <div class="hot-movie-score">
                      <i>{{ item.score.toFixed(1) }}</i>
                    </div>
                    <div>
                      <p class="hot-movie-name">{{ item.name }}</p>
                    </div>
                    <div class="buy" @click="buyit">购票</div>
                  </div>
                </dd>
              </div>
            </div>
          </div>
          <div class="hot-panel">
            <div class="hot-panel-header">
              <span class="hot-panel-header-more">
                <a @click="allfilm">
                  <span
                    >全部<svg
                      t="1654044671881"
                      class="icon"
                      viewBox="0 0 1024 1024"
                      version="1.1"
                      xmlns="http://www.w3.org/2000/svg"
                      p-id="1237"
                      width="20"
                      height="20"
                    >
                      <path
                        d="M612.266667 514.133333c0-2.133333 0-2.133333 0 0L369.066667 755.2c-8.533333 8.533333-8.533333 21.333333 0 29.866667 8.533333 8.533333 21.333333 8.533333 29.866666 0l243.2-243.2c17.066667-17.066667 17.066667-42.666667 0-57.6L401.066667 241.066667c-8.533333-8.533333-21.333333-8.533333-29.866667 0-8.533333 8.533333-8.533333 21.333333 0 29.866666l241.066667 243.2z"
                        fill="#ef4238"
                        p-id="1238"
                      ></path></svg
                  ></span>
                </a>
              </span>
              <span class="hot-panel-header-title"><h2>热播电影</h2></span>
              <span class="panel-subtitle">
                <a style="color: aliceblue" @click="allfilm">爱情</a>
              </span>
              <span class="panel-subtitle">
                <a style="color: aliceblue" @click="allfilm">喜剧</a>
              </span>
              <span class="panel-subtitle">
                <a style="color: aliceblue" @click="allfilm">动作</a>
              </span>
              <span class="panel-subtitle">
                <a style="color: aliceblue" @click="allfilm">恐怖</a>
              </span>
              <span class="panel-subtitle">
                <a style="color: aliceblue" @click="allfilm">动画</a>
              </span>
            </div>
            <div class="hot-panel-content">
              <dl class="hot-movie-list">
                <!-- --------------------------------热播电影------------------------- -->
                <dd
                  class="dd"
                  @click="toMovie(item.movieID)"
                  v-for="(item, index) in movies2"
                  :key="item.id"
                >
                  <div v-if="index < 8" class="hot-movie-item2">
                    <a href="#">
                      <div class="movie-poster-short">
                        <div class="fl">
                          <img :src="item.avatar" alt="" />
                        </div>
                        <div class="movie-info2">
                          <div class="movie-score2">
                            <i>{{ item.score.toFixed(1) }}</i>
                          </div>
                          <div class="movie-title-padding2">
                            {{ item.name }}
                          </div>
                          <div class="buy">购票</div>
                        </div>
                      </div>
                    </a>
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>
      <div class="aside">
        <div class="ranking-box-wrapper">
          <div class="panel2">
            <div class="panel2-header">
              <div class="panel2-title">
                <h2>今日票房</h2>
              </div>
            </div>
            <div class="panel2 content">
              <ul class="ranking-box">
               
                <li
                  class="ranking-index2"
                  
                  v-for="(item,index) in movies6"
                  :key="item.id"
                  @click="toMovie(item.movieID)"
                >
                  <i class="i">{{ index + 1 }}</i>
                  <a href="#">
                    <span class="ranking-movie-name">
                      {{ item.name }}
                    </span>
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>
        <div class="box-total-wrapper">
          <div class="ticket">
            <p>
              <span class="font">1001.89</span>
              <span class="num">万</span>
            </p>
            <p class="meta-info">magnolia票房数据</p>
          </div>
        </div>
    
      </div>
    </div>
  </div>
  <Footer />
</template>

<script setup>
import swiper from "../../components/swiper.vue";
import { getAllMovie, sortByRead, sortByExpect } from "../../apis/movie";
import { onMounted, reactive, ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter()
let movies = reactive([]);
const movies2 = reactive([]);
const page = ref(1);
const movies3 = reactive([]);
const movies4 = reactive([]);
const movies5 = reactive([]);
const movies6 = reactive([]);
const movies7 = reactive([]);
const allfilm = () => {
  router.push({ path: "/film" });
};
const alllist = () => {
  router.push({ path: "/list" });
};
const toMovie = (movie_id) => {
  router.push({ path: "/theFilm", query:{ movie_id } });
};
//获取所有电影
const getAllMovies = async () => {
  let res = await getAllMovie();
  // movies = res.data.data.movieInfo
  movies.splice(0, movies.length, ...res.data.data.movieInfo);
  console.log(movies);
};
//获取热度电影
const getHotMovies = async () => {
  let res = await sortByRead();
  console.log(res);
  if (res.data.status_code == 0) {
    movies2.splice(0, movies2.length, ...res.data.data.movieInfo);
  }
};
//获取最受期待
const getExpectMovie = async () => {
  let res = await sortByExpect();
  if (res.data.status_code == 0) {
    movies6.splice(0, movies6.length, ...res.data.data.movieInfo);
  }
};
onMounted(() => {
  getAllMovies(), getHotMovies(), getExpectMovie();
});
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
}
.content {
  display: flex;
  width: 100%;
  margin: 30px auto;
}
.aside {
  flex: 1;
  width: 400px;
  height: 1500px;
  position: relative;
  left: 50px;
  /* background-color: pink; */
}
.main {
  width: 50%;
}
.all {
  position: relative;
  left: -150px;
}
.panel-header {
  margin-top: 5px;
  width: 760px;
  height: 26px;
  line-height: 26px;
}
.panel-content {
  margin-top: 35px;
  width: 700px;
  height: 569px;
  margin-left: 45px;
  /* background-color: red; */
}
.dd {
  position: relative;
  top: 40px;
}
.title h2 {
  /* float: left; */
  position: relative;
  left: -200px;
  text-align: center;
  color: #ef4238 !important;
  font-weight: 320;
}
.panel2-header {
  width: 320px;
}
.panel2-title h2 {
  color: #ef4238 !important;
  font-weight: 320;
  margin-left: -200px;
  padding: 3px 0 0 3px;
}
.panel2-content {
  height: 311px;
  width: 320px;
}
.ranking-box li {
  list-style: none;
}
.more a {
  float: right;
  font-size: 14px;
  text-decoration: none;
  color: #ef4238 !important;
}
.i {
  position: relative;
  left: 20px;
}
.more a:hover {
  text-decoration: underline;
}
.more .icon {
  /* float: right;
    margin: 4px 2px 0 0; */
  position: relative;
  left: -150px;
  top: 5px;
}
.img-link img {
  width: 170px;
  height: 118px;
  border: 1px solid #efefef;
}
.movie-list {
  margin: -29px 0 20px -25px;
}
.movie-item {
  float: left;
  cursor: pointer;
  position: relative;
  border: 1px solid #efefef;
  margin: 10px 0 0 30px;
}
.movie-item img {
  height: 220px;
  width: 160px;
}
.fl {
  float: left;
  cursor: pointer;
  position: relative;
}
.fl img {
  height: 224px;
  width: 160px;
  border: 1px solid #efefef;
}
.movie-item .hot-movie-score {
  position: absolute;
  right: 20px;
  bottom: 60px;
}
.movie-item .buy {
  text-align: center;
  color: #ef4238;
  font-size: 14px;
  line-height: 39px;
  cursor: pointer;
}
.movie-item .buy:hover {
  color: #efefef;
  background-color: #ef4238;
}
.ranking-index1 {
  border: 1px solid #efefef;
  margin-bottom: 10px;
}

.ranking-index1 a {
  text-decoration: none;
  display: block;
  margin-bottom: 3px;
}
.ranking-index2 {
  width: 250px;
  float: left;
  height: 80px;
  line-height: 80px;
  margin-top: 20px;
  text-align: center;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}
.ranking-index2 i {
  float: left;
}
.ranking-index2:hover {
  background-color: #efefef;
}
.ranking-index2 a {
  display: block;
  padding-top: 5px;
  float: right;
  padding-left: 4px;
  text-decoration: none;
  margin-bottom: 8px;
  color: #333;
}
.hot-movie-name {
  width: 160px;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  text-align: center;
}

.ranking-index3 {
  padding-top: 30px;
  padding-left: 20px;
  height: 50px;
  float: left;
}
.ranking-index3:hover {
  background-color: #efefef;
}
.ranking-index3 a {
  display: block;
  text-decoration: none;
  margin-bottom: 8px;
  float: right;
  padding-left: 4px;
  padding-top: 5px;
  color: #333;
}
.ranking-index4 {
  padding-top: 30px;
  padding-left: 20px;
  height: 50px;

  float: left;
}
.ranking-index4:hover {
  background-color: #efefef;
}
.ranking-index4 a {
  display: block;
  text-decoration: none;
  margin-bottom: 8px;
  float: right;
  padding-left: 4px;
  padding-top: 5px;
  color: #333;
}
.ranking-index5 {
  padding-top: 30px;
  padding-left: 20px;
  height: 50px;

  float: left;
}
.ranking-index5:hover {
  background-color: #efefef;
}
.ranking-index5 a {
  display: block;
  text-decoration: none;
  float: right;
  padding-top: 5px;
  padding-left: 4px;
  color: #333;
}
.hot-movie-list {
  height: 500px;
}

.ranking-top-left {
  float: left;
  margin-bottom: 11px;
}
.ranking-top-left img {
  height: 78px;
  width: 90px;
}
.ranking-top-right {
  position: relative;
  /* overflow: hidden; */
  /* display: inline; */
  width: 238px;
  height: 78px;
}
.ranking-movie-name {
  text-overflow: ellipsis;
}
.movie-info2 {
  border: 1px solid #efefef;
}
.ranking-top-moive-name {
  position: absolute;
  left: 100px;
  top: 10px;
  font-size: 16px;
  color: #333;
  line-height: 1.4;
  text-overflow: ellipsis;
  vertical-align: center;
}
.ranking-top-wish {
  position: absolute;
  top: 60px;
  left: 100px;
  color: #ef4238;
  font-size: 14px;
  line-height: 1;
}
i {
  color: #ef4238;
  display: inline-block;
  width: 20px;
  font-size: 24px;
}
.ticket {
  position: relative;
  height: 100px;
  text-align: center;
  border: 1px solid #efefef;
  margin-bottom: 40px;
}

.ticket::before {
  content: "今日大盘";
  float: left;
  width: 20px;
  height: 83px;
  padding: 10px;
  color: #fff;
  background-color: #ef4238;
  text-align: center;
  font-weight: 400;
  font-size: 17px;
  line-height: 21px;
}
.ticket p {
  width: 340px;
  height: 40px;
}
.ticket .meta-info {
  position: absolute;
  color: #999;
  font-size: 14px;
  top: 68px;
}
.ticket p .more {
  color: #ef4238;
  margin-top: 40px;
  float: right;
  line-height: 16px;
  text-decoration: none;
}
.ticket p .more .icon {
  margin: 0 auto;
}
.ticket .font {
  position: absolute;
  font-size: 30px;
  font-weight: 700;
  color: #ef4238;
  left: 110px;
  top: 27px;
}
.ticket .num {
  position: absolute;
  color: #ef4238;
  left: 230px;
  top: 40px;
}
.box-total-wrapper {
  margin-left: 0px;
  width: 500px;
  margin-top: 100px;
}
.pull-right {
  background-color: #fff;
}
.textcolor_orange h2 {
  padding-left: 50px;
  color: #ffb400 !important;
  float: left;
  font-size: 26px;
  font-weight: 100;
}

.pane3-header .panel-more a {
  /* float: right; */
  padding-top: 10px;
  position: relative;
  left: 100px;
  text-decoration: none;
  color: #ffb400 !important;
  font-size: 14px;
  line-height: 16px;
}
.pane3-header .panel-more a:hover {
  text-decoration: underline;
}

.ranking-index-1 a {
  display: block;
  width: 380px;
  padding-top: 50px;
  padding-left: 45px;
  text-decoration: none;
}

.ranking2-top-left {
  float: left;
  width: 140px;
}
.ranking2-top-left img {
  width: 140px;
  height: 194px;
}
.ranking2-top-right {
  border: 1px solid #efefef;
  float: right;
  width: 218px;
  height: 191px;
  margin-top: 20px;
  margin-right: 20px;
  cursor: pointer;
}
.ranking2-top-right:hover {
  background-color: #efefef;
}
.ranking-top-right-main {
  float: left;
  left: 200px;
  padding-left: 20px;
  margin-top: 60px;
  font-size: 18px;
  color: #333;
  line-height: 1.4;
}
.ranking-top-right-main .time {
  margin-top: 12px;
  color: #999;
  font-size: 15px;
}
.ranking-top-right-main .want {
  margin-top: 12px;
  font-size: 14px;
  color: #ffb400;
}
.ranking-index-2 {
  height: 185px;
}
.ranking-index-2 a {
  position: relative;
  float: left;
  border: 1px solid #efefef;
  display: block;
  text-decoration: none;
  width: 170px;
  height: 0px;
  margin-top: 1px;
  margin-left: 45px;
}
.ranking-index-2 a:hover {
  background-color: #efefef;
}
.ranking-num-info4 {
  float: right;
  color: #fdb863;
  font-size: 14px;
}
.name-link {
  font-size: 18px;
  color: #333;
  margin-top: 8px;
  margin-left: 10px;
  margin-right: 10px;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}
.rankingindex {
  position: absolute;
  padding-left: 4px;
  width: 16px;
  line-height: 20px;
  color: #fff;
  background-color: #ffb400;
}
.ranking-num-info {
  display: inline-block;
  font-size: 14px;
  margin-top: 6px;
  margin-left: 10px;
  color: #fdb863;
}
.ranking-index-3 a {
  display: block;
  text-decoration: none;
  border: 1px solid #efefef;

  margin-top: 30px;
  float: right;
}
.ranking-index-3 a:hover {
  background-color: #efefef;
}
.name-link3 {
  font-size: 18px;
  color: #333;
  margin-top: 8px;
  margin-left: 10px;
  margin-right: 10px;
}
.ranking-num-info3 {
  display: inline-block;
  font-size: 14px;
  margin-top: 6px;
  margin-left: 10px;
  color: #fdb863;
}
.ranking-index-4 a {
  display: block;
  height: 55px;
  width: 360px;
  margin-top: 20px;
  margin-left: 45px;
}
.ranking-index-4 a:hover {
  background-color: #efefef;
}
.ranking-index-4 {
  height: 55px;
  line-height: 55px;
  float: left;
}
.hopeindex {
  font-size: 18px;
  color: #999;
}
.ranking-hope-movie {
  display: inline-block;
  overflow: hidden;
  text-overflow: ellipsis;
  font-size: 16px;
  color: #333;
  vertical-align: top;
  padding-left: 5px;
}
.hot-panel {
  width: 100vw;
  height: 540px;
  /* background-color: red; */
  position: relative;
}
.hot-panel-header {
  width: 1300px;
  height: 50px;
  line-height: 100px;
  background-color: black;
  height: 100px;
  margin: 0;
}
.hot-panel-header-title h2 {
  display: inline;
  font-size: 26px;
  color: #ef4238;
  position: relative;
  left: -400px;
  font-weight: 300;
}
.hot-panel-header-more {
  font-size: 14px;
  line-height: 16px;
  float: right;
  margin-top: 10px;
}
.hot-panel-header-more a {
  text-decoration: none;
  color: #ef4238 !important;
}
.hot-panel-header-more a .icon {
  float: right;
}
.hot-panel-header-more a:hover {
  text-decoration: underline;
}
.panel-subtitle a {
  font-size: 20px;
  color: #333;
  text-decoration: none;
  line-height: 26px;
  margin-left: 20px;
  position: relative;
  left: -250px;
  font-weight: 300;
}
.panel-subtitle a:hover {
  text-decoration: underline;
}
.hot-panel-content {
  position: relative;
  left: 200px;
  width: 750px;
  height: 350px;
}
.movie-poster-long {
  position: relative;
  width: 350px;
  height: 220px;
}
.movie-poster-short {
  position: relative;
  width: 160px;
  height: 220px;
}
.movie-poster-short .fl {
  float: right;
}
.hot-movie-item {
  float: left;
  margin-top: 20px;
  margin-left: 20px;
}
.hot-movie-item2 {
  float: left;
  margin-left: 25px;
  margin-top: 40px;
  padding-bottom: 30px;
}
.hot-movie-item2 .buy {
  text-align: center;
  color: #ef4238;
  font-size: 14px;
  line-height: 39px;
  cursor: pointer;
}
.hot-movie-item2 .buy:hover {
  background-color: #ef4238;
  color: #efefef;
}
.movie-info {
  position: absolute;
  width: 350px;
  height: 24px;
  bottom: 7px;
  left: 10px;
}
.movie-title-padding {
  font-size: 18px;
  color: #f5f1f1;
}
.movie-score i {
  position: absolute;
  font-size: 18px;
  right: 25px;
  color: #ef4238;
}
.movie-title-padding2 {
  width: 160px;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  text-align: center;
  color: black;
}
.movie-score2 i {
  position: absolute;
  font-size: 24px;
  right: 25px;
  color: #ef4238;
  bottom: 1px;
}
</style>