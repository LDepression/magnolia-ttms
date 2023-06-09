<script setup>
import { createCinema } from "../../../../../apis/cinema";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { reactive, ref } from "vue";
const router = useRouter();
const cinema = reactive({
  name: "",
  rows: 8,
  cols: 8,
  cinemaID: "",
  siganature: "",
});
const addCinema = async (name, rows, cols) => {
  let res = await createCinema(name, rows, cols);
  console.log(res);
  if (res.data.status_code == 0) {
    ElMessage({
      message: "添加影院成功",
      center: true,
      type: "success",
    });
    cinema.name= "";
    cinema.rows = "";
    cinema.colscols= "";
    cinema.cinemaID= "";
    cinema.siganature = ''
    router.push({path:'/studio'})
  }
};
</script>
<template>
  <div class="container">
    <el-form class="form">
      <el-form-item class="act" label="演出厅编号">
        <el-input
          style="width: 200px"
          type="text"
          v-model="cinema.cinemaID"
          placeholder="请输入演出厅编号"
        />
      </el-form-item>
      <el-form-item class="pwd" label="演出厅名称">
        <el-input
          style="width: 200px"
          type="text"
          v-model="cinema.name"
          placeholder="请输入演出厅名称"
        />
      </el-form-item>
      <el-form-item class="act" label="座位&nbsp&nbsp 行数">
        <el-input
          v-model="cinema.rows"
          style="width: 200px"
          type="number"
          placeholder="请输入座位行数"
        />
      </el-form-item>
      <el-form-item class="act" label="座位&nbsp&nbsp  列数">
        <el-input
          v-model="cinema.cols"
          style="width: 200px"
          type="number"
          placeholder="请输入座位列数"
        />
      </el-form-item>
      <el-form-item class="act" label="演出厅描述">
        <el-input
          style="width: 200px"
          type="text"
          v-model="cinema.siganature"
          placeholder="请输入演出厅描述"
        />
      </el-form-item>
      <el-button class="btn" @click="addCinema(cinema.name, cinema.rows, cinema.cols)" type="primary"
        >提交</el-button
      >
      <el-button class="btn" @click="router.push('/studio')" type="primary">返回</el-button>
    </el-form>
  </div>
</template>
<style scoped>
.container {
  width: 400px;
  height: 500px;
  border: 1px solid black;
  display: flex;
  border-radius: 20px;
  position: absolute;
  left: 200px;
  top: 100px;
  bottom: 0;
  right: 0;
  margin: auto;
  align-items: center;
  justify-content: center;
}
.form {
  margin-top: -50px;
}
.btn {
  position: relative;
  top: 50px;
}
</style>