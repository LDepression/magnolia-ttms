<script setup>import { onMounted, ref } from "vue";
import axios from "axios";
import Cookies from "js-cookie";
import { changeAvatar, fileUpLoad,changeDescription,changeEmail,getCode } from "../../../apis/user";
const user_info = JSON.parse(localStorage.getItem('user_info'))
const handleFileChange = () => {
  // console.log(event);
  const File = event.target.files[0];
  //构造format对象
  const formData = new FormData();
  formData.append("File", File);

  axios({
    method: "POST",
    url: "http://127.0.0.1/api/v1/file/upload",
    data: formData,
    headers: {
      "Content-Type": "multipart/form-data",
      x_token: Cookies.get("AccessToken"),
    },
  })
    .then((response) => {
      
      user_info.AvatarURL = response.data.data.URL;
      localStorage.setItem("user_info",JSON.stringify(user_info))
      changeConfirm()
    })
    .catch((error) => {
      console.error(error);
    });
const changeConfirm = () => {
  console.log(111);
  window.location.reload();
}
const changeInfo = async () => {
  let res = await changeEmail(user_info.email)
  console.log(res);
  if(res.data.status_code == 0){
    user_info.email = newEmail
  }
  let result =await changeDescription(user_info.gender,user_info.signature)
  if(res.data.status_code == 0){
    localStorage.setItem("user_info")
  }
}
let newEmail = ref('')
const getTestCode = async () => {
  let res = await getCode(user_info.email)
  console.log(res);
}
const radio = ref(3);
}</script>



<template>
  <div class="top">基本信息</div>
  <hr />
  <div class="allInfo">
    <div class="avatar">
      <!-- <img :src="user_info.AvatarURL" alt="" /> -->
      <img :src="user_info.AvatarURL" alt="" />
      <div>
        <form action="POST" enctype="multipart/form-data">
          <input type="file" class="file" name="File" @change="handleFileChange" />
        </form>
      </div>
      <el-button class="btn" type="primary" @click="changeConfirm" style="width: 80px"
        >确认修改</el-button>
      <div class="text">支持JPG,PNG格式,且文件需小于1M</div>
    </div>

    <div style="width: 70%;position: relative;left:60px;top:30px">
        <el-form label-width="80px">
          <el-form-item label="用户名" prop="userName">
            <el-input  v-model="user_info.user_name"></el-input>
            <!-- <el-input ></el-input> -->
          </el-form-item>
          <el-form-item label="电子邮箱" prop="email">
            <!-- <el-input v-model="user_info.email"></el-input>
            <el-input  v-model="newEmail"></el-input> -->
            <el-input v-model="user_info.email"></el-input>
            <el-button class="btn3" type="primary" @click="getTestCode" style="width: 80px">获取验证码</el-button>
           <el-input class="newEmail" v-model="newEmail"></el-input>
          </el-form-item>
       
          <el-form-item class="gender" label="性别" prop="sex">
            <el-radio-group v-model="radio">
                <!-- <el-radio-group> -->
                <el-radio :label="3">男</el-radio>
                <el-radio  :label="6">女</el-radio>
            </el-radio-group>
          </el-form-item>
          
          <el-form-item class="signature" label="个性签名" prop="autograph">
            <el-input v-model="user_info.signature"></el-input>
            <!-- <el-input ></el-input> -->
          </el-form-item>
            <el-button class="newInfo" type="primary" @click="changeInfo()">保存信息</el-button>
            <!-- <el-button type="primary">保存信息</el-button> -->
        </el-form>
      </div>

  </div>
</template>
<style lang="css" scoped>
.choose_sex{
    position: relative;
}
.newInfo{
  position: relative;
  top: 120px;
  left: -150px;
}
.newEmail{
  position: relative;
  top: 50px;
  left: -280px;
}
.signature{
  position: relative;
  top: 50px;
}
       
.gender{
  position: relative;
  top: 50px;
}
.password {
  position: relative;
  top: 210px;
  left: 140px;
}
.el-input{
    width: 200px;
}
.file{
  position: relative;
  top: 60px;
  left: 120px;
}
.btn2 {
  position: relative;
  top: 250px;
  left: 150px;
}
.email2{
  position: relative;
  top: 180px;
  left: -10px;
}
.btn3{
  position: relative;
  top: 0px;
  left: 0px;
}

.top {
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  color: red;

}
.email {
  position: relative;
  top: 100px;
  left: 150px;
}
.descript {
  position: relative;
  top: 160px;
  left: 140px;
}
.email span {
  position: relative;
  left: 0px;
  top: 20px;
}
.allInfo {
  width: 100%;
  height: 100%;
  display: flex;
}
.right {
  color: black;
  font-size: 13px;
}
.right span {
  position: relative;
  left: 120px;
}
.username {
  position: relative;
  left: 120px;
  top: 50px;
}
.avatar {
  width: 40%;
  height: 100%;
  background-color: #fff;
}
.radio {
  position: relative;
  top: 80px;
  left: 120px;
}
.avatar img {
  width: 200px;
  height: 200px;
  border-radius: 50%;
  position: relative;
  top: 20px;
}
.text {
  position: relative;
  top: 150px;
  font-size: 14px;
  left: -30px;
  color: black;
}
.btn {
  position: relative;
  top: 100px;
  left: 0px;
}
</style>