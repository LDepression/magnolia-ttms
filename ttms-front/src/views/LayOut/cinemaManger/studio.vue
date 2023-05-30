<script setup>
import { reactive, ref } from "vue";
const input = ref("");
const tableData = reactive([
  {
    date: "6816685683",
    name: "Tom",
    address: "No. 189,",
  },
  {
    date: "6816685682",
    name: "Tom",
    address: "No. 189, ",
  },
  {
    date: "6816685684",
    name: "Tom",
    address: "No. 189, ",
  },
  {
    date: "6816685681",
    name: "Tom",
    address: "No. 189, ",
  },
]);
</script>
<template>
  <div class="search">
    <span>演出厅名称: </span>
    <el-input
      v-model="input"
      style="width: 300px"
      placeholder="请输入您要查找的演出厅"
    />
    <el-button class="btn" type="primary">查询</el-button>
    <el-button type="danger">新增</el-button>
  </div>
  <div class="bigTable">
  <div class="table">

    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="date" label="演出厅管理" width="168" />
      <el-table-column prop="name" label="演出厅名称" width="168" />
      <el-table-column prop="address" label="座位行数" width="168"/>
      <el-table-column prop="address" label="座位列数" width="168"/>
      <el-table-column prop="address" label="演出厅描述" width="168"/>
      <el-table-column prop="address" label="操作" width="168"/>
    </el-table>
  </div>
</div>
</template>
<style scoped>
.search {
  /* position: relative;
  left: -68px; */
}
.btn {
  margin-left: 30px;
}
.table{
  margin-top: 30px;

}
.bigTable{
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>