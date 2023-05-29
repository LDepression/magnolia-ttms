<<<<<<< HEAD
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