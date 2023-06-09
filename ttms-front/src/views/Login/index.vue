<script setup>
import { ElMessage } from "element-plus";
import { Login, getCode,ping } from "../../apis/user";
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import Cookies from 'js-cookie'
const router = useRouter();
const LoginType = ref(2);
var expires = new Date()
expires.setTime(expires.getTime()+30*1000)
let user = reactive({
  username: "",
  Password: "",
  Email: "",
  EmailCode: "",
});
let user_info = reactive({
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
  let res = await getCode(user.Email);
  if (res.status_code === 0) {
    ElMessage({
      message: "请前往邮箱查看验证码",
      center: true,
      type: "error",
    });
  }
};
const isLogin = () => {
  if (LoginType.value == 2) {
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
     
      Login(user.Email, user.Password, LoginType.value ).then((res) => {
        
        if (res.data.status_code == 0) {
          //存储token  信息   页面跳转至home home页面顶部显示用户信息
          
          user_info = res.data.data.user_info
          Cookies.set('AccessToken',res.data.data.AccessToken, {expires:1/24 }) 
          localStorage.setItem("RefreshToken", res.data.data.RefreshToken);
          localStorage.setItem("AccessToken", res.data.data.AccessToken);
          localStorage.setItem("user_info", JSON.stringify(res.data.data.user_info));
          if (res.data.data.user_info.Role == "vistor") {
            router.push("/home");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "success",
            });
          }
          if(res.data.data.user_info.Role == "administer"){
            router.push("/layout");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "success",
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
      
      Login( user.Email, user.EmailCode,LoginType).then((res) => {
        if (res.status_code == 0) {
          //存储token  信息   页面跳转至home    home页面顶部显示用户信息
          user_info = res.data.user_info;
          Cookies.set('AccessToken', res.data.data.AccessToken, {expires:1/24 }) 
          localStorage.setItem("RefreshToken", res.data.RefreshToken);
          if (res.data.user_info.Role == "vistor") {
            router.push("/home");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "success",
            }); 
          }
          if(res.data.user_info.Role == "administer"){
            router.push("/layout");
            ElMessage({
              message: "登录成功！",
              center: true,
              type: "success",
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
    <div class="color"></div>
        <div class="color"></div>
        <div class="color"></div>
            <div class="logo2">
                <img src="../../assets/img/logo2.png" alt="">
            </div>
            <div class="logo3">
                <img src="../../assets/img/logo3.png" alt="">
            </div>
            <div class="logo4">
                <img src="../../assets/img/logo4.png" alt="">
            </div>
    
    <el-form>
      <div class="form">
        <el-form-item class="act" label="邮箱 &nbsp&nbsp">
          <el-input
            style="width: 200px;background-color: transparent;"
          class="inp"
            placeholder="请输入您的邮箱"
            v-model="user.Email"
          />
        </el-form-item>
        <el-form-item class="pwd" v-if="LoginType == 2" label="密码 &nbsp&nbsp">
          <el-input
            style="width: 200px"
            placeholder="请输入您的密码"
            v-model="user.Password"
          />
        </el-form-item>
        <el-form-item class="pwd" v-if="LoginType == 1" label="验证码">
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
@import url('../../assets/css/register.css');
template{
  background: linear-gradient(to bottom, #f1f4f9, #dff1ff);
}
.container {
  width: 400px;
  height: 400px;
  border: 1px solid red;
  display: flex;
  border-radius: 30px;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  margin: auto;
 
  align-items: center;
  justify-content: center;
}
.act{
  position: relative;
  left: -50px;
  background-color: transparent;
}
.pwd {
  margin-top: 30px;
  position: relative;
  left: -50px;
}
.btn {
  border-radius: 20px;
  background-color: transparent;
  position: relative;
  top: 0px;
}

.tit {
  position: absolute;
  top: 50px;
  color: black;
  font-size: 26px;
}
.form {
  position: relative;
  left: 12%;
}
.inp{
  background-color: transparent;
}
.logo2{
position: absolute;
left: -300px;
}
.logo3{
  position: absolute;
  left:500px;
}
.logo4{
  position: absolute;
  top: -150px;
  left: -500px;
}
</style>