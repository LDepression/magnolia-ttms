<script setup>
import {useRouter} from 'vue-router'
const router = useRouter()
</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-aside class="aside" width="200px"><div>
        <h3 style="color:black">个人中心</h3>
          <el-menu>
            <el-menu-item class="item" @click="router.push('/myOrder')">我的订单</el-menu-item>
            <el-menu-item class="item" @click="router.push('/userInfo')">基本信息</el-menu-item>
          </el-menu>
      </div></el-aside>
      <el-main class="main"><router-view></router-view></el-main>
    </el-container>
  </div>
</template>
<style lang="css" scoped>
.aside{
  background-color: rgb(243, 243, 243);
  height: 100vh;
}
.main{
  background-color: #fff;
  width: 82vw;
}
.item{
  text-indent:40px;
  height: 40px;
  background-color:rgb(243, 243, 243);
  color: black;
  position: relative;
  left: 0px;
  width: 199px;
}
.item:hover{
  background-color: red;
}
</style>