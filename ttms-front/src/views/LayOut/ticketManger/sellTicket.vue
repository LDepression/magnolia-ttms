<script setup></script>
<template>
    <div class="board">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>票务管理</el-breadcrumb-item>
        <el-breadcrumb-item>售票管理</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
     <!--卡片视图-->
     <el-card class="box-card">
      
      <el-row :gutter="20">
        <el-col :span="5">
          <el-input v-model="inputUserName" placeholder="请输入用户名" clearable></el-input>
        </el-col>
        <el-col :span="4">
          <el-select v-model="selectedState" placeholder="请选择订单状态" clearable>
            <el-option v-for="item in payStates" :key="item.id" :label="item.name" :value="item.id"></el-option>
          </el-select>
        </el-col>

        <el-col :span="5">
          <el-date-picker
              v-model="selectedDate"
              type="date"
              value-format="yyyy-MM-dd"
              placeholder="选择日期时间">
          </el-date-picker>
        </el-col>

        <el-col :span="2">
          <el-button icon="el-icon-search" @click="getBillList">搜索</el-button>
        </el-col>
        <el-col :span="2">
          <el-button type="danger" @click="multipleDelete">批量删除</el-button>
        </el-col>
      </el-row>

      <!--订单列表-->
      <el-table :data="billList" style="width: 100%;  font-size:14px;" border stripe @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55"></el-table-column>
        <el-table-column prop="billId" label="#" width="40"></el-table-column>
        <el-table-column prop="sysUser.userName" label="用户名" width="80px"></el-table-column>
        <el-table-column prop="sysSession.sysMovie.movieName" label="电影名称"></el-table-column>
        <el-table-column prop="sysSession.sysHall.hallName" label="影厅名称"></el-table-column>
        <el-table-column prop="sysSession.sessionDate" label="场次日期"></el-table-column>
        <el-table-column prop="sysSession.playTime, sysSession.endTime" label="播放时间">
          <template slot-scope="scope"> {{scope.row.sysSession.playTime}}-{{scope.row.sysSession.endTime}} </template>
        </el-table-column>
        <el-table-column prop="seats" label="座位"></el-table-column>
        <el-table-column prop="createTime" label="订票时间"></el-table-column>
        <el-table-column prop="payState" label="订单状态" width="80">
          <template slot-scope="scope">
            <span v-if="scope.row.payState === true" style="color: #13ce66">已完成</span>
            <span v-if="scope.row.payState !== true && scope.row.cancelState !== true " style="color: #145ddc">未支付</span>
            <span v-if="scope.row.cancelState === true && scope.row.cancelTime ===null" style="color: #e6a23c">超时取消</span>
            <span v-if="scope.row.cancelState === true && scope.row.cancelTime !==null" style="color: crimson">用户取消</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template slot-scope="scope">
            <el-tooltip effect="dark" content="修改订单信息" placement="top" :enterable="false" :open-delay="500">
              <el-button type="primary" icon="el-icon-edit" size="mini"
                         @click="showEditDialog(scope.row.billId)"></el-button>
            </el-tooltip>
            <el-tooltip effect="dark" content="删除订单" placement="top" :enterable="false" :open-delay="500">
              <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteBillById(scope.row.billId)"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <!--分页区域-->
      <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="queryInfo.pageNum"
          :page-sizes="[5, 7, 9]"
          :page-size="queryInfo.pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total">
      </el-pagination>
    </el-card>

</template>
<style lang="css" scoped></style>