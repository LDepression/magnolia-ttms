
<template>
  <div class="invite">
    <!--1 搜索区域 -->
    <div class="header">
      
      <!-- <el-button type="primary" round>
        <router-link to="/add" style="color:#fff">页面添加</router-link>
      </el-button> -->
    </div>  
    <!-- 2 表格区域 -->
    <div class="wrapper">
    <el-table
    :data="tableData"
    stripe
    border
    style="width: 100%">
     <el-table-column
      type="selection"
      width="55">
    </el-table-column>
    <el-table-column
      prop="cinemaname"
      label="影厅"
      width="250">
    </el-table-column>
    <el-table-column
      prop="name"
      label="电影"
      width="350"
      show-overflow-tooltip>
    </el-table-column>
        <el-table-column
      prop="price"
      label="价格"
      width="120">
    </el-table-column>
    <el-table-column
      prop="start_at"
      label="开始时间"
      width="300">
    </el-table-column>
        <el-table-column
      prop="version"
      label="版本">
    </el-table-column>
    
      <el-table-column
      label="操作"
      width="180  ">
      <template v-slot="scope">
        <el-button
          size="mini"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
    </div> 
        <!-- 3 分页展示 -->
    <my-pagnation :total="total" :pageSize="pageSize" @changePage='changePage'></my-pagnation>  
  </div>
</template>

<script setup>
import { onMounted, reactive } from 'vue'
import MyPagnation from '../../../components/Mypagantion.vue'
const dialogVisible = ref(false)
const total = ref(10)
const tableData = reactive([])
const pageSize = ref(1)
const add = () => {
  dialogVisible = true
}
const changeDialog = () => {
  dialogVisible = false
}
const http = (page) => {
  //getPlanpage  获取演出计划
}
const handleDelete = (index,rows) => {

}
onMounted(()=>{http(1)})
</script>

<style scoped>
  .invite {
    margin: 20px;
  }
</style>