<template>
  <div>
    <div class="header">
        <h3>
          <span @click="list='listByVisit'" :class="{textcolor:list=='listByVisit'}">热映口碑榜&nbsp;&nbsp;&nbsp;</span>
          <span @click="list='listByFollow'" :class="{textcolor:list=='listByFollow'}">最受期待榜</span>
        </h3>
    </div>
      <component :is="list"></component>
  </div>
</template>

<script >
import { defineComponent, ref } from 'vue'
import listByVisit from '../../components/listByVisit.vue'
import listByFollow from "../../components/listByFollow.vue"

export default defineComponent({
    components: {
      listByVisit,
      listByFollow
    },
    setup() {
      const list = ref('listByVisit')
      return { list }
    }
})
</script>

<style scoped>
.textcolor{
  color: #eae852;
}
.header {
  position: relative;
  top: 0px;
  background: url(../../assets/img/background2.jpg);
  display: flex;
  align-items: center;
  justify-content: center;
  color:#fff;
  height:60px;
  width:950px;
  margin:0 auto;
  cursor: pointer;
}

</style>