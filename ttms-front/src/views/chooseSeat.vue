<script setup>
import { onMounted, ref, reactive,nextTick } from "vue";
import { getSeat } from "../apis/seat";
import { toPay, tryLock, isPay } from "../apis/order";
import { ElMessage } from "element-plus";
import QrcodeVue from 'qrcode.vue'
import { useRoute, useRouter } from "vue-router";
const seat = reactive([]);
const selectSeat = reactive([]);
const isLock = ref(false);

const timeRemaining = reactive({
  minutes: 15,
  seconds: 0,
});
const router = useRouter()
const exist = ref(false);
const seatIDs = reactive([])
const poll = ref(null);
const route = useRoute();
const IsLock = ref()
const planID = route.query.planID;
const movieInfo = JSON.parse(route.query.movieInfo);
console.log(movieInfo);
const qrcodeRef = ref(null);
// import { v4 as uuidv4 } from 'uuid';
// const orderID = uuidv4()
const orderID = ref();
console.log(movieInfo[0]);
const changeState = (row, column) => {
  let newState = seat[row][column];
  console.log(newState);
  let seatID = newState.seatID;
  if (
    newState.Status == "normal" &&
    newState.TicketStatus == "for_sale" &&
    selectSeat.length < 5
  ) {
    newState.Status = "选中";
    selectSeat.push([row + 1, column + 1]);
  } else {
    for (let i in selectSeat) {
      if (selectSeat[i][0] == row + 1 && selectSeat[i][1] == column + 1)
        selectSeat.splice(i, 1);
      newState.Status = "normal";
    }
  }

  seat[row].splice(column, 1, newState);
};

const deadline = () => {
  poll.value = setInterval(payment, 1000);
  let timer = setInterval(() => {
    if (timeRemaining.minutes >= 0) {
      if (timeRemaining.minutes == 0 && timeRemaining.seconds == 0) {
        exist.value = false;
        clearInterval(timer);
        clearInterval(poll.value);
        return;
      }
      if (timeRemaining.seconds == 0) {
        timeRemaining.seconds = 60;
        timeRemaining.minutes--;
      }
      timeRemaining.seconds--;
    }
  });
};
import QRCode  from "qrcode"
// const getCode = (orderID) => {
//   QRCode.toDateURL(orderID)
// }
function qrcode() {
      let qrcode = new QRCode('qrcode', {
        width: 124,
        height: 124,
        text: 'https://img1.baidu.com/it/u=1960110688,1786190632&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1686416400&t=f72539efe87c5aa880fb470b346b5db0',
      });
    }
const payment = async (orderID, seatID, planID, cinemaID) => {
  //调用是否支付
  let res = await isPay(orderID);
  if (res.data.status_code == 0) {
    if (res.data.data.exist == true) {
      clearInterval(poll);
      payorder(orderID, seatID, planID, movieInfo[0].cinema_id);
    }
  }
};
const ToPay = () => {
  if (!confirm("确定购买吗？")) {
    return;
  }
  if (selectSeat.length > 0) {
    let arr = [];
    for (let i of selectSeat) {
      arr.push(seat[i[0] - 1][i[1] - 1]);
    }
    console.log(arr);
    rocking(arr);
    if (exist.value == false) {
      exist.value = true;
    }
  } else {
    ElMessage({
      message: "请选座",
      type: "warning",
    });
  }
};
const payorder = async (orderID, seatID, planID, cinemaID) => {
  let arr = [];
  for (let i of selectSeat) {
    arr.push(seat[i[0] - 1][i[1] - 1]);
  }
  console.log(arr);
  let res = await toPay({ orderID, seatIDs, planID, cinemaID });
  console.log(res);
  if (res.data.status_code == 0) {

    exist.value = false;
    ElMessage({
      message: "支付成功",
      type: "success",
    });
    router.push({ path: "/myOrder" });
  }
  //调用toPay
};

