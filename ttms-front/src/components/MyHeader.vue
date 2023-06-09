<template>
  <div class="top">
    <div class="logo">
      <img src="../assets/img/logo.png" alt="logo" />
      <span>mogalian</span>
    </div>
    <div class="nav">
      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" active-text-color="#ffd04b" text-color="#fff" style="background: url(../assets/img/background.jpg)">
        <el-menu-item index="1" @click="changeroute(1)">首页</el-menu-item>
        <el-menu-item index="2" @click="changeroute(2)">电影</el-menu-item>
        <el-menu-item index="3" @click="changeroute(3)">榜单</el-menu-item>
      </el-menu>
    </div>
    <div class="search">
      <el-input
        placeholder="请输入内容"
        v-model="keyvalue"
        class="input-with-select"
      >
        <el-button @click="goSearch" icon="el-icon-search"></el-button>
      </el-input>
    </div>
    <div class="user">
      <img @click="toHomePage" :src="photo" alt="user" />
      <span @click="router.push('/userInfo')">个人中心</span>
    </div>
  </div>
</template>
<script>
import { ref,computed } from "vue";
import {useRouter} from 'vue-router'
import home from "../views/home/index.vue";
import films from "../views/films/index.vue";
import { ElMessage } from "element-plus";
import list from "../views/list/index.vue";
// import Search from "../views/Search";
// import HomePage from "../views/HomePage.vue";

export default {
  setup() {
    const router = useRouter()
    const activeIndex = ref("1");
    const keyvalue = ref("");
    const isLogin = ref(false);

    // const photo = computed(() => {
    //   let user = JSON.parse(localStorage.getItem("user"));
    //   if (user) {
    //     isLogin.value = false;
    //     return user.avatar;
    //   }
    //   return require("../assets/img/user.png");
    // });

    function changeroute(T) {
      switch (T) {
        case 1:
          router.replace({
            path: "home",
          });
          break;
        case 2:
          router.replace({
            path: "film",
          });
          break;
        case 3:
          router.replace({
            path: "list",
          });
          break;
      }
    }

    function tologin() {
      if (!isLogin.value) {
        localStorage.clear();
        router.push({ name: "login" });
      }
    }

    function goSearch() {
      if (keyvalue.value !== "") {
        router.push({ name: "search", query: { keyvalue: keyvalue.value } });
        keyvalue.value = "";
      } else {
        ElMessage({
          message: "输入一点东西",
          type: "warning",
        });
      }
    }

    function toHomePage() {
      router.push({ name: "homepage" });
    }

    return {
      activeIndex,
      keyvalue,
      isLogin,
      home,
      films,
      list,
      router,
      // search,
      // HomePage,
      // photo,
      changeroute,
      tologin,
      goSearch,
      toHomePage,
    };
  },
};
</script>
<style scoped lang="css">
.top{
    display:flex;
    width: 100vw;
    justify-content:center;
    box-shadow: 3px 3px 1px #888888;
    background: url(../assets/img/background.jpg);
    background-size: 100%;
    font-size: 1.5em;
    font-weight: bold;
    align-items:center;
    color: white;
    min-width: 988px;
    margin-bottom: 3px;
    position: relative;
    left: 00px;
    top: -0px;
}
.user{
        cursor: pointer;
        margin-left:50px;
        position: relative;
        font-size: 16px;
        display:flex;
        align-items: center;
}
.user img{
    background-color:#888888;
    border-radius:50%;
    width:30px;
    height:30px;
    margin-right: 8px;
}
.logo{
    height:30px;
    margin-right: 50px;
    display:flex;
}
.logo img{
    height:120%;
}
.logo span{
    align-self: center;
    margin-left: 8px;
    }
.nav{
    margin: 0 100px;
}

</style>