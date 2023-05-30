<script setup>
</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">
        <div class="tit">
          <img src="../../assets/img/1.gif" alt="" />
          <span>magnolia影院</span>
        </div>
        <div class="nav">
          <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            active-text-color="#ffd04b"
            text-color="#fff"
            style="background: url(../assets/img/background.jpg)"
          >
            <el-menu-item index="1" >首页</el-menu-item>
            <el-menu-item index="2" >电影</el-menu-item>
            <el-menu-item index="3" >榜单</el-menu-item>
          </el-menu>
        </div>
        <div class="search">
            <el-input placeholder="请输入内容" style="width: 200px;" class="input-with-select">
              <el-button class="btn" type="primary" >登录</el-button>
        </el-input>
        </div>
        <!-- <div class="user" >
            <img  src="" alt="user">
            <span>登录</span>
        </div> -->
        </el-header
      >
      <el-main class="main">Main</el-main>
    </el-container>
  </div>
</template>


<style scoped>
.header {
  width: 100vw;
  background-color: red;
  height: 58px;
  line-height: 50px;
  position: absolute;
  top: 0px;
}
.nav{
  height: 50px;
}
.el-menu-demo{
  position: absolute;
  top: 0px;
  text-decoration: none;
  left: 43%;
  border: none;
}
.tit img {
  width: 40px;
  height: 40px;
  padding-top:10px;
}
.tit {
  position: relative;
  left: -350px;
}
.tit span {
  margin-left: 20px;
  position: relative;
  bottom: 10px;
}
.input-with-select{
  position: absolute;
  top: 10px;
  left: 860px;
}
.btn{
 
 
  background-color: blue;
}

</style>