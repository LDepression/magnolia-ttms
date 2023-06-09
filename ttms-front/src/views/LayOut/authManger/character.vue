<script setup>
import {onBeforeUnmount,computed, onMounted, onUnmounted, onUpdated, watch,reactive,toRefs, watchEffect} from 'vue'
import { useStore } from "vuex";
import {getUsers,updateRole} from '../../../apis/manager'
const store = useStore()
const userList = computed(() => store.state.userList)

const tableData =reactive({data: []})
const getAllUsers = async () => {
  let res = await getUsers();
  tableData.data.splice(0,tableData.data.length,...res.data.data.user_infos)
  
};
const updaterole = async (Username) => {
  console.log(Username);
  let res = await updateRole(Username)
  getAllUsers()
  console.log(res);
}

onMounted(()=>{ 
  getAllUsers()
})


</script>
<template>
  <div class="board">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>权限管理</el-breadcrumb-item>
        <el-breadcrumb-item>角色信息管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <el-button type="primary" >添加角色</el-button>
   
    <div class="bigTable">
    <div class="table">
    <el-table :data="tableData.data" border style="width: 100%">
        <el-table-column prop="username" label="角色用户名" width="168" />
        <el-table-column prop="role" label="角色身份" width="168">
        </el-table-column>
        <el-table-column prop="gender" label="角色描述" width="168" />
        <el-table-column prop="signature" label="操作" width="168">
          <template v-slot="{row}">
          <el-button class="btn1" @click="updaterole(row.username)" type="primary">修改身份</el-button>
        </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
<style>
.table {
  margin-top: 30px;
}
.bigTable {
  display: flex;
  justify-content: center;
  align-items: center;
}
.btn{
  position: relative;
  left: 10px;
}
.btn1{
  position: relative;
  left: 23px;
}
</style>