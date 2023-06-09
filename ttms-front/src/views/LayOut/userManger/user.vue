<script setup>
import { computed, onMounted, onUnmounted, onUpdated, reactive, ref ,  } from "vue";
import { findUser } from "../../../apis/user";
import { getUsers,deleteUser } from "../../../apis/manager";
import { mapState, useStore,mapMutations} from 'vuex'
import emitter from '../../../mitt/mitt'

const UserName = ref("");
const page = ref("3");


const tableData = reactive({
  data: [],
});

const user = reactive({});
const UserFind = async (UserName) => {
  let res = await findUser(UserName);
  console.log(res.data);
  tableData.data.splice(0,tableData.data.length,res.data.data)
};
const getAllUsers = async () => {
  let res = await getUsers();
  console.log(...res.data.data.user_infos);
   tableData.data.splice(0,tableData.data.length,...res.data.data.user_infos)
 
  // console.log(tableData.data);
  
  // setUserList(tableData.data)
};
const delUser = async (userID) => {
  let res = await deleteUser(userID);
    
  getAllUsers()
  console.log(res);
}

onMounted(() => {
  
  getAllUsers()
  
}
 
);
onUnmounted(()=>{
  
})
</script>
<template>
  <div class="search">
    <span>用户名称: </span>
    <el-input
      v-model="UserName"
      style="width: 300px"
      placeholder="请输入您要查找的用户"
    />
    <el-button class="btn" @click="UserFind(UserName)" type="primary"
      >查询</el-button
    >
  </div>
  <div class="bigTable">
    <div class="table">
      <el-table :data="tableData.data" border style="width: 100%">
        <el-table-column prop="username" label="用户名字" width="168" />
        <el-table-column prop="avatar" label="用户头像" width="168">
          <template v-slot="{row}">
          <img :src="row.avatar" style="width: 50px; height: 50px" />
        </template>
        </el-table-column>
        <el-table-column prop="gender" label="用户性别" width="168" />
        <el-table-column prop="signature" label="个性签名" width="168" />
        <el-table-column prop="role" label="用户权限" width="168">
         <el-button type="primary">修改权限</el-button>
         </el-table-column>
         <el-table-column prop="userID" label="用户删除" width="168">
         <template v-slot="{row}">
          <el-button @click="delUser(row.userID)" type="primary">删除用户</el-button>
         </template>
           
          
         
         </el-table-column>
      </el-table>
    </div>
  </div>
  <el-pagination
    :page-size="4"
    :pager-count="5"
    layout="prev, pager, next"
    :total="tableData.data.length"
    class="pagination"
  />
</template>
<style lang="css" scoped>
.pagination{
  margin-left: 400px;
  margin-top: 50px;
}
.search {
  /* position: relative;
  left: -68px; */
}
.btn {
  margin-left: 30px;
}
.table {
  margin-top: 30px;
}
.bigTable {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>