<script setup>
import { reactive, ref } from "vue";
import { register, isRepeat, getCode } from "../../apis/user";
import { ElMessage } from "element-plus";
const email = ref("");
const user = reactive({
  Email: "",
  EmailCode: "",
  Password: "",
  UserName: "",
  Gender: "",
});
const getTest = async () => {
  let answer = await getCode(user.Email);
  console.log(answer);
};

const blank = async () => {
  if (user.Password == "" || user.Email == "" || user.EmailCode == "") {
    ElMessage({
      message: "账号密码邮箱验证码不能为空!!!",
      center: true,
      type: "error",
    });
  } else {
    //查重
    let result = await isRepeat(user.UserName);
    console.log(result);
    if (result.status_code == 20001) {
      ElMessage({
        message: "该用户名已被注册，请重新输入一个用户名!!!",
        center: true,
        type: "error",
      });
    }
    console.log(result.data.status_code);
    if (result.data.status_code === 20002) {
      
      let res = await register(
        user.Email,
        user.EmailCode,
        user.Password,
        user.UserName,
        user.Gender
      );
      console.log(res);
      if (res.data.status_code == 0) {
        ElMessage({
          message: "注册成功",
          center: true,
          type: "success",
        });
      }
      if (res.data.status_code == 3007) {
        ElMessage({
          message: "验证码错误或过期!!",
          center: true,
          type: "error",
        });
      }
    }
  }
};
</script>
<template>
  <div class="container">
    <div class="tit">magolia影院</div>
    <el-form>
      <el-form-item class="act" label="邮箱">
        <el-input
          v-model="user.Email"
          style="width: 270px"
          placeholder="请输入您的邮箱"
        />
      </el-form-item>
      <el-form-item class="pwd" label="验证">
        <el-input
          v-model="user.EmailCode"
          style="width: 200px"
          placeholder="请输入您的验证码"
        />
        <el-button class="getCode" type="primary" @click="getTest()" round
          >获取验证码</el-button
        >
      </el-form-item>
      <el-form-item class="act" label="用户名">
        <el-input
          v-model="user.UserName"
          style="width: 270px"
          placeholder="请输入您的用户名"
        />
      </el-form-item>
      <el-form-item class="pwd" label="密码">
        <el-input
          v-model="user.Password"
          style="width: 270px"
          placeholder="设置您的密码"
        />
      </el-form-item>
      <el-button class="btn" type="primary" @click="blank()">注册</el-button>
      <el-button class="btn" type="primary">返回登录</el-button>
    </el-form>
  </div>
</template>
<style lang="css" scoped>
@import url('../../assets/css/register.css');
.container {
  width: 400px;
  height: 500px;
  border: 1px solid red;
  position: relative;
  left: 140%;
  top: -50px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pwd {
  margin-top: 30px;
}
.getCode {
  position: absolute;
  left: 170px;
}
.btn {
  position: relative;
  top: 50px;
}
.tit {
  position: absolute;
  top: 220px;
}
</style>