<script setup>
import { onMounted, reactive, ref } from "vue";
import { findUser } from "../../../apis/user";
import { useRouter } from "vue-router";
import {getMovieDetail,deleteMovie,getAllMovie} from '../../../apis/movie'
const UserName = ref("");
const router = useRouter()
const tableData = reactive({
  data:[
  
  {
    name: "妇联",
    area: "美国",
    duration: "男",
    action: "我是傻逼",
  },
  {
    name: "妇联",
    area: "美国",
    duration: "男",
    action: "我是傻逼",
  },
  {
    name: "妇联",
    area: "美国",
    duration: "男",
    action: "我是傻逼",
  },
  {
    name: "妇联",
    area: "美国",
    duration: "男",
    action: "我是傻逼",
  },
]
});
const getMovie = async () => {
  // let res = await getMovieDetail()
  let res = await getAllMovie()
  if(res.data.status_code == 0){
    tableData.data.splice(0,tableData.data.length,...res.data.data.movieInfo)
  }
  console.log(res);
}
// const delete = async (movie_id) => {
//   let res = await deleteMovie(movie_id)
// }
const user = reactive({})
const delMovie = async (movieID) => {
  let res = await deleteMovie(movieID)
  getMovie()
  console.log(res);
}
  // tableData.data = res.data.data

onMounted(()=>{getMovie()})
</script>
<template>
  <div class="search">
   电影信息
   <el-button class="btn" type="primary" @click="router.push('/addMovie')" >添加电影</el-button>

  </div>
  <div class="bigTable">
    <div class="table">
      <el-table :data="tableData.data" border style="width: 100%">
        <el-table-column prop="name" label="电影名称" width="168" />
        <el-table-column prop="director" label="电影导演" width="168">
        </el-table-column>
        <el-table-column prop="duration" label="上映时间" width="168" />
        <el-table-column prop="action" label="操作" width="168">
          <template v-slot="{row}">
            <el-button class="btn" type="primary" @click="router.push({name:'movieDetail',params:{movieID}})">修改信息</el-button>
          <el-button class="btn" type="danger" @click="delMovie(row.movieID)" >删除电影</el-button>
         
        </template>
          <!-- <el-button class="btn" type="primary" >修改信息</el-button>
          <el-button class="btn" type="danger" @click="delMovie()" >删除电影</el-button> -->
        </el-table-column>
      </el-table>
    </div>
  </div>

  <el-pagination
    :page-size="6"
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
  margin-top: 10px;
}
.example-pagination-block{
  position: relative;
  left: 300px;
}
.example-pagination-block + .example-pagination-block {
  margin-top: 10px;
}
.example-pagination-block .example-demonstration {
  margin-bottom: 16px;
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