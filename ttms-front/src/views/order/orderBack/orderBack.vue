<template>
  <div>
    <div class="films">
      <ul class="film">
        <li v-for="(item, index) in movies" :key="item">
          <div class="plan">
            <div class="planmovie"><img :src="item.avatar" /></div>
            <span class="add">
              <el-row>
                <el-button type="primary" @click="add(index)"
                  >添加演出计划</el-button
                >
              </el-row>
            </span>
          </div>
        </li>
      </ul>
    </div>
    <el-dialog
      title="添加剧目"
      :visible.sync="dialogVisible"
      width="70%"
      :before-close="handleClose"
    >
      <!-- 内容区域 -->
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        class="demo-ruleForm"
      >
        <el-form-item
          label="开始时间(例如:2022-06-23 12:00:00)"
          prop="start_at"
        >
          <el-input v-model="ruleForm.start_at"></el-input>
        </el-form-item>
        <el-form-item label="输入影院" prop="cinema_id">
          <el-input v-model="ruleForm.cinema_id"></el-input> </el-form-item
        ><el-form-item label="输入价格" prop="price">
          <el-input v-model="ruleForm.price"></el-input>
        </el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')"
          >确定</el-button
        >
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form>
    </el-dialog>
    <my-pagnation
      :total="total"
      :pageSize="pageSize"
      @changePage="changePage"
    ></my-pagnation>
  </div>
</template>

<script setup>
import { onMounted, reactive } from "vue";
import MyPagnation from "../../../components/Mypagantion.vue";
const movies = reactive([]);
const total = ref(10);
const pageSize = ref(1);
const movie_id = ref("");
const dialogVisible = ref(false);
const ruleForm = reactive({
  start_at: "",
  cinema_id: "",
  version: "",
  price: "",
  type: [],
});
const rules = reactive({
  start_at: [{ required: true, message: "请输入开始时间", trigger: "blur" }],
  cinema_id: [{ required: true, message: "请输入影厅", trigger: "blur" }],
  version: [{ required: true, message: "请输入版本", trigger: "blur" }],
  price: [{ required: true, message: "请输入价格", trigger: "blur" }],
});
const submitForm = (formName) => {
      const formRef = $refs[formName]
      formRef.validate((valid) => {
        if (valid) {
          console.log(ruleForm)
          $api.getPlan({
            start_at: new Date(ruleForm.start_at).getTime() / 1000,
            cinema_id: parseInt(ruleForm.cinema_id),
            version: ruleForm.version,
            price: parseFloat(ruleForm.price),
            movie_id: movieId.value
          }).then((res) => {
            console.log(res.data)
            if (res.data.status_code === 0) {
              dialogVisible.value = false
              ElMessage({
                message: '添加成功!!',
                center: true,
                type: 'success'
              })
            }
            if (res.data.status_code === 3003) {
              dialogVisible.value = false
              ElMessage({
                message: '时间冲突添加失败!!',
                center: true,
                type: 'error'
              })
            }
            if (res.data.status_code === 1001) {
              dialogVisible.value = false
              ElMessage({
                message: '起始时间晚于当前时间添加失败!!',
                center: true,
                type: 'error'
              })
            }
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }

    const resetForm = (formName) => {
      const formRef = $refs[formName].resetFields()
      
    }
    const add = (index) => {
      dialogVisible = true
      movie_id = movies[index].id
    }
    const http = (page,pageSize) => {
      //调用getMovieList接口
    }
    const changePage = (num) => {
      http(num)
    }
    onMounted(()=>{http(1)})
</script>

<style>
.plan {
  position: relative;
}
.films {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  height: 1000px;
}

.film li {
  float: left;
  height: 270px;
  width: 250px;
  margin-left: 100px;
  margin-top: 60px;
}
.planmovie img {
  width: 218px;
  height: 300px;
}
.add {
  position: absolute;
  bottom: 4px;
}
.add .el-row {
  margin-left: 50px;
  /* margin-top: 320px; */
}
</style>