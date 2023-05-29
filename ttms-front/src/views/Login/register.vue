<<<<<<< HEAD
<script setup>
import {reactive, ref} from 'vue'
const user = reactive({
  email:'',
  testCode:'',
  password:''
})
const blank = () => {
 
            if(user.password==''|| user.email == '' || user.testCode=='')
            {
                 ELMessage({
                                message: '账号密码邮箱验证码不能为空!!!',
                                center:true,
                                type: 'error'
                            })
            }
            // else {
                 
            //         this.$api.isRepeat({
            //             Username:this.user.Username
            //         })
            //         .then(res=>{
            //             console.log(res.data);
            //             if(res.data.status_code==3001)
            //             this.$message({
            //                     message: '用户名已存在!!',
            //                     center:true,
            //                     type: 'error'
            //                 })
            //         })
            //         this.$api.getRegister({
            //             Email: this.user.email,
            //             Password:this.user.Password,
            //             Username:this.user.Username,
            //             VerifyCode:this.user.comfirm,
            //             InviteCode:this.user.invite
            //         })
            //         .then(res=>{
            //             console.log(res);
            //             if(res.data.status_code==0){
            //                 this.$message({
            //                         message: '注册成功',
            //                         center: true,
            //                         type: 'success'
            //                         });
            //                 this.change();
            //             }
            //             if(res.data.status_code==3007){
            //                 this.$message({
            //                     message: '验证码错误或过期!!',
            //                     center:true,
            //                     type: 'error'
            //                 })
            //             }
            //         })
            //         .catch (err => {
            //         console.log(err);
            // })
                
            // }
    
}
</script>
<template>
  <div class="container">
    <div class="tit">magolia影院</div>
    <el-form>
      <el-form-item class="act" label="邮箱">
        <el-input v-model="e_mail" style="width: 270px;" placeholder="请输入您的邮箱" />
      </el-form-item>
      <el-form-item class="pwd" label="验证">
        <el-input v-model="testCode" style="width: 200px;" placeholder="请输入您的验证码" />
        <el-button class="getCode" type="primary" round>获取验证码</el-button>
      </el-form-item>
      <el-form-item class="pwd" label="密码">
        <el-input v-model="password" style="width: 270px;" placeholder="设置您的密码" />
      </el-form-item>
      
      <el-button class="btn"  type="primary">返回登录</el-button>
    </el-form>
  </div>
</template>
<style lang="css" scoped>
.container{
  width: 400px;
  height: 600px;
  border: 1px solid red;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pwd{
  margin-top: 30px;
  
}
.getCode{
  position: absolute;
  left: 170px;
}
.btn{
  position: relative;
  top: 50px;
}
.tit{
  position: absolute;
  top: 220px;
}
=======
<script setup>
import {reactive, ref} from 'vue'
const user = reactive({
  email:'',
  testCode:'',
  password:''
})
const blank = () => {
 
            if(user.password==''|| user.email == '' || user.testCode=='')
            {
                 ELMessage({
                                message: '账号密码邮箱验证码不能为空!!!',
                                center:true,
                                type: 'error'
                            })
            }
            // else {
                 
            //         this.$api.isRepeat({
            //             Username:this.user.Username
            //         })
            //         .then(res=>{
            //             console.log(res.data);
            //             if(res.data.status_code==3001)
            //             this.$message({
            //                     message: '用户名已存在!!',
            //                     center:true,
            //                     type: 'error'
            //                 })
            //         })
            //         this.$api.getRegister({
            //             Email: this.user.email,
            //             Password:this.user.Password,
            //             Username:this.user.Username,
            //             VerifyCode:this.user.comfirm,
            //             InviteCode:this.user.invite
            //         })
            //         .then(res=>{
            //             console.log(res);
            //             if(res.data.status_code==0){
            //                 this.$message({
            //                         message: '注册成功',
            //                         center: true,
            //                         type: 'success'
            //                         });
            //                 this.change();
            //             }
            //             if(res.data.status_code==3007){
            //                 this.$message({
            //                     message: '验证码错误或过期!!',
            //                     center:true,
            //                     type: 'error'
            //                 })
            //             }
            //         })
            //         .catch (err => {
            //         console.log(err);
            // })
                
            // }
    
}
</script>
<template>
  <div class="container">
    <div class="tit">magolia影院</div>
    <el-form>
      <el-form-item class="act" label="邮箱">
        <el-input v-model="e_mail" style="width: 270px;" placeholder="请输入您的邮箱" />
      </el-form-item>
      <el-form-item class="pwd" label="验证">
        <el-input v-model="testCode" style="width: 200px;" placeholder="请输入您的验证码" />
        <el-button class="getCode" type="primary" round>获取验证码</el-button>
      </el-form-item>
      <el-form-item class="pwd" label="密码">
        <el-input v-model="password" style="width: 270px;" placeholder="设置您的密码" />
      </el-form-item>
      
      <el-button class="btn"  type="primary">返回登录</el-button>
    </el-form>
  </div>
</template>
<style lang="css" scoped>
.container{
  width: 400px;
  height: 600px;
  border: 1px solid red;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pwd{
  margin-top: 30px;
  
}
.getCode{
  position: absolute;
  left: 170px;
}
.btn{
  position: relative;
  top: 50px;
}
.tit{
  position: absolute;
  top: 220px;
}
>>>>>>> 73d85f9 (添加注册功能)
</style>