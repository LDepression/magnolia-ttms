<script setup>
import { onMounted, reactive } from 'vue';
import {getAllOrder, returnTicket} from '../../../apis/order'
const orderList = reactive([])
const  getList = async () => {
  let res = await getAllOrder()
  orderList.splice(0,orderList.length,...res.data.data.OrderInfos)
  console.log(orderList);
}
const back = async (OrderID,planID,SeatIDs) => {
  let res = await returnTicket(OrderID,planID,SeatIDs)
  console.log(res);
  getList()
}
onMounted(()=>{getList()})
</script>
<template>
  <div class="top">
   我的订单
  </div>
  <hr>
  <div class="container" v-for="(item,index) in orderList" :key="index">
    <div class="littleTop">
      <div class="time">{{item.UpdateTime}}</div>
      <div>订单号：20030812</div>
      <div>ia</div>
    </div>
    <div class="mainContext">
      <img src="../../../assets/img/2.gif" alt="">
      <div class="description">
        <div><h3>《复仇者联盟》</h3></div>
        <div>软协影院</div>
        <div>{{item.CinemaName}}{{item.SeatIDs.slice(0,2)}}号座位</div>
        <div>周三 6月1日 14:30</div>
      </div>
      <div class="money">￥{{item.Price}}</div>
      <div class="status">{{item.Status}}</div>
      <div class="detail">查看详情</div>
      <el-button class="newInfo" type="primary" v-if="item.Status=='已支付'" @click="back(item.OrderID,item.PlanID,[Number(item.SeatIDs)])">退票</el-button>
    </div>
  </div>
  <!-- <div class="container">
    <div class="littleTop">
      <div class="time">2022-09-15</div>
      <div>订单号：20030812</div>
      <div>ia</div>
    </div>
    <div class="mainContext">
      <img src="../../../assets/img/2.gif" alt="">
      <div class="description">
        <div><h3>《复仇者联盟》</h3></div>
        <div>软协影院</div>
        <div>3号厅 7排9座</div>
        <div>周三 6月1日 14:30</div>
      </div>
      <div class="money">￥40</div>
      <div class="status">已完成</div>
      <div class="detail">查看详情</div>
    </div>
  </div> -->
</template>
<style lang="css" scoped>
.container{
  border: 1px solid rgb(243, 243, 243);
  font-size: 14px;
  margin-top: 10px;
}
.newInfo{
  position: relative;
  top: 55px;
}
.description{
  color: black;
  display: flex;
  height: 80%;
  margin: auto;
  flex-direction: column;
  margin-left: 30px;
  
}
.description div{
  font-size: 12px;

}
.time{
  position: relative;
  left: -100px;
}
.mainContext{
  width: 100%;
 
  height: 150px;
  background-color: rgb(red, green, blue);
  display: flex;
  justify-content: space-between;
  color: black;
  font-size: 14px;
}
.money{
  position: relative;
  left: -300px;
  top: 34%;
}
.detail{
  position: relative;
  left: -30px;
  top: 40%;
}
.status{
  position: relative;
  left: -150px;
  top: 40%;
}
.mainContext img{
  width: 70px;
  height: 100px;
  margin-top: 25px;
  margin-left: 20px;
}
.top{
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  color: red;
}
.littleTop{
  width: 100%;
  height: 50px;
  display: flex;
  justify-content: space-around;
  align-items: center;
  color: black;
  background-color:rgb(243, 243, 243) ;
}
/* .littleTop div{
  margin-left: 50px;
} */
</style>