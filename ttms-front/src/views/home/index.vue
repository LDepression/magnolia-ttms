<<<<<<< HEAD
<script setup>
</script>
=======
<<<<<<< HEAD
<script setup></script>
>>>>>>> 7efc7b64e6d16b80b94af0028ebc52c104220b5c
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
<<<<<<< HEAD


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

=======
<style lang="css">
@import url(../../assets/css/bootstrap-theme.css);
@import url(../../assets/css/bootstrap.css);
@import url(../../assets/css/font-awesome.min.css);
@import url(../../assets/css/slick.css);
@import url(../../assets/css/slicknav.css);
@import url(../../assets/css/animate.css);
@import url(../../assets/css/theme.css);
=======
<script setup></script>
<template>
  <header id="site-header">
    <div class="container">
      <div class="row">
        <div class="col-md-2" style="margin: 5px">
          <a href="index.html"
            ><img src="../images/logo.png" width="127" height="48"
          /></a>
        </div>
        <div class="col-md-9">
          <nav id="site-nav" class="nav navbar-default">
            <ul class="nav navbar-nav">
              <li><a href="index.html">主页</a></li>
              <li><a href="board.html">榜单</a></li>
              <li><a href="contact.html">联系我们</a></li>
              <li><a href="login.html">登陆</a></li>
            </ul>
          </nav>
        </div>
      </div>
    </div>

    <!--轮播图-->
  </header>
  <div class="main-slider-wrapper clearfix">
    <div id="main-slider">
      <div class="slide"><img src="../images/slider/1.jpg" /></div>
      <div class="slide"><img src="../images/slider/2.jpg" /></div>
      <div class="slide"><img src="../images/slider/3.jpg" /></div>
      <div class="slide"><img src="../images/slider/4.png" /></div>
    </div>
    <div id="slider-contents">
      <div class="container text-center">
        <div class="jumbotron">
          <h1>寻找您的理想剧目</h1>
          <div class="contents clearfix">
            <p>
              汉唐剧院是观众喜爱的观影平台，为您提供了在线详情服务。<br />
              同时还为您提供剧目预告片、票房查询、剧目排行榜、剧目资讯等信息。
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!--搜索-->
  <div id="advance-search" class="main-page clearfix">
    <div class="container">
      <button class="btn top-btn">寻找您的剧目</button>
      <form action="#" id="adv-search-form" class="clearfix">
        <input
          type="text"
          id="subject"
          name="filmname"
          placeholder="请输入剧目名 *"
          aria-required="true"
          required="required"
        />
        <button type="submit" class="btn btn-primary btn-lg">查询</button>
      </form>
    </div>
  </div>

  <!--剧目列表-->
  <section id="home-property-listing">
    <div class="container">
      <div class="row">
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/1.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">我和我的家乡</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>喜剧片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=1'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/2.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">姜子牙</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>动漫片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=2'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/3.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">夺冠</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>生活片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=3'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/4.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">急先锋</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>动作片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=4'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/5.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">七号房的礼物</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>生活片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=5'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/6.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">天道王</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>战争片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=6'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/7.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">一点就到家</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>生活片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=7'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
        <div class="col-lg-3 col-sm-6 layout-item-wrap">
          <article class="property layout-item clearfix">
            <figure class="feature-image">
              <a class="clearfix zoom" href="single-property.html"
                ><img
                  data-action="zoom"
                  src="../images/property/8.jpg"
                  alt="Property Image"
              /></a>
            </figure>
            <div class="property-contents clearfix">
              <header class="property-header clearfix">
                <div class="pull-left">
                  <h6 class="entry-title">喜宝</h6>
                  <span class="property-location">
                    <i class="fa fa-map-marker"></i>生活片
                  </span>
                </div>
                <button
                  class="btn btn-primary btn-price pull-right btn-3d"
                  data-hover="详情"
                  onclick="javascript:window.location='schedule.html?playId=8'"
                >
                  <strong>详情</strong>
                </button>
              </header>
            </div>
          </article>
        </div>
      </div>
    </div>
  </section>

  <footer id="footer">
    <div class="site-footer">
      <div class="container">
        <div class="row">
          <div class="col-md-4 col-sm-6">
            <section class="widget address-widget clearfix">
              <h4 class="title">联系方式</h4>
              <ul>
                <li>
                  <i class="fa fa-map-marker"></i>
                  <a href="http://cs.xiyou.edu.cn/"
                    >西安市长安区西长安街618号东区计算机学院</a
                  >
                </li>
                <li><i class="fa fa-phone"></i>029-88166712</li>
                <li><i class="fa fa-envelope"></i> shuxf@xiyou.edu.cn</li>
                <li>@ Copyright &copy; 2020. 西安邮电大学计算机学院.</li>
              </ul>
            </section>
          </div>
        </div>
      </div>
    </div>
  </footer>
  <a href="#top" id="scroll-top"><i class="fa fa-angle-up"></i></a>
</template>
<style lang="css">
@import url(../../assets/css/bootstrap-theme.css);
@import url(../../assets/css/bootstrap.css);
@import url(../../assets/css/font-awesome.min.css);
@import url(../../assets/css/slick.css);
@import url(../../assets/css/slicknav.css);
@import url(../../assets/css/animate.css);
@import url(../../assets/css/theme.css);
>>>>>>> 73d85f9 (添加注册功能)
>>>>>>> 7efc7b64e6d16b80b94af0028ebc52c104220b5c
</style>