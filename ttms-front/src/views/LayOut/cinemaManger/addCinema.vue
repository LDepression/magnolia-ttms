<script setup></script>
<template>
  <div class="container">
    <el-form class="form">
      <el-form-item class="act" label="演出厅编号">
        <el-input style="width: 200px" type="text" placeholder="请输入演出厅编号" />
      </el-form-item>
      <el-form-item class="pwd" label="演出厅名称">
        <el-input style="width: 200px" type="text" placeholder="请输入演出厅名称" />
      </el-form-item>
      <el-form-item class="act" label="座位&nbsp&nbsp 行数">
        <el-input
        
          style="width: 200px"
          type="number"
          placeholder="请输入座位行数"
        />
      </el-form-item>
      <el-form-item class="act" label="座位&nbsp&nbsp  列数">
        <el-input
          style="width: 200px"
          type="number"
          placeholder="请输入座位列数"
        />
      </el-form-item>
      <el-form-item class="act" label="演出厅描述">
        <el-input style="width: 200px" type="text" placeholder="请输入演出厅描述" />
      </el-form-item>
      <el-button class="btn" type="primary">提交</el-button>
      <el-button class="btn" type="primary">返回</el-button>
    </el-form>
  </div>
</template>
<style scoped>
.container {
  width: 400px;
  height: 500px;
  border: 1px solid red;
  display: flex;
  position: absolute;
  left: 200px;
  top: 0;
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