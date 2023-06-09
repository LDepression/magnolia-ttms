<template>
  <div class="cardcup">
    <el-carousel :interval="4000" type="card" height="300px"  class="card">
        <el-carousel-item v-for="item,index in swiperImg" :key="index">
            <img :src='item.url' style="width:100%">
        </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
    name:'Swiper',
    data(){
        return {
            swiperImg:[
              {url: new URL('../assets/img/1.gif', import.meta.url).href},
              {url: new URL('../assets/img/2.gif', import.meta.url).href},
              {url: new URL('../assets/img/3.gif', import.meta.url).href},
              {url: new URL('../assets/img/4.gif', import.meta.url).href},
              {url: new URL('../assets/img/5.gif', import.meta.url).href}
            ],
        }
    },
}
</script>

<style scoped>
.cardcup{
    /* display:flex; */
    padding-top: 30px;
    padding-bottom: 0px;
   
    justify-content: center;
    /* position: relative;
    top: -80px;
    left: -20px; */
    background-color: #fff;
}
.card{
    width: 90%;
    margin: auto;
}
</style>