const rocking = async (arr) => {
  let user_id = JSON.parse(localStorage.getItem("user_info")).user_id;
  //调用锁票
  console.log(arr);
  console.log(arr[0].SeatID);
  for (let i in arr) {
    seatIDs.push(arr[i].SeatID)
  }
  console.log(seatIDs);
  
  let res = await tryLock(Number(planID),movieInfo[0].cinema_id,seatIDs)
  console.log(res);
  IsLock.value = res.data.data.IsLock;
  console.log(IsLock.value);
  orderID.value = res.data.data.orderID;
  console.log(orderID.value);
  payment(orderID.value, seatIDs, Number(planID), movieInfo[0].cinema_id)
  // nextTick(()=>{
  //   qrcode()
  // })
  if (res.data.status_code == 0) {
    ElMessage({
      message: "请尽快支付",
      type: "warning",
    });
  } else {
    ElMessage({
      message: "请刷新重试",
      type: "warning",
    });
  }
};
// const rocking = () => {
//   let user_id = JSON.parse(localStorage.getItem("user")).userid;
//             let plan_id=this.seat[0][0].plan_id;
//             //调用锁票
// }
// const toPay = () => {
//   if(!confirm("确定购买吗？"))
//             {
//                 return
//             }
//             if(selectSeat.length>0)
//             {
//                 let arr=[];
//                 for(let i of selectSeat)
//                 {
//                     arr.push(seat[i[0]-1][i[1]-1].seat_id)
//                 }
//                 rocking(arr)
//                 if(this.isPay==false)
//                 {
//                     this.isPay=true;
//                 }
//             }
//             else{
//                ElMessage({
//                         message: '请选座',
//                         type: 'warning'
//                         });
//             }
// }
const getSeatList = async (planID) => {
  let res = await getSeat(planID);
  console.log(res.data.data);
  seat.splice(0, seat.length);
  for (let i = 0; i < res.data.data.Tickets.length; i++) {
    seat.push(res.data.data.Tickets[i]);
  }
};
onMounted(() => {
  getSeatList(planID);
});
</script>
<template>
  <div class="container">
    <div class="seat">
      <div class="flag">
        <img src="../assets/img/seat/可选座位.png" alt="zuowei" />
        <p style="color:black">可选座位</p>
        <img src="../assets/img/seat/已售座位.png" alt="zuowei" />
        <span  style="color:black">已售座位</span>
        <img src="../assets/img/seat/已选座位.png" alt="zuowei" />
        <span  style="color:black">已选座位</span>
      </div>
      <div class="movieTv">
        <img src="../assets/img/yingmu.jpg" alt="" />
        <span>银幕中央</span>
      </div>

      <!-- TicketStatus: lock for_sale saled -->
      <!-- Status broken normal -->
      <div class="seatList">
        <div class="seat1" v-for="index in seat.length" :key="index">
          <div
            class="seat2"
            v-for="(item, x) in seat[index - 1]"
            :key="item.SeatID"
          >
            <!-- <img src="../assets/img/seat/可选座位.png" alt=""> -->
            <img
              v-show="
                item.Status == 'normal' && item.TicketStatus == 'for_sale'
              "
              @click="changeState(index - 1, x)"
              src="../assets/img/seat/可选座位.png"
              alt="zuowei"
            />
            <img
              v-show="
                item.TicketStatus == 'saled' ||
                item.TicketStatus == 'lock' ||
                item.Status == 'broken'
              "
              src="../assets/img/seat/已售座位.png"
              alt="zuowei"
            />
            <img
              v-show="item.Status == '选中'"
              @click="changeState(index - 1, x)"
              src="../assets/img/seat/已选座位.png"
              alt="zuowei"
            />
          </div>
          <!-- <div v-show="item.seats_status=='走廊'" class="emptySeat"></div> -->
        </div>
      </div>
    </div>

    <div class="movie">
      <div class="header">
        <img :src="movieInfo[0].movieAvatar" alt="" />
        <div class="title">
          <h2></h2>

          <p class="labelname">
            时长：{{ movieInfo[0].duration }}<span> 分钟</span>
          </p>
        </div>
        <p class="labelname">
          影厅：<span>{{ movieInfo[0].cinemaName }}</span>
        </p>
        <p class="labelname">
          场次：<span>{{ movieInfo[0].startAt.slice(0, 10) }} </span>
        </p>
        <p class="labelname">
          票价：<span>￥{{ movieInfo[0].price }}/张</span>
        </p>
      </div>
      <div class="price">
        <p class="labelname">
          座位：
          <el-tag
            v-for="(item, index) in selectSeat"
            :key="index"
            color="#f56c6c"
            style="color: #fff"
            >{{ item[1] }}排{{ item[0] }}座</el-tag
          >
        </p>
        <div class="allprice">
          总价：
          <div>￥{{ movieInfo[0].price * selectSeat.length }}</div>
        </div>
      </div>
      <div class="enter">
        <el-button
          @click="ToPay()"
          @click.once="deadline()"
          style="width: 300px"
          type="danger"
          round
          >确认选座</el-button
        >
      </div>
    </div>
  </div>
  <div class="pay" v-if="exist">
    <div class="time"> {{timeRemaining.minutes}}:{{timeRemaining.seconds>=10?timeRemaining.seconds:'0'+timeRemaining.seconds}}</div>
    <!-- <div id="qrcode"></div> -->
    <qrcode-vue :value="orderID"></qrcode-vue>
  </div>
