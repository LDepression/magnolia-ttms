<script setup>
import { ElMessage } from "element-plus";
import { Login, getCode,isRepeat } from "../../apis/user";
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
const LoginType = ref(2);
const user = reactive({
  username: "",
  Password: "",
  Email: "",
  EmailCode: "",
});
const user_info = reactive({
  user_name: "",
  AvatarURL: "",
  Role: "",
});
const changeLogin = () => {
  console.log(LoginType.value);
  if (LoginType.value == 1) {
    LoginType.value = 2;
    return;
  }
  if (LoginType.value == 2) {
    LoginType.value = 1;
    return;
  }
};
const getTestCode = async () => {
  let res = await getCode(Email);
  if (res.status_code === 0) {
    ElMessage({
      message: "请前往邮箱查看验证码",
      center: true,
      type: "error",
    });
  }
};

const isLogin = () => {
  if (LoginType.value == 1) {
    if (user.Password == "" && user.Email != "") {
      ElMessage({
        message: "密码不能为空!!",
        center: true,
        type: "error",
      });
    }
    if (user.Email == "" && user.Password != "") {
      ElMessage({
        message: "邮箱不能为空!!",
        center: true,
        type: "error",
      });
    }
    if (user.Password == "" && user.Email == "") {
      ElMessage({
        message: "账号密码不能为空!!",
        center: true,
        type: "error",
      });
    } else {
      let { Email, Password } = user;
      Login({ Email, Password, LoginType }).then((res) => {
        if (res.status_code == 0) {
          //存储token  信息   页面跳转至home    home页面顶部显示用户信息
          user_info = res.data.user_info;
          localStorage.setItem("AccessToken", res.data.AccessToken);
          localStorage.setItem("RefreshToken", res.data.RefreshToken);
          if (res.data.user_info.Role == "vistor") {
            router.push("/home");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "error",
            });
          }
          if(res.data.user_info.Role == "administer"){
            router.push("/layout");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "error",
            });
          }
        }else{
          ElMessage({
              message: "账号密码错误",
              center: true,
              type: "error",
            });
        }
      });
    }
  } else {
    if (user.EmailCode == "" && user.Email != "") {
      ElMessage({
        message: "验证码不能为空!!",
        center: true,
        type: "error",
      });
    }
    if (user.Email == "" && user.EmailCode != "") {
      ElMessage({
        message: "邮箱不能为空!!",
        center: true,
        type: "error",
      });
    }
    if (user.Password == "" && user.Email == "") {
      ElMessage({
        message: "邮箱验证码不能为空!!",
        center: true,
        type: "error",
      });
    } else {
      let { Email, EmailCode } = user;
      Login({ Email, EmailCode, LoginType }).then((res) => {
        if (res.status_code == 0) {
          //存储token  信息   页面跳转至home    home页面顶部显示用户信息
          user_info = res.data.user_info;
          localStorage.setItem("AccessToken", res.data.AccessToken);
          localStorage.setItem("RefreshToken", res.data.RefreshToken);
          if (res.data.user_info.Role == "vistor") {
            router.push("/home");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "error",
            });
          }
          if(res.data.user_info.Role == "administer"){
            router.push("/layout");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "error",
            });
          }
        }else{
          ElMessage({
              message: "账号密码错误",
              center: true,
              type: "error",
            });
        }
      });
    }
  }
};
</script>

<template>
  <div class="container">
    <div class="tit">magolia影院</div>

    <el-form>
      <div class="form">
        <el-form-item class="act" label="邮箱 &nbsp&nbsp">
          <el-input
            style="width: 200px"
            placeholder="请输入您的邮箱"
            v-model="user.Email"
          />
        </el-form-item>
        <el-form-item class="pwd" v-if="LoginType == 1" label="密码 &nbsp&nbsp">
          <el-input
            style="width: 200px"
            placeholder="请输入您的密码"
            v-model="user.Password"
          />
        </el-form-item>
        <el-form-item class="pwd" v-if="LoginType == 2" label="验证码">
          <el-input
            style="width: 200px"
            placeholder="请输入您的验证码"
            v-model="user.EmailCode"
          />
        </el-form-item>
      </div>
      <el-button
        class="btn"
        type="primary"
        @click="changeLogin"
        style="width: 80px"
        >切换登录</el-button
      >
      <el-button
        class="btn"
        type="primary"
        style="width: 80px"
        @click="getTestCode()"
        >获取验证码</el-button
      >
      <el-button
        class="btn"
        type="primary"
        style="width: 80px"
        @click="isLogin()"
        >登录</el-button
      >
      <el-button
        class="btn"
        type="primary"
        style="width: 80px"
        @click="$router.push('/register')"
        >注册</el-button
      >
    </el-form>
  </div>
</template>
<style lang="css" scoped>
.container {
  width: 400px;
  height: 600px;
  border: 1px solid red;
  display: flex;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
  align-items: center;
  justify-content: center;
}
.pwd {
  margin-top: 30px;
}
.btn {
  position: relative;
  top: 50px;
}
.tit {
  position: absolute;
  top: 100px;
  font-size: 26px;
}
.form {
  position: relative;
  left: 12%;
}
</style>