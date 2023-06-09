<template>
  <div>
    <Myheader/>
    <router-view></router-view>
    <Footer></Footer>
  </div>
</template>

<script setup>
// import MyHeader from "./MyHeader.vue";
import Myheader from "./MyHeader.vue";
import Footer from "./Footer.vue"

    // created () {
    //   if (window.localStorage) {
        
    //   }
    // }

</script>

<style scoped>

</style>