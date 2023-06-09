<script setup>
import { onMounted, reactive, ref } from "vue";
import { getCinemaList,deleteCinema } from "../../../../apis/cinema";
import {useRouter} from 'vue-router'

const input = ref("");
const page = ref("1");
const router = useRouter()
const tableData = reactive({
  data: [],
});
const getCinema = async () => {
  let res = await getCinemaList();
  if (res.data.status_code == 0) {
  console.log(res.data.data.list);
  localStorage.setItem("cinemaInfo",JSON.stringify(res.data.data.list))
    tableData.data.splice(0,tableData.data.length,...res.data.data.list)
  }
};
const delCinema = async (cinema_id) => {
  let res = await deleteCinema(cinema_id)
  getCinema()
}
const searchList = reactive([])
const edit = (id) => {
  console.log(11);
 
  const target = tableData.data.find(item => item.ID == id);
  
  if (target) {
   
    const cinemaInfo = encodeURIComponent(JSON.stringify(target));
    const path = `/editCinema/${cinemaInfo}`;
   console.log(cinemaInfo);
    router.push({
      path: path
    });
  } else {
    console.error(`未找到 ID 为 ${id} 的元素`);
  }
}
const findCinema = (input) => {
  console.log(111);
  console.log(input);
  searchList.splice(0, searchList.length) // 清空 searchList 数组
  if (input) {
    console.log(2);
    for (let i = 0; i < tableData.data.length; i++) {
      let item = tableData.data[i]
      console.log(item);
      if (item.name.indexOf(input) !== -1) {
        searchList.push(item)
      }
    }
  } else {
    searchList.push(...tableData.data) // 将 tableData.data 的所有项添加到 searchList 数组中
  }
  tableData.data.splice(0, tableData.data.length, ...searchList) // 用搜索结果更新 tableData.data 数组
}
onMounted(() => {
  getCinema();
});
</script>
<template>
  <div class="search">
    <span>演出厅名称: </span>
    <el-input
      v-model="input"
      style="width: 300px"
      placeholder="请输入您要查找的演出厅"
    />
    <el-button class="btn" @click="findCinema(input)" type="primary">查询</el-button>
    <el-button type="danger" @click="router.push({path:'/addCinema'})">新增</el-button>
  </div>
  <div class="bigTable">
    <div class="table">
      <el-table :data="tableData.data" border style="width: 100%">
        <el-table-column prop="name" label="演出厅名称" width="148" />
        <el-table-column prop="ID" label="演出厅编号" width="148" />
        <el-table-column prop="avatar" label="演出厅头像" width="148">
          <template v-slot="{ row }">
          <img :src="row.avatar" style="width: 50px; height: 50px" />
        </template>
        </el-table-column>
        <el-table-column prop="rows" label="座位行数" width="148" />
        <el-table-column prop="cols" label="座位列数" width="148" />
        <el-table-column prop="avatar" label="操作" width="148">
          <template v-slot="{ row }">
      <el-button class="btn3" type="primary" style="width: 50px" @click="delCinema(row.ID)">删除</el-button>
      <el-button class="btn3" type="primary" style="width: 50px" @click="edit(row.ID)">编辑</el-button>
    </template>
        </el-table-column>
        <!-- <el-table-column prop="address" label="演出厅描述" width="118"/>
      <el-table-column prop="address" label="演出厅状态" width="118"/>
      <el-table-column prop="address" label="操作" width="118"/> -->
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
           
.btn3{
  width: 50px;
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