</template>
<style lang="css" scoped>
.seat1 {
  display: inline-block;
}

.time {
  font-size: 35px;
  text-align: center;
}
#qrcode {
  display: flex;
  justify-content: center;
  margin: 40px 0px;
}
.pay {
  padding: 20px;
  position: fixed;
  width: 300px;
  border: 1px solid #ccc;
  box-shadow: 3px 3px 3px #ccc;
  top: 50%;
  left: 50%;
  background: #fff;
  text-align: center;
  transform: translate(-50%, -50%);
}
.enter {
  text-align: center;
  margin: 40px 0;
}
.allprice {
  color: #000;
  padding: 30px 0px;
  border-bottom: 1px solid rgb(155, 150, 150);
}
.allprice > div {
  display: inline-block;
  color: #f56c6c;
  font-size: 30px;
}
.labelname {
  color: rgb(155, 150, 150);
  font-size: 14px;
}
.labelname span {
  color: #000;
}
.labelname:nth-child(6) {
  padding-bottom: 10px;
  border-bottom: 1px solid rgb(155, 150, 150);
}
.title {
  display: inline-block;
  margin-left: 20px;
  vertical-align: top;
}
.header img {
  width: 160px;
  height: 222.62px;
}
.title h2 {
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  width: 230px;
}
.movie {
  width: 40%;
  padding: 54px;
  position: relative;
  left: 10%;
  background-color: #f3f0f0;
}
.flag {
  position: relative;
  left: 120px;
}
.emptySeat {
  width: 48px;
  height: 48px;
}
.seatRowList {
  list-style: none;
  white-space: nowrap;
}
.container {
  position: relative;
  left: 20%;
}
.seatItem {
  display: inline-block;
  margin-left: 6px;
  cursor: pointer;
}
.emptySeat {
  margin: 0px;
}
.seatList {
  overflow-x: auto;
  list-style: none;
}
/* 滚动条宽度 */
.seatList::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}
/* 滚动条轨道 */
.seatList::-webkit-scrollbar-track {
  background: rgb(239, 239, 239);
  border-radius: 2px;
}
/*  小滑块 */
.seatList::-webkit-scrollbar-thumb {
  background: #40a0ff49;
  border-radius: 10px;
}
.seatList::-webkit-scrollbar-thumb:hover {
  background: #40a0ff;
}

.seatRow {
  margin-top: 6px;
  display: flex;
  align-items: center;
  color: #999;
}
.flag {
  display: flex;
  align-items: center;
  margin-left: 80px;
}
.flag img {
  margin-left: 40px;
  margin-right: 5px;
}
.main {
  margin: 0 auto;
  margin-top: 100px;
  display: flex;
  border-top: 1px solid #ccc;
  border-bottom: 1px solid #ccc;
}
.seat {
  width: 70%;
}
.seatlayout > div {
  text-align: center;
  color: #999;
  margin: 15px 0;
}
</style>