<<<<<<< HEAD
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
=======
<script setup>
import { ElMessage } from "element-plus";
import { ref,reactive } from "vue";
const user = reactive({
  username: "",
  password: "",
  email: "",
});
const isBlank = () => {
  if (user.password == "" && user.email != "") {
    ElMessage({
      message: "密码不能为空!!",
      center: true,
      type: "error",
    });
  }
  if (user.email == "" && user.password != "") {
    ElMessage({
      message: "邮箱不能为空!!",
      center: true,
      type: "error",
    });
  }
  if (user.password == "" && user.email == "") {
    ElMessage({
      message: "账号密码不能为空!!",
      center: true,
      type: "error",
    });
  }
  // else {
  //     let {email,password} = this.user
  //     this.$api.getLogin({
  //         email,password
  //     })
  //     .then(res => {
  //         if(res.data.status_code == 0) {
  //             //1.存储登录信息 2.跳转页面 3.顶部区域显示用户
  //             this.info=''
  //             let obj = {
  //                 user:{
  //                     username:res.data.data.username,
  //                     avatar:res.data.data.avatar,
  //                     userId:res.data.data.user_id,
  //                     privilege:res.data.data.privilege
  //                 },
  //                 token:res.data.data.refresh_token
  //             }
  //             // console.log(obj);
  //             localStorage.setItem('token',obj.token)
  //             localStorage.setItem('user',JSON.stringify(obj.user))
  //             // //跳转
  //             if (res.data.data.privilege=='用户') {
  //                 this.$router.push({path:'/'})
  //                this.$message({
  //                         message: '登录成功',
  //                         center: true,
  //                         type: 'success'
  //                         });
  //             }
  //             if(res.data.data.privilege=='管理员') {
  //                 this.$router.push({ path: "/layout" })
  //                 this.$message({
  //                         message: '登录成功',
  //                         center: true,
  //                         type: 'success'
  //                         });
  //             }
  //         }
  //         else{
  //             //账号和密码错误
  //             this.info='账号或密码错误'
  //         }
  //     })
  //     .catch (err=>{
  //         console.log(err);
  //     })
  // }
};
</script>

<template>
  <div class="container">
    <div class="tit">magolia影院</div>
    <el-form>
      <el-form-item class="act" label="邮箱">
        <el-input style="width: 200px" placeholder="请输入您的邮箱" v-model="user.email" />
      </el-form-item>
      <el-form-item class="pwd" label="密码">
        <el-input style="width: 200px" placeholder="请输入您的密码" v-model="user.password" />
      </el-form-item>
      <el-button class="btn" type="primary" @click="isBlank()">登录</el-button>
      <el-button class="btn" type="primary" @click="$router.push('/register')"
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
  top: 220px;
}
>>>>>>> 73d85f9 (添加注册功能)
</style>