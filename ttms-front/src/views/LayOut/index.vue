<script setup>
import { reactive } from "vue"
import {useRouter} from 'vue-router'
const user_info = reactive({})
const router = useRouter()
if(localStorage.getItem("user_info")){
  user_info.value = localStorage.getItem("user_info")
}else{

}
</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">moliga影院后台管理系统</el-header>
      <el-container>
        <el-aside class="aside" width="200px">
          <el-row class="tac">
            <el-col class="tac" :span="12">
              
              <el-menu
                text-color="#fff"
                default-active="2"
                class="el-menu-vertical-demo"
                @open="handleOpen"
                @close="handleClose"
                
              >
                <el-sub-menu index="1" class="bigItem" >
                  <template #title>
                    <el-icon><location /></el-icon>
                    <span>剧院管理</span>
                  </template>
                    <el-menu-item class="item" @click="router.push('/studio')" index="1-1">演出厅管理</el-menu-item>
                    <el-menu-item class="item" @click="router.push('/play')"  index="1-2">剧目管理</el-menu-item>
                    <el-menu-item class="item" @click="router.push('/playSchedule')" index="1-3">演出计划</el-menu-item>
                    
                </el-sub-menu>
                <el-sub-menu index="2">
                  <template #title>
                    <el-icon><location /></el-icon>
                    <span>用户管理</span>
                  </template>
                    <el-menu-item class="item" @click="router.push('/employee')" index="2-2">员工管理</el-menu-item>
                    <el-menu-item class="item" @click="router.push('/user')" index="2-3">用户管理</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="3">
                  <template #title>
                    <el-icon><location /></el-icon>
                    <span>票务管理</span>
                  </template>
                    <el-menu-item class="item" index="3-1">售票管理</el-menu-item>
                    <el-menu-item class="item" index="3-3">销售统计</el-menu-item>
                </el-sub-menu>
                <el-sub-menu index="5">
                  <template #title>
                    <el-icon><location /></el-icon>
                    <span>权限管理</span>
                  </template>
                    <el-menu-item class="item" @click="router.push('/character')" index="5-2">角色管理</el-menu-item>
                </el-sub-menu>
              </el-menu>
            </el-col>
          </el-row>
        </el-aside>
        <el-main class="main"><router-view></router-view></el-main>
      </el-container>
    </el-container>
  </div>
</template>
<style lang="css" scoped>
.aside {
  background-color: #2d3a4b;
  height: 100vh;
  color: #fff;
}
.main{
  /* background: linear-gradient(to bottom, #f1f4f9, #dff1ff); */
}
.header {
  height: 10vh;
  width: 100vw;
  background-color: black;
  line-height: 10vh;
}
/* .tac{
  color: #fff;
} */
.el-menu-vertical-demo{
  width: 201px;
  background-color: #2d3a4b;
  
}
/* .bigItem:hover{
  background-color: blue !important;
} */
.item{
  background-color: #2e3a4b;
  text-indent: 25px;
}
</style